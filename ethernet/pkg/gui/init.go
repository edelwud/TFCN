package gui

import (
	"../ethernet"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"sync"
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

func CreateTransmitterBox(receiver *widgets.QTextEdit, status *widgets.QTreeWidget) *widgets.QGroupBox {
	transmitterTextEdit := widgets.NewQTextEdit(nil)
	transmitterTextEdit.SetPlaceholderText("Input text here")

	mutex := sync.Mutex{}

	sendButton := widgets.NewQPushButton2("Send", nil)
	sendButton.ConnectClicked(func(_ bool) {
		go func() {
			mutex.Lock()
			text := transmitterTextEdit.ToPlainText()

			receiver.SetText("")
			status.Clear()

			for _, symbol := range text {
				currentCollisionNumber := 0
				transferred := true

				for ethernet.CheckChannelStatus() != nil {
					time.Sleep(100 * time.Millisecond)
				}

				ethernet.SendSymbol(symbol)
				for ethernet.CollisionWindow() != nil {
					if currentCollisionNumber == ethernet.MaxCollisionNumber {
						transferred = false
						break
					}
					currentCollisionNumber++
					ethernet.SendSymbol(symbol)
				}

				for ethernet.CheckChannelStatus() != nil {
					time.Sleep(100 * time.Millisecond)
				}

				if transferred {
					collisionString := ""
					for index := 0; index < currentCollisionNumber; index++ {
						collisionString += "* "
					}
					receivedText := receiver.ToPlainText()

					if symbol == '\t' {
						AddRowToStatusTable(status, "\\t", collisionString)
					} else if symbol == '\n' {
						AddRowToStatusTable(status, "\\n", collisionString)
					} else {
						AddRowToStatusTable(status, string(symbol), collisionString)
					}

					receivedText += string(symbol)
					receiver.SetText(receivedText)
				}
			}
			mutex.Unlock()
		}()
	})

	transmitterLayout := widgets.NewQGridLayout2()
	transmitterLayout.AddWidget(transmitterTextEdit)
	transmitterLayout.AddWidget(sendButton)

	transmitterGroup := widgets.NewQGroupBox2("Transmitter:", nil)
	transmitterGroup.SetLayout(transmitterLayout)
	return transmitterGroup
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

	layout.AddWidget2(transmitter, 0, 0, 0)
	layout.AddWidget2(receiver, 1, 0, 0)
	layout.AddWidget2(status, 2, 0, 0)
	return layout
}
