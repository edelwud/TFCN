package gui

import (
	"../serial"
	"bytes"
	"regexp"
	"time"

	"github.com/pkg/errors"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
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

			buf = BeautifyFlagBits(buf)
			buf = BeautifyBitStuffed(buf)

			receivedArea.SetHtml(string(buf))
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

func BeautifyFlagBits(data []byte) []byte {
	var result []byte
	for _, bit := range data {
		result = append(result, bit)
		if len(result) >= 8 {
			if string(result[len(result)-8:]) == serial.BitStuffingFlag {
				result = result[:len(result)-8]
				result = append(result, ConfigureBackgroundColor([]byte(serial.BitStuffingFlag), "yellow")...)
			}
		}
	}
	return result
}

func BeautifyBitStuffed(data []byte) []byte {
	var result []byte
	for _, bit := range data {
		result = append(result, bit)
		if len(result) >= 49 {
			if bytes.Equal(result[len(result)-49:len(result)-1],
				ConfigureBackgroundColor([]byte(serial.BitStuffingFlag), "yellow")) {
				result = result[:len(result)-1]
				result = append(result, ConfigureBackgroundColor([]byte{serial.BitToStuff}, "cyan")...)
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
			packet, err := receiver.Read()
			if err != nil {
				continue
			}

			dataReceived <- packet.Data
			packet.DeBitStuffing()

			buf := BeautifyFlagBits(packet.Data)

			receiverTextEdit.SetHtml(string(buf))
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
