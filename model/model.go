package model

import (
	"xerror/pages"
	"xerror/pages/explorer"
	"xerror/pages/mainmenu"

	tea "charm.land/bubbletea/v2"
)

type Model struct {
	Page   pages.Page
	Dir    string
	Width  int
	Height int
}

func New(dir string) Model {
	return Model{
		Page: mainmenu.New(),
		Dir:  dir,
	}
}

func (m Model) Init() tea.Cmd {
	return m.Page.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

	case mainmenu.OpenExplorerMsg:
		m.Page = explorer.New(m.Dir)

	case explorer.BackMsg:
		m.Page = mainmenu.New()
	}

	var cmd tea.Cmd
	m.Page, cmd = m.Page.Update(msg)
	return m, cmd
}

func (m Model) View() tea.View {
	content := m.Page.View(m.Width, m.Height)
	v := tea.NewView(content)
	v.AltScreen = true
	return v
}