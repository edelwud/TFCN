package serial

type Packet struct {
	Start byte
	Data  []byte
	End   byte
}

const BitStuffingFlag = "01101101"
const BitToStuff = '1'

func NewPacket(data []byte, start byte, end byte) *Packet {
	return &Packet{
		start,
		data,
		end,
	}
}

func (packet Packet) ToBytes() []byte {
	dataCopy := make([]byte, len(packet.Data))
	dataCopy = append([]byte{packet.Start}, packet.Data...)
	dataCopy = append(packet.Data, packet.End)
	return dataCopy
}

func (packet Packet) Get8BitArray() []string {
	var result []string
	dataCopy := make([]byte, len(packet.Data))
	copy(dataCopy, packet.Data)

	for i := 0; i < len(dataCopy)/8; i++ {
		result = append(result, string(dataCopy[:8]))
		dataCopy = dataCopy[8:]
	}
	if len(dataCopy) > 0 {
		result = append(result, string(dataCopy))
	}
	return result
}

func (packet *Packet) BitStuffing() {
	var result []byte
	var temp []byte
	for _, bit := range packet.Data {
		temp = append(temp, bit)
		result = append(result, bit)
		if len(temp) >= 8 {
			if string(temp[len(temp)-8:]) == BitStuffingFlag {
				result = append(result, BitToStuff)
			}
		}
	}
	packet.Data = result
}
