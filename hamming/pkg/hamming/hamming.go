package hamming

import (
	"math"
)

const SequenceLength = 17

func GenerateHammingCode(buf []byte) (result []byte) {
	result = append(result, buf...)

	var index = 0
	var power float64 = 1
	var base float64 = 2

	for index < len(result) {
		var temp []byte
		if index > 0 {
			temp = append(temp, result[:index]...)
		}
		temp = append(temp, 48)
		temp = append(temp, result[index:]...)
		result = temp

		index = int(math.Pow(base, power)) - 1
		power += 1
	}

	power = 0
	index = 0
	step := 2

	for (step - 1) < len(result) {
		index = step - 2
		entity := 0
		for index < len(result) {
			k := 0
			for k < step-1 {
				if index >= len(result) {
					break
				}
				if result[index] == '1' {
					entity += 1
				}
				k += 1
				index += 1
			}
			index += step - 1
		}
		if entity%2 == 1 {
			result[step-2] = 49
		}
		step = int(math.Pow(base, power)) + 1
		power += 1
	}

	return result
}

func ValidateHammingCode(buf []byte) (result []byte) {
	var power float64 = 0
	var base float64 = 2
	var index = 0
	var step = 2

	result = append(result, buf...)

	var errorIndex int

	for step-1 < len(result) {
		index = step - 2
		entity := 0
		flag := true

		for index < len(result) {
			k := 0

			if flag {
				flag = false
				index += 1
				k += 1
			}

			for k < step-1 {
				if index >= len(result) {
					break
				}
				if result[index] == '1' {
					entity += 1
				}
				k += 1
				index += 1
			}
			index += step - 1
		}

		if (entity % 2) != int(result[step-2]-48) {
			errorIndex += step - 2
		}
		power += 1
		step = int(math.Pow(base, power)) + 1
	}

	if errorIndex != 0 {
		if result[errorIndex] == '0' {
			result[errorIndex] = '1'
		} else {
			result[errorIndex] = '0'
		}
	}
	return result
}

func ParseMessage(buffer []byte) string {
	var base float64 = 2
	var power float64 = 0
	var result []byte

	for index, bit := range buffer {
		if index+1 == int(math.Pow(base, power)) {
			power += 1
			continue
		}
		result = append(result, bit)
	}
	return string(result)
}

func SplitUserInput(userInput string) (result []string) {
	for len(userInput) >= SequenceLength {
		result = append(result, userInput[:SequenceLength])
		userInput = userInput[SequenceLength:]
	}
	return result
}
