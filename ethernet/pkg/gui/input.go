package gui

import (
	"errors"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"regexp"
)

func ValidateTextEdit(transmitterTextEdit *widgets.QTextEdit) error {
	text := transmitterTextEdit.ToPlainText()
	otherSymbols := regexp.MustCompile("[^0-1]+")
	if index := otherSymbols.FindIndex([]byte(text)); index != nil {
		newText := text[:index[0]]
		newText += text[index[1]:]
		transmitterTextEdit.SetText(newText)

		newCursor := transmitterTextEdit.TextCursor()
		newCursor.SetPosition(len(newText), gui.QTextCursor__MoveAnchor)
		transmitterTextEdit.SetTextCursor(newCursor)
		return errors.New("symbols are not accepted")
	}
	return nil
}
