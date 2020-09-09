package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func InitApplication(application string, windowTitle string) fyne.Window {
	a := app.NewWithID(application)
	a.Settings().SetTheme(theme.LightTheme())

	w := a.NewWindow(windowTitle)
	w.SetMaster()

	programInput := widget.NewMultiLineEntry()
	programInput.OnChanged = func(newString string) {
		println(newString)
	}
	inputBox := widget.NewVBox(
		widget.NewLabel("Text input"),
		programInput,
	)

	programResult := widget.NewMultiLineEntry()
	programResult.Disable()

	outputBox := widget.NewVBox(
		widget.NewLabel("Program result"),
		programResult,
	)

	status := CreateStatusBox()

	w.SetContent(widget.NewVBox(
		inputBox,
		outputBox,
		status,
	))
	return w
}