package table

import (
	"testing"

	"github.com/charmbracelet/bubbles/table"
)

func TestInit(t *testing.T) {
	colData := table.Column{Title: "Rank", Width: 4}
	rowData := []string{"1"}

	Init(colData, rowData)
}
