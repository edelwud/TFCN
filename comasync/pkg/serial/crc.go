package serial

import (
	"bytes"
	"errors"
	"hash/crc32"
)

var (
	ErrorChecksumConfirmation = errors.New("error: checksum confirmation error")
)

func ChecksumToSlice(checksum uint32) []byte {
	result := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		result[i] = byte((checksum >> (8 * i)) & 0xff)
	}
	return result
}

func GenerateCRCPacket(packet []byte) []byte {
	checksum := crc32.ChecksumIEEE(packet)
	return append(packet, ChecksumToSlice(checksum)...)
}

func ConfirmChecksum(packet []byte, length uint32) error {
	checksum := crc32.ChecksumIEEE(packet[:length-4])
	if bytes.Equal(ChecksumToSlice(checksum), packet[length-4:length]) == false {
		return ErrorChecksumConfirmation
	}
	return nil
}
