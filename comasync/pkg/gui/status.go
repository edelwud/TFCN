package gui

import (
	"github.com/therecipe/qt/widgets"
)

func AddRowToStatusTable(table *widgets.QTreeWidget, option string, value string) {
	table.AddTopLevelItem(widgets.NewQTreeWidgetItem2([]string{option, value}, 0))
}

func CreateStatusTable() *widgets.QTreeWidget {
	statusTable := widgets.NewQTreeWidget(nil)
	statusTable.SetColumnCount(2)
	statusTable.SetHeaderLabels([]string{"Option name", "Option value"})
	statusTable.SetIndentation(0)
	return statusTable
}
