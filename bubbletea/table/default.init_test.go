package table

import (
	"testing"

	"github.com/charmbracelet/bubbles/table"
)

func TestInit(t *testing.T) {
	colData := []table.Column{
		{Title: "#", Width: 4},
		{Title: "Title", Width: 10},
		{Title: "Description", Width: 4},
	}

	rowData := []table.Row{
		{"1", "Title 1", "Description 1"},
		{"2", "Title 2", "Description 2"},
		{"3", "Title 3", "Description 3"},
		{"4", "Title 4", "Description 4"},
	}

	Init(colData, rowData)
}
