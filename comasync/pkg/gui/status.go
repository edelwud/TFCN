package gui

import (
	"github.com/therecipe/qt/widgets"
)

func CreateStatusTable() *widgets.QTableWidget {
	statusTable := widgets.NewQTableWidget(nil)
	statusTable.SetColumnCount(2)
	statusTable.SetHorizontalHeaderLabels([]string{"Option name", "Option value"})
	return statusTable
}
