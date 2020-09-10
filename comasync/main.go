package main

import (
	"./pkg/com"
	"./pkg/gui"
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	transmitter, err := com.OpenCOM("COM1")
	if err != nil {
		gui.ShowErrorMessage(err.Error())
		return
	}

	receiver, err := com.OpenCOM("COM2")
	if err != nil {
		gui.ShowErrorMessage(err.Error())
		return
	}

	widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("COM Async Library")

	centralWidget := widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(gui.InitGUI(transmitter, receiver))
	window.SetCentralWidget(centralWidget)
	window.Show()

	widgets.QApplication_Exec()
}
