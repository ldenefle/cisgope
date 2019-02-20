package gui

import (
	"github.com/gdamore/tcell"
	cscope "github.com/ldenefle/cisgope/cscope"
	"github.com/rivo/tview"
	log "github.com/sirupsen/logrus"
)

func fillTable(table *tview.Table, symbols []cscope.Symbol) {
	for row, symbol := range symbols {
		for column, cell := range symbol.Serialize() {
			color := tcell.ColorWhite
			align := tview.AlignLeft
			if column == 0 {
				align = tview.AlignRight
				color = tcell.ColorDarkCyan
			}
			tableCell := tview.NewTableCell(cell).
				SetTextColor(color).
				SetAlign(align).
				SetSelectable(true).
				SetReference(symbol)
			table.SetCell(row, column, tableCell)
		}
	}
}

func Display(db cscope.Cscope) {
	var app = tview.NewApplication()
	callersTable := tview.NewTable().SetSelectable(true, false)
	calleesTable := tview.NewTable().SetSelectable(true, false)
	var searchField *tview.InputField

	var refresh = func(symbol string) {
		callersTable.Clear()
		calleesTable.Clear()
		callers, _ := db.Cmd(3, symbol)
		fillTable(callersTable, callers)
		callees, _ := db.Cmd(2, symbol)
		fillTable(calleesTable, callees)
	}
	log.Info("Display is starting")
	callersTable.SetDoneFunc(func(key tcell.Key) {
		app.SetFocus(calleesTable)
	}).SetSelectedFunc(func(row, column int) {
		symbol := callersTable.GetCell(row, column).GetReference().(cscope.Symbol)
		refresh(symbol.Name)
	})
	calleesTable.SetDoneFunc(func(key tcell.Key) {
		app.SetFocus(callersTable)
	}).SetSelectedFunc(func(row, column int) {
		symbol := calleesTable.GetCell(row, column).GetReference().(cscope.Symbol)
		refresh(symbol.Name)
	})
	callersTable.SetBorder(true).SetTitle("Callers")
	calleesTable.SetBorder(true).SetTitle("Callees")
	searchField = tview.NewInputField().
		SetLabel("Search").
		SetLabelWidth(80).
		SetDoneFunc(func(key tcell.Key) {
			refresh(searchField.GetText())
			app.SetFocus(callersTable)
		})

	searchBar := tview.NewForm().
		AddFormItem(searchField)

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(50, 50).
		SetBorders(true).
		SetBordersColor(tcell.ColorNavy)

	// Layout for screens wider than 100 cells.
	grid.AddItem(callersTable, 1, 0, 1, 1, 0, 0, false).
		AddItem(calleesTable, 1, 1, 1, 1, 0, 0, false).
		AddItem(searchBar, 0, 0, 1, 2, 0, 0, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch key := event.Key(); key {
		case tcell.KeyRune:
			ch := event.Rune()
			switch ch {
			case 's':
				app.SetFocus(searchField)
			}
		}
		return event
	})
	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
