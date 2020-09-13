package main

import (
	"./pkg/gui"
	"./pkg/serial"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	transmitter, err := serial.Open("COM1", &serial.Config{
		BaudRate: 115200,
		ByteSize: 8,
		Parity:   0,
		StopBits: 0,
	})
	if err != nil {
		gui.ShowErrorMessage(err.Error())
		return
	}

	receiver, err := serial.Open("COM2", &serial.Config{
		BaudRate: 115200,
		ByteSize: 8,
		Parity:   0,
		StopBits: 0,
	})
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
