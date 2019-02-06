package gui

import (
	"github.com/gdamore/tcell"
	cscope "github.com/ldenefle/cisgope/cscope"
	"github.com/rivo/tview"
	log "github.com/sirupsen/logrus"
)

func Display(db cscope.Cscope) {
	var callersList *tview.List
	var calleesList *tview.List

	var refresh = func(symbol string) {
		callersList.Clear()
		calleesList.Clear()
		callers, _ := db.Cmd(3, symbol)
		for _, symbol := range callers {
			callersList = callersList.InsertItem(0, symbol.String(), "", 0, nil)
		}
		callees, _ := db.Cmd(4, symbol)
		for _, symbol := range callees {
			calleesList = calleesList.InsertItem(0, symbol.String(), "", 0, nil)
		}
	}
	log.Info("Display is starting")
	callersList = tview.NewList()
	calleesList = tview.NewList()
	searchField := tview.NewInputField().
		SetLabel("Search").
		SetLabelWidth(80).
		SetChangedFunc(func(text string) {
			refresh(text)
		})

	searchBar := tview.NewForm().
		AddFormItem(searchField)

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(50, 50).
		SetBorders(true).
		SetBordersColor(tcell.ColorNavy)

	// Layout for screens wider than 100 cells.
	grid.AddItem(callersList, 1, 0, 1, 1, 0, 0, false).
		AddItem(calleesList, 1, 1, 1, 1, 0, 0, false).
		AddItem(searchBar, 0, 0, 1, 2, 0, 0, true)

	if err := tview.NewApplication().SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
