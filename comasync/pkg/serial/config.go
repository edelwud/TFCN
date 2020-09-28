package serial

import "strconv"

type Config struct {
	BaudRate       uint32
	ByteSize       uint8
	Parity         uint8
	StopBits       uint8
	MaxReadBuffer  uint32
	MaxWriteBuffer uint32
	ReadTimeout    uint32
	WriteTimeout   uint32
}

const (
	XoffSymbol = 0xff
	XonSymbol  = 0x0
)

func (config *Config) Serialize() map[string]string {
	var parity string
	switch config.Parity {
	case 0:
		parity = "no parity"
		break
	case 1:
		parity = "odd parity"
		break
	case 2:
		parity = "even parity"
		break
	case 3:
		parity = "mark parity"
		break
	case 4:
		parity = "space parity"
		break
	}

	var stopBits string
	switch config.StopBits {
	case 0:
		stopBits = "1"
		break
	case 1:
		stopBits = "1.5"
		break
	case 2:
		stopBits = "2"
		break
	}

	return map[string]string{
		"Bit to stuff":          string(BitToStuff),
		"Frame flag":            BitStuffingFlag + string(CompletedFlag),
		"Baud rate":             strconv.Itoa(int(config.BaudRate)) + " baud",
		"Byte size":             strconv.Itoa(int(config.ByteSize)) + " bit",
		"Parity":                parity,
		"Stop bits":             stopBits + " bit",
		"Max read buffer size":  strconv.Itoa(int(config.MaxReadBuffer)) + " bytes",
		"Max write buffer size": strconv.Itoa(int(config.MaxWriteBuffer)) + " bytes",
		"Timeout read":          strconv.Itoa(int(config.ReadTimeout)) + " msec",
		"Timeout write":         strconv.Itoa(int(config.WriteTimeout)) + " msec",
	}
}
