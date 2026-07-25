package pages

import tea "charm.land/bubbletea/v2"

type RedrawMsg struct{}

type Page interface {
	Init() tea.Cmd
	Update(tea.Msg) (Page, tea.Cmd)
	View(width, height int) string
}

func Redraw() tea.Msg {
	return RedrawMsg{}
}