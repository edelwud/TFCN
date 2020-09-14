package gui

import (
	"../serial"
	"github.com/therecipe/qt/widgets"
	"time"
)

func CreateStatusBox(transmitter serial.Serial, _ serial.Serial) *widgets.QGroupBox {
	statusLayout := widgets.NewQGridLayout2()
	statusTable := CreateStatusTable()

	config := transmitter.GetConfig()
	for name, value := range config.Serialize() {
		AddRowToStatusTable(statusTable, name, value)
	}

	statusLayout.AddWidget(statusTable)

	statusGroup := widgets.NewQGroupBox2("Status table:", nil)
	statusGroup.SetLayout(statusLayout)
	return statusGroup
}

func CreateTransmitterBox(transmitter serial.Serial) *widgets.QGroupBox {
	transmitterTextEdit := widgets.NewQTextEdit(nil)
	transmitterTextEdit.SetPlaceholderText("Input text here")

	go transmitterTextEdit.ConnectTextChanged(func() {
		text := " " + transmitterTextEdit.ToPlainText()
		if err := transmitter.Write([]byte(text)); err != nil {
			ShowErrorMessage(err.Error())
		}
		time.Sleep(time.Millisecond * 15)
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
			receiverTextEdit.SetText(string(buf[1:]))
			receiverTextEdit.VerticalScrollBar().SetValue(
				receiverTextEdit.VerticalScrollBar().Maximum())
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
