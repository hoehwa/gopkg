package listfancy

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

// TODO: convert it to class method.
func InitCallout(callout Callout) {
	rand.Seed(time.Now().UTC().UnixNano())

	if _, err := tea.NewProgram(newModel(callout)).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
