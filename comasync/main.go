package main

import (
	gui "./pkg/gui"
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("COM Async Library")

	var centralWidget = widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(gui.InitLayout())
	window.SetCentralWidget(centralWidget)

	window.Show()

	widgets.QApplication_Exec()
}
