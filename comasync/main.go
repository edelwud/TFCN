package main

import (
	"./pkg/gui"
	"./pkg/serial"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	config := serial.Config{
		BaudRate:       256000,
		ByteSize:       8,
		Parity:         1,
		StopBits:       0,
		MaxReadBuffer:  4096,
		MaxWriteBuffer: 4096,
		ReadTimeout:    10,
		WriteTimeout:   10,
	}

	transmitter, err := serial.Open("COM1", &config)
	if err != nil {
		gui.ShowErrorMessage(err.Error())
		return
	}

	receiver, err := serial.Open("COM2", &config)
	if err != nil {
		gui.ShowErrorMessage(err.Error())
		return
	}

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("COM Async Library")

	centralWidget := widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(gui.InitGUI(transmitter, receiver))
	window.SetCentralWidget(centralWidget)
	window.Show()

	widgets.QApplication_Exec()
}
