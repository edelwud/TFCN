package gui

import (
	"fyne.io/fyne/widget"
)

func CreateStatusBox() *widget.AccordionContainer {
	status := widget.NewAccordionContainer()
	statusItem := widget.NewAccordionItem("Status", widget.NewLabel("Heke"))
	status.Append(statusItem)
	return status
}