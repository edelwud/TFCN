package gui

import (
	h "../hamming"
	"github.com/pkg/errors"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"math"
	"regexp"
	"strings"
	"time"
)

func CreateStatusBox() (*widgets.QGroupBox, *widgets.QTreeWidget) {
	statusLayout := widgets.NewQGridLayout2()
	statusTable := CreateStatusTable()

	statusTable.SortItems(0, core.Qt__AscendingOrder)

	statusLayout.AddWidget(statusTable)

	statusGroup := widgets.NewQGroupBox2("Status table:", nil)
	statusGroup.SetLayout(statusLayout)
	return statusGroup, statusTable
}

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
		return errors.New("Symbols are not accepted")
	}
	return nil
}

func ConfigureBackgroundColor(data []byte, color string) []byte {
	var newBytes []byte
	newBytes = append(newBytes, []byte("<font style=\"background: "+color+"\">")...)
	newBytes = append(newBytes, data...)
	newBytes = append(newBytes, []byte("</font>")...)
	return newBytes
}

func CreateTransmitterBox(receiver *widgets.QTextEdit, status *widgets.QTreeWidget) *widgets.QGroupBox {
	transmitterTextEdit := widgets.NewQTextEdit(nil)
	transmitterTextEdit.SetPlaceholderText("Input text here")

	transmitterTextEdit.ConnectTextChanged(func() {
		if err := ValidateTextEdit(transmitterTextEdit); err != nil {
			return
		}

		userInput := transmitterTextEdit.ToPlainText()

		time.Sleep(time.Millisecond * 15)

		var message [][]byte
		for _, part := range h.SplitUserInput(userInput) {
			message = append(message, h.GenerateHammingCode([]byte(part)))
		}

		var result []string
		for _, buf := range message {
			result = append(result, h.ParseMessage(h.ValidateHammingCode(buf)))
		}

		receiver.SetText(strings.Join(result, ""))

		status.Clear()
		for _, buf := range message {
			AddRowToStatusTable(
				status,
				h.ParseMessage(h.ValidateHammingCode(buf)),
				BeautifyHammingCode(h.ValidateHammingCode(buf)))
		}
	})

	transmitterLayout := widgets.NewQGridLayout2()
	transmitterLayout.AddWidget(transmitterTextEdit)

	transmitterGroup := widgets.NewQGroupBox2("Transmitter:", nil)
	transmitterGroup.SetLayout(transmitterLayout)
	return transmitterGroup
}

func BeautifyHammingCode(buffer []byte) string {
	var base float64 = 2
	var power float64 = 0
	var result []byte
	var flag = false

	for index, bit := range buffer {
		if index+1 == int(math.Pow(base, power)) {
			power += 1
			flag = true
		}
		if flag {
			result = append(result,
				[]byte("<font style=\"background: red\">")...)
			result = append(result, bit)
			result = append(result, []byte("</font>")...)
			flag = false
		} else {
			result = append(result, bit)
		}
	}
	return string(result)
}

func CreateReceiverBox() (*widgets.QGroupBox, *widgets.QTextEdit) {
	receiverTextEdit := widgets.NewQTextEdit(nil)
	receiverTextEdit.SetReadOnly(true)

	receiverLayout := widgets.NewQGridLayout2()
	receiverLayout.AddWidget(receiverTextEdit)

	receiverGroup := widgets.NewQGroupBox2("Receiver:", nil)
	receiverGroup.SetLayout(receiverLayout)
	return receiverGroup, receiverTextEdit
}

func InitGUI() *widgets.QGridLayout {
	layout := widgets.NewQGridLayout2()

	receiver, receiverTextEdit := CreateReceiverBox()
	status, statusTable := CreateStatusBox()
	transmitter := CreateTransmitterBox(receiverTextEdit, statusTable)

	layout.AddWidget2(receiver, 1, 0, 0)
	layout.AddWidget2(transmitter, 0, 0, 0)
	layout.AddWidget2(status, 2, 0, 0)
	return layout
}
