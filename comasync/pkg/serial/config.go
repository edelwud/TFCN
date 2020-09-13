package serial

type Config struct {
	BaudRate       uint32
	ByteSize       uint8
	Parity         uint8
	StopBits       uint8
	MaxReadBuffer  uint32
	MaxWriteBuffer uint32
}
