// Copyright 2011 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"math/rand"
	"strings"
	"time"
)

import (
	"github.com/lxn/walk"
)

type Foo struct {
	Bar     string
	Baz     float64
	Quux    time.Time
	checked bool
}

type FooModel struct {
	walk.TableModelBase
	items []*Foo
}

// Make sure we implement all required interfaces.
var _ walk.TableModel = &FooModel{}
var _ walk.ItemChecker = &FooModel{}

func NewFooModel() *FooModel {
	m := new(FooModel)
	m.ResetRows()
	return m
}

// Called by the TableView from SetModel to retrieve column information. 
func (m *FooModel) Columns() []walk.TableColumn {
	return []walk.TableColumn{
		{Title: "#"},
		{Title: "Bar"},
		{Title: "Baz", Format: "%.2f", Alignment: walk.AlignFar},
		{Title: "Quux", Format: "2006-01-02 15:04:05", Width: 150},
	}
}

// Called by the TableView from SetModel and every time the model publishes a
// RowsReset event.
func (m *FooModel) RowCount() int {
	return len(m.items)
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *FooModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return row

	case 1:
		return item.Bar

	case 2:
		return item.Baz

	case 3:
		return item.Quux
	}

	panic("unexpected col")
}

// Called by the TableView to retrieve if a given row is checked.
func (m *FooModel) Checked(row int) bool {
	return m.items[row].checked
}

// Called by the TableView when the user toggled the check box of a given row.
func (m *FooModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked

	return nil
}

func (m *FooModel) ResetRows() {
	// Create some random data.
	m.items = make([]*Foo, rand.Intn(50000))

	now := time.Now()

	for i := range m.items {
		m.items[i] = &Foo{
			Bar:  strings.Repeat("*", rand.Intn(5)+1),
			Baz:  rand.Float64() * 1000,
			Quux: time.Unix(rand.Int63n(now.Unix()), 0),
		}
	}

	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()
}

type MainWindow struct {
	*walk.MainWindow
	model *FooModel
}

func main() {
	walk.Initialize(walk.InitParams{PanicOnError: true})
	defer walk.Shutdown()

	rand.Seed(time.Now().UnixNano())

	mainWnd, _ := walk.NewMainWindow()

	mw := &MainWindow{
		MainWindow: mainWnd,
		model:      NewFooModel(),
	}

	mw.SetLayout(walk.NewVBoxLayout())
	mw.SetTitle("Walk TableView Example")

	resetRowsButton, _ := walk.NewPushButton(mw)
	resetRowsButton.SetText("Reset Rows")

	resetRowsButton.Clicked().Attach(func() {
		// Get some fresh data.
		mw.model.ResetRows()
	})

	tableView, _ := walk.NewTableView(mw)

	// Everybody loves check boxes.
	tableView.SetCheckBoxes(true)

	// Don't forget to set the model.
	tableView.SetModel(mw.model)

	mw.SetMinMaxSize(walk.Size{320, 240}, walk.Size{})
	mw.SetSize(walk.Size{800, 600})
	mw.Show()

	mw.Run()
}
