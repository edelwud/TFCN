package main

import (
	_ "./pkg/com"
	"./pkg/gui"
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("COM Async Library")

	centralWidget := widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(gui.InitGUI())
	window.SetCentralWidget(centralWidget)
	window.Show()

	widgets.QApplication_Exec()
}
