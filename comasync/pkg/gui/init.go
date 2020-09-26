package gui

import (
	"../serial"
	"bytes"
	"github.com/pkg/errors"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"regexp"
	"time"
)

var (
	dataReceived = make(chan []byte)
)

func CreateStatusBox(transmitter serial.Serial, _ serial.Serial) *widgets.QGroupBox {
	statusLayout := widgets.NewQGridLayout2()
	statusTable := CreateStatusTable()

	config := transmitter.GetConfig()
	for name, value := range config.Serialize() {
		AddRowToStatusTable(statusTable, name, value)
	}
	statusTable.SortItems(0, core.Qt__AscendingOrder)

	receivedArea := widgets.NewQTextEdit(nil)
	receivedArea.SetFixedHeight(75)
	receivedArea.SetReadOnly(true)

	go func() {
		for {
			buf := <-dataReceived
			receivedArea.SetHtml(string(BeautifyBits(buf)))
		}
	}()

	statusLayout.AddWidget(receivedArea)
	statusLayout.AddWidget(statusTable)

	statusGroup := widgets.NewQGroupBox2("Status table:", nil)
	statusGroup.SetLayout(statusLayout)
	return statusGroup
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

func BeautifyBits(data []byte) []byte {
	var result []byte
	for _, bit := range data {
		result = append(result, bit)
		if len(result) >= 8 {
			if string(result[len(result)-8:]) == serial.BitStuffingFlag {
				result = result[:len(result)-8]
				result = append(result, ConfigureBackgroundColor([]byte(serial.BitStuffingFlag), "red")...)
			}
		}
		if len(result) >= 46 {
			if bytes.Equal(result[len(result)-46:len(result)-1], ConfigureBackgroundColor([]byte(serial.BitStuffingFlag), "red")) {
				result = result[:len(result)-1]
				result = append(result, ConfigureBackgroundColor([]byte{serial.BitToStuff}, "blue")...)
			}
		}
	}
	return result
}

func CreateTransmitterBox(transmitter serial.Serial) *widgets.QGroupBox {
	transmitterTextEdit := widgets.NewQTextEdit(nil)
	transmitterTextEdit.SetPlaceholderText("Input text here")

	transmitterTextEdit.ConnectTextChanged(func() {
		if err := ValidateTextEdit(transmitterTextEdit); err != nil {
			return
		}
		text := transmitterTextEdit.ToPlainText()
		if err := transmitter.Write([]byte(text)); err != nil {
			ShowErrorMessage(err.Error())
		}
		time.Sleep(time.Millisecond * 10)
	})

	transmitterLayout := widgets.NewQGridLayout2()
	transmitterLayout.AddWidget(transmitterTextEdit)

	transmitterGroup := widgets.NewQGroupBox2("Transmitter:", nil)
	transmitterGroup.SetLayout(transmitterLayout)
	return transmitterGroup
}

func CreateReceiverBox(receiver serial.Serial) *widgets.QGroupBox {
	receiverTextEdit := widgets.NewQTextEdit(nil)
	receiverTextEdit.SetReadOnly(true)

	go func() {
		for {
			buf := make([]byte, receiver.GetConfig().MaxReadBuffer)
			if _, err := receiver.Read(buf); err != nil {
				continue
			}

			dataReceived <- buf

			result := string(BeautifyBits(buf))
			receiverTextEdit.SetHtml(result)
		}
	}()

	receiverLayout := widgets.NewQGridLayout2()
	receiverLayout.AddWidget(receiverTextEdit)

	receiverGroup := widgets.NewQGroupBox2("Receiver:", nil)
	receiverGroup.SetLayout(receiverLayout)
	return receiverGroup
}

func InitGUI(transmitter serial.Serial, receiver serial.Serial) *widgets.QGridLayout {
	layout := widgets.NewQGridLayout2()
	layout.AddWidget2(CreateReceiverBox(receiver), 1, 0, 0)
	layout.AddWidget2(CreateTransmitterBox(transmitter), 0, 0, 0)
	layout.AddWidget2(CreateStatusBox(transmitter, receiver), 2, 0, 0)
	return layout
}
