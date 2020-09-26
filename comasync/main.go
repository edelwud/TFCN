package main

import (
	"./pkg/gui"
	"./pkg/serial"
	"github.com/therecipe/qt/widgets"
	"os"
)

const (
	TransmitterComm = "COM1"
	ReceiverComm    = "COM2"
	WindowWidth     = 400
	WindowHeight    = 600
	WindowTitle     = "COM Async Library"
)

var config = serial.Config{
	BaudRate:       9600,
	ByteSize:       8,
	Parity:         1,
	StopBits:       1,
	MaxReadBuffer:  4096,
	MaxWriteBuffer: 4096,
	ReadTimeout:    10,
	WriteTimeout:   100,
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	transmitter, err := serial.Open(TransmitterComm, &config)
	if err != nil {
		gui.ShowErrorMessage(err.Error())
		return
	}

	receiver, err := serial.Open(ReceiverComm, &config)
	if err != nil {
		gui.ShowErrorMessage(err.Error())
		return
	}

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle(WindowTitle)
	window.SetFixedWidth(WindowWidth)
	window.SetFixedHeight(WindowHeight)

	centralWidget := widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(gui.InitGUI(transmitter, receiver))
	window.SetCentralWidget(centralWidget)
	window.Show()

	widgets.QApplication_Exec()
}
