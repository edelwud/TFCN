package com

import "github.com/Andeling/serial"

type Port interface {
	Write(buf []byte) (int, error)
	Read(buf []byte) (int, error)
	Close() error
}

func OpenCOM(comPort string) (Port, error) {
	com, err := serial.Open(comPort, &serial.Config{
		BaudRate:    115200,
		DataBits:    8,
		Parity:      serial.ParityNone,
		StopBits:    serial.StopBitsOne,
		FlowControl: serial.FlowControlXonXoff,
	})
	if err != nil {
		return nil, err
	}
	return com, nil
}
