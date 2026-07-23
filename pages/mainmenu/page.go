package mainmenu

import (
	"xerror/styles"
	"xerror/pages"
	tea "github.com/charmbracelet/bubbletea"
)

type OpenExplorerMsg struct{}

type Model struct {
	items  []string
	cursor int
}

func New() pages.Page {
	return Model{
		items: []string{
			"Open Explorer",
			"Quit",
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (pages.Page, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}

		case "enter":

			switch m.cursor {

			case 0:
				return m, func() tea.Msg {
					return OpenExplorerMsg{}
				}

			case 1:
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m Model) View(width, height int) string {

	s := styles.Title.Render("Main Menu") + "\n\n"

	for i, item := range m.items {

		if i == m.cursor {
			s += styles.Selected.Render("> "+item) + "\n"
		} else {
			s += "  " + item + "\n"
		}
	}

	return styles.Panel.Width(40).Render(s)
}