package gui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func ShowErrorMessage(reason string) {
	widgets.NewQMessageBox2(
		widgets.QMessageBox__Critical,
		"Error",
		reason,
		widgets.QMessageBox__Close,
		nil,
		core.Qt__Window).Show()
}
