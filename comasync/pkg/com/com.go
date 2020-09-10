package com

import (
	"github.com/kbinani/win"
	"golang.org/x/sys/windows"
)

type COM struct {
	Handle win.HANDLE
}

func (com *COM) Initialize(comPortName string) {
	com.Handle = win.CreateFile(
		comPortName,
		windows.GENERIC_READ|windows.GENERIC_WRITE,
		0,
		nil,
		windows.OPEN_EXISTING,
		windows.FILE_FLAG_OVERLAPPED,
		0)

	if com.Handle == 0 {
		panic("Cannot to connect to " + comPortName)
	}
}
