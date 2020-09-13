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

	statusLabel := widgets.NewQLabel2("OK", nil, 0)

	go func() {
		for {
			buf := make([]byte, 4096)
			n, err := receiver.Read(buf)
			if err == serial.ErrorChecksumConfirmation {
				statusLabel.SetText("Error: checksum confirmation error")
			}
			if err != nil {
				continue
			}
			statusLabel.SetText("OK")
			receiverTextEdit.SetText(string(buf[:n-4]))
		}
	}()

	receiverLayout := widgets.NewQGridLayout2()
	receiverLayout.AddWidget(receiverTextEdit)
	receiverLayout.AddWidget(statusLabel)

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
