package mainmenu

import (
	"xerror/pages"
	"xerror/styles"
//	"xerror/tools"
	"xerror/assets"
	tea "charm.land/bubbletea/v2"
)

type OpenExplorerMsg struct{}

type Model struct {
	items  []item
	cursor int
}

type item struct {
	tag      string
	normal   string
	selected string
}

var quit = item{tag: "quit", normal: "┌────────────────┐\n│      Quit      │\n└────────────────┘", selected: "┌────────────────┐\n│      Quit      │<==\n└────────────────┘"}
var explorer = item{tag: "explore", normal: "┌────────────────┐\n│    Explorer    │\n└────────────────┘", selected: "┌────────────────┐\n│    Explorer    │<==\n└────────────────┘"}

func New() pages.Page {
	return Model{
		items: []item{
			explorer,
			quit,
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (pages.Page, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case "enter":
			switch m.cursor {
			case 0:
				return m, func() tea.Msg { return OpenExplorerMsg{} }
			case 1:
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m Model) View(width, height int) string {
	s := assets.Logo
	for i, item := range m.items {
		if i == m.cursor {
			s += styles.Selected.Render(item.selected) + "\n"
		} else {
			s += item.normal + "\n"
		}
	}
	return styles.Panel(true).Width(90).Render(s)
}

func (i item) selecte() string {
	return i.selected
}
