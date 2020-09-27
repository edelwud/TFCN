package serial

import (
	"golang.org/x/sys/windows"
	"sync"
	"unsafe"
)

type Serial interface {
	Open(string) error
	Close() error
	Write([]byte) error
	Read() (*Packet, error)
	GetConfig() Config
}

type SerialPort struct {
	Name     string
	DCB      *DCB
	Timeouts *CommTimeouts
	Handle   windows.Handle
	Config   *Config
	Mux      sync.Mutex
}

const (
	CleatInBuffer         = 0x0008
	ClearOutBuffer        = 0x0004
	CancelWriteOperations = 0x0001
	CancelReadOperations  = 0x0002
)

var (
	kernel32                = windows.NewLazyDLL("kernel32.dll")
	procSetupComm           = kernel32.NewProc("SetupComm")
	procGetOverlappedResult = kernel32.NewProc("GetOverlappedResult")
	procPurgeComm           = kernel32.NewProc("PurgeComm")
)

func (port *SerialPort) Open(com string) error {
	var err error
	port.Handle, err = windows.CreateFile(
		windows.StringToUTF16Ptr("\\\\.\\"+com),
		windows.GENERIC_WRITE|windows.GENERIC_READ,
		0,
		nil,
		windows.OPEN_EXISTING,
		windows.FILE_ATTRIBUTE_NORMAL|windows.FILE_FLAG_OVERLAPPED,
		windows.InvalidHandle)
	if err != nil {
		return err
	}

	return nil
}

func (port *SerialPort) GetConfig() Config {
	return *port.Config
}

func (port *SerialPort) Close() error {
	if err := windows.CloseHandle(port.Handle); err != nil {
		return err
	}
	return nil
}

func (port *SerialPort) Clear(flags uint32) error {
	if r, _, err := procPurgeComm.Call(uintptr(port.Handle), uintptr(flags)); r == 0 {
		return err
	}
	return nil
}

func (port *SerialPort) Write(buffer []byte) error {
	port.Mux.Lock()
	var overlapped windows.Overlapped

	packet := NewPacket(buffer, XoffSymbol, XonSymbol)
	packet.BitStuffing()
	if err := port.Clear(ClearOutBuffer | CancelWriteOperations); err != nil {
		return err
	}

	if err := windows.WriteFile(
		port.Handle,
		packet.ToBytes(),
		nil,
		&overlapped,
	); err != windows.ERROR_IO_PENDING {
		return err
	}
	port.Mux.Unlock()
	return nil
}

func (port *SerialPort) Read() (*Packet, error) {
	buffer := make([]byte, port.GetConfig().MaxReadBuffer)

	var overlapped windows.Overlapped
	var err error
	overlapped.HEvent, err = windows.CreateEvent(nil, 1, 0, nil)
	if err != nil {
		return nil, err
	}

	if err := port.Clear(CancelReadOperations); err != nil {
		return nil, err
	}

	var read uint32
	err = windows.ReadFile(port.Handle, buffer, &read, &overlapped)
	if err == nil {
		return nil, nil
	}
	if err != windows.ERROR_IO_PENDING {
		return nil, err
	}
	if r, _, err := procGetOverlappedResult.Call(uintptr(port.Handle),
		uintptr(unsafe.Pointer(&overlapped)),
		uintptr(unsafe.Pointer(&read)), 1); r == 0 {
		return nil, err
	}
	if err := windows.CloseHandle(overlapped.HEvent); err != nil {
		return nil, err
	}
	if read == 0 {
		return nil, err
	}

	return NewPacket(buffer[:read], XoffSymbol, XonSymbol), nil
}

func Open(com string, config *Config) (Serial, error) {
	serial := &SerialPort{}
	serial.Config = config
	if err := serial.Open(com); err != nil {
		return nil, err
	}

	if r, _, err := procSetupComm.Call(
		uintptr(serial.Handle),
		uintptr(config.MaxReadBuffer),
		uintptr(config.MaxWriteBuffer),
	); r == 0 {
		return nil, err
	}

	serial.DCB = &DCB{}
	if err := serial.DCB.Build(serial.Handle, config); err != nil {
		return nil, err
	}

	serial.Timeouts = &CommTimeouts{}
	if err := serial.Timeouts.Configure(serial.Handle, config.ReadTimeout, config.WriteTimeout); err != nil {
		return nil, err
	}

	if err := serial.Clear(ClearOutBuffer | CleatInBuffer | CancelWriteOperations | CancelReadOperations); err != nil {
		return nil, err
	}

	return serial, nil
}
