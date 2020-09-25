package serial

type Packet struct {
	Start byte
	Data  []byte
	End   byte
}

func NewPacket(data []byte, start byte, end byte) *Packet {
	return &Packet{
		start,
		data,
		end,
	}
}

func (packet Packet) ToBytes() []byte {
	packet.Data = append([]byte{packet.Start}, packet.Data...)
	packet.Data = append(packet.Data, packet.End)
	return packet.Data
}
