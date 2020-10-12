package main

import (
	"./pkg/gui"
	"github.com/therecipe/qt/widgets"
	"os"
)

const (
	WindowWidth  = 500
	WindowHeight = 600
	WindowTitle  = "COM Async Library"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle(WindowTitle)
	window.SetFixedWidth(WindowWidth)
	window.SetFixedHeight(WindowHeight)

	centralWidget := widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(gui.InitGUI())
	window.SetCentralWidget(centralWidget)
	window.Show()

	widgets.QApplication_Exec()
}
