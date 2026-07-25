package styles

import "charm.land/lipgloss/v2"

func Panel(border bool) lipgloss.Style {
	padding := 0
	if border {
		padding = 1
	}
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(padding)
}

var Selected = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#fffb00")).
	Bold(true)

var Title = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("39"))
