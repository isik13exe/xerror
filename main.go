package main

import (
	"log"

	"xerror/model"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(
		model.New(),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}