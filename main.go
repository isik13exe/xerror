package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"xerror/model"

	tea "charm.land/bubbletea/v2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: xerror <dir>")
		os.Exit(1)
	}

	path := os.Args[1]
	name := filepath.Base(path)

	p := tea.NewProgram(model.New(name))

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
