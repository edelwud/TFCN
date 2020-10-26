package gui

import (
	"github.com/therecipe/qt/widgets"
)

func AddRowToStatusTable(table *widgets.QTreeWidget, option string, value string) {
	wItem := widgets.NewQTreeWidgetItem3(table, 0)

	optionEdit := widgets.NewQTextEdit(nil)
	optionEdit.SetReadOnly(true)
	optionEdit.SetFixedHeight(25)
	optionEdit.SetMinimumHeight(25)
	optionEdit.VerticalScrollBar().SetVisible(false)
	optionEdit.VerticalScrollBar().SetDisabled(true)
	optionEdit.SetHtml(option)
	optionEdit.SetStyleSheet("border: none")
	table.SetItemWidget(wItem, 0, optionEdit)

	valueEdit := widgets.NewQTextEdit(nil)
	valueEdit.SetReadOnly(true)
	valueEdit.SetFixedHeight(25)
	valueEdit.SetMinimumHeight(25)
	valueEdit.VerticalScrollBar().SetVisible(false)
	valueEdit.VerticalScrollBar().SetDisabled(true)
	valueEdit.SetHtml(value)
	valueEdit.SetStyleSheet("border: none")
	table.SetItemWidget(wItem, 1, valueEdit)
}

func CreateStatusTable() *widgets.QTreeWidget {
	statusTable := widgets.NewQTreeWidget(nil)
	statusTable.SetColumnCount(2)
	statusTable.SetHeaderLabels([]string{"Parsed code", "Hamming code"})
	statusTable.SetIndentation(5)
	statusTable.SetColumnWidth(0, 180)
	return statusTable
}
