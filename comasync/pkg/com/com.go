package com

import (
	"github.com/kbinani/win"
	"golang.org/x/sys/windows"
)

type com struct {
	Handle win.HANDLE
}

func (c *com) Initialize(comPortName string) {
	c.Handle = win.CreateFile(
		comPortName,
		windows.GENERIC_READ|windows.GENERIC_WRITE,
		0,
		nil,
		windows.OPEN_EXISTING,
		windows.FILE_FLAG_OVERLAPPED,
		0)

	if c.Handle == 0 {
		panic("Cannot to connect to " + comPortName)
	}
}
