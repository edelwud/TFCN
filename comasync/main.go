package main

import (
	gui "./pkg/gui"
)

func main() {
	w := gui.InitApplication("com.edelwud.comasync", "COM Async library")
	w.ShowAndRun()
}