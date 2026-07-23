package model

import (
	"xerror/pages/explorer"
	"xerror/pages/mainmenu"
	"xerror/pages"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Page   pages.Page
	Width  int
	Height int
}

func New() Model {
	return Model{
		Page: mainmenu.New(),
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
		m.Page = explorer.New()

	case explorer.BackMsg:
		m.Page = mainmenu.New()
	}

	var cmd tea.Cmd
	m.Page, cmd = m.Page.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	return m.Page.View(m.Width, m.Height)
}