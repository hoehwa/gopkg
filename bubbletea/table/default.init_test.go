package table

import (
	"testing"
)

func TestInit(t *testing.T) {
	colData := Columns{
		{Title: "#", Width: 4},
		{Title: "Title", Width: 10},
		{Title: "Description", Width: 4},
	}

	rowData := Rows{
		{"1", "Title 1", "Description 1"},
		{"2", "Title 2", "Description 2"},
		{"3", "Title 3", "Description 3"},
		{"4", "Title 4", "Description 4"},
	}

	Init(colData, rowData)
}
