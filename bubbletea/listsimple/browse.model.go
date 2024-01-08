package listsimple

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pkg/browser"
)

type browseModel struct {
	list     list.Model
	choice   string
	quitting bool
}

func (bm browseModel) Init() tea.Cmd {
	return nil
}

func (bm browseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		bm.list.SetWidth(msg.Width)
		return bm, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			bm.quitting = true
			return bm, tea.Quit

		case "enter":
			i, ok := bm.list.SelectedItem().(item)
			if ok {
				bm.choice = string(i)
			}

			githubTree := strings.Split(bm.choice, ",")[0]
			url := fmt.Sprintf("https://github.com/%s", githubTree)
			browser.OpenURL(url)

			return bm, tea.Quit
		}
	}

	var cmd tea.Cmd
	bm.list, cmd = bm.list.Update(msg)
	return bm, cmd
}

func (bm browseModel) View() string {
	if bm.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("%s", bm.choice))
	}
	if bm.quitting {
		return quitTextStyle.Render(`
		You can explore trending repos in the web browser.
		https://github.com/trending`)
	}
	return "\n" + bm.list.View()
}
