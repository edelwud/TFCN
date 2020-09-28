package serial

type Packet struct {
	Start byte
	Data  []byte
	End   byte
}

const BitStuffingFlag = "0110110"
const BitToStuff = '0'
const CompletedFlag = '1'

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

func (packet Packet) ToString() string {
	return string(packet.Data)
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
		if len(temp) >= 7 {
			if string(temp[len(temp)-7:]) == BitStuffingFlag {
				result = append(result, BitToStuff)
			}
		}
	}

	response := make([]byte, 0)
	response = append(response, []byte(BitStuffingFlag)...)
	response = append(response, CompletedFlag)
	response = append(response, result...)
	response = append(response, []byte(BitStuffingFlag)...)
	response = append(response, CompletedFlag)

	packet.Data = response
}

func (packet *Packet) DeBitStuffing() {
	var result []byte
	var flag = false
	for _, bit := range packet.Data[8 : len(packet.Data)-9] {
		if flag {
			flag = false
			continue
		}
		result = append(result, bit)
		if len(result) >= 7 {
			if string(result[len(result)-7:]) == BitStuffingFlag {
				flag = true
			}
		}
	}
	packet.Data = result
}
