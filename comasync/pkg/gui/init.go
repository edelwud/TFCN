package gui

import (
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

func CreateTransmitterBox() *widgets.QGroupBox {
	transmitterTextEdit := widgets.NewQTextEdit(nil)
	transmitterTextEdit.SetPlaceholderText("Input text here")

	transmitterLayout := widgets.NewQGridLayout2()
	transmitterLayout.AddWidget(transmitterTextEdit)

	transmitterGroup := widgets.NewQGroupBox2("Transmitter:", nil)
	transmitterGroup.SetLayout(transmitterLayout)
	return transmitterGroup
}

func CreateReceiverBox() *widgets.QGroupBox {
	receiverTextEdit := widgets.NewQTextEdit(nil)
	receiverTextEdit.SetReadOnly(true)

	receiverLayout := widgets.NewQGridLayout2()
	receiverLayout.AddWidget(receiverTextEdit)

	receiverGroup := widgets.NewQGroupBox2("Receiver:", nil)
	receiverGroup.SetLayout(receiverLayout)
	return receiverGroup
}

func InitGUI() *widgets.QGridLayout {
	layout := widgets.NewQGridLayout2()
	layout.AddWidget2(CreateTransmitterBox(), 0, 0, 0)
	layout.AddWidget2(CreateReceiverBox(), 1, 0, 0)
	layout.AddWidget2(CreateStatusBox(), 2, 0, 0)
	return layout
}
