package styles

import "github.com/charmbracelet/lipgloss"

var Panel = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Padding(1)

var Selected = lipgloss.NewStyle().
	Foreground(lipgloss.Color("205")).
	Bold(true)

var Title = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("39"))