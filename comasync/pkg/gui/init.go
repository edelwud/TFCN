package gui

import (
	"../com"
	"github.com/therecipe/qt/widgets"
)

func CreateStatusBox() *widgets.QGroupBox {
	statusLayout := widgets.NewQGridLayout2()
	statusTable := CreateStatusTable()

	AddRowToStatusTable(statusTable, "hey", "world")
	statusLayout.AddWidget(statusTable)

	statusGroup := widgets.NewQGroupBox2("Status table:", nil)
	statusGroup.SetLayout(statusLayout)
	return statusGroup
}

func CreateTransmitterBox(transmitter com.Port) *widgets.QGroupBox {
	transmitterTextEdit := widgets.NewQTextEdit(nil)
	transmitterTextEdit.SetPlaceholderText("Input text here")
	transmitterTextEdit.ConnectTextChanged(func() {
		if _, err := transmitter.Write([]byte("1" + transmitterTextEdit.ToPlainText())); err != nil {
			ShowErrorMessage(err.Error())
		}
	})

	transmitterLayout := widgets.NewQGridLayout2()
	transmitterLayout.AddWidget(transmitterTextEdit)

	transmitterGroup := widgets.NewQGroupBox2("Transmitter:", nil)
	transmitterGroup.SetLayout(transmitterLayout)
	return transmitterGroup
}

func CreateReceiverBox(receiver com.Port) *widgets.QGroupBox {
	receiverTextEdit := widgets.NewQTextEdit(nil)
	receiverTextEdit.SetReadOnly(true)
	go func() {
		for {
			buf := make([]byte, 2048)
			if _, err := receiver.Read(buf); err != nil {
				continue
			}
			receiverTextEdit.SetText(string(buf[1:]))
		}
	}()

	receiverLayout := widgets.NewQGridLayout2()
	receiverLayout.AddWidget(receiverTextEdit)

	receiverGroup := widgets.NewQGroupBox2("Receiver:", nil)
	receiverGroup.SetLayout(receiverLayout)
	return receiverGroup
}

func InitGUI(transmitter com.Port, receiver com.Port) *widgets.QGridLayout {
	layout := widgets.NewQGridLayout2()
	layout.AddWidget2(CreateReceiverBox(receiver), 1, 0, 0)
	layout.AddWidget2(CreateTransmitterBox(transmitter), 0, 0, 0)
	layout.AddWidget2(CreateStatusBox(), 2, 0, 0)
	return layout
}
