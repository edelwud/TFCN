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

func (config *Config) Serialize() map[string]string {
	var parity string
	switch config.Parity {
	case 0:
		parity = "NO PARITY"
		break
	case 1:
		parity = "ODD PARITY"
		break
	case 2:
		parity = "EVEN PARITY"
		break
	case 3:
		parity = "MARK PARITY"
		break
	case 4:
		parity = "SPACE PARITY"
		break
	}

	var stopBits string
	switch config.StopBits {
	case 0:
		stopBits = "1 STOP BIT"
		break
	case 1:
		stopBits = "1.5 STOP BITS"
		break
	case 2:
		stopBits = "2 STOP BITS"
		break
	}

	return map[string]string{
		"Baud rate":             strconv.Itoa(int(config.BaudRate)),
		"Byte size":             strconv.Itoa(int(config.ByteSize)),
		"Parity":                parity,
		"Stop bits":             stopBits,
		"Max read buffer size":  strconv.Itoa(int(config.MaxReadBuffer)),
		"Max write buffer size": strconv.Itoa(int(config.MaxWriteBuffer)),
		"Read timeout":          strconv.Itoa(int(config.ReadTimeout)),
		"Write timeout":         strconv.Itoa(int(config.WriteTimeout)),
	}
}
