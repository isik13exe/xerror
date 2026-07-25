package explorer

import (
	"fmt"
	"strings"

	"xerror/pages"
	"xerror/pages/explorer/editor"
	"xerror/pages/explorer/exp"
	"xerror/styles"
	"os"
	"github.com/ionut-t/goeditor"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type BackMsg struct{}

type Focus int

const (
	Left Focus = iota
	Right
)

type Model struct {
	root    *exp.Node
	current *exp.Node
	cursor  int
	focus   Focus
	ed      editor.Model
}

func New(dirpath string) pages.Page {
	root, err := exp.BuildTree(dirpath)
	if err != nil {
		panic(err)
	}
	return Model{
		root:    root,
		current: root,
		focus:   Left,
		ed:      editor.New(40, 20),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (pages.Page, tea.Cmd) {
	var cmds []tea.Cmd

	// When editor has focus, forward keys to it
	if m.focus == Right {
		switch msg := msg.(type) {
			case goeditor.SaveMsg:
				path := m.ed.FilePath()
				if msg.Path != nil {
					path = *msg.Path
				}
				if path != "" {
					_ = os.WriteFile(path, []byte(msg.Content), 0o644)
				}
				return m, nil
		
			case goeditor.QuitMsg:
				m.ed.Clear()
				m.focus = Left
				return m, nil
			}
		var cmd tea.Cmd
		m.ed, cmd = m.ed.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
		if key, ok := msg.(tea.KeyPressMsg); ok && key.String() == "ctrl+w" {
			m.ed.Blur()
			m.focus = Left
			return m, tea.Batch(cmds...)
		}
		return m, tea.Batch(cmds...)
	}

	// Left panel (tree)
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		entries := m.current.Entries()

		switch msg.String() {
			
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(entries)-1 {
				m.cursor++
			}
		case "enter":
			if len(entries) == 0 {
				return m, nil
			}
			selected := entries[m.cursor]
			if selected.IsDir {
				m.current = selected
				m.cursor = 0
			} else {
				if err := m.ed.OpenFile(selected.Path); err != nil {
					return m, nil
				}
				m.focus = Right
				return m, func() tea.Msg { return pages.RedrawMsg{} }
			}
		case "esc", "backspace":
			if m.current.Parent != nil {
				m.current = m.current.Parent
				m.cursor = 0
				return m, nil
			}
			return m, func() tea.Msg { return BackMsg{} }
		case "tab":
			if m.ed.FilePath() != "" {
				m.ed.Focus()
				m.focus = Right
			}
		}
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View(width, height int) string {
	leftWidth := 34
	rightWidth := width - leftWidth - 4
	if rightWidth < 20 {
		rightWidth = 20
	}
	panelHeight := height - 4
	m.ed.SetSize(rightWidth-2, panelHeight-2)

	// Left: combined folders + files
	left := strings.Builder{}
	left.WriteString(fmt.Sprintf(" %s\n\n", m.current.Name))

	entries := m.current.Entries()
	if len(entries) == 0 {
		left.WriteString("  <empty>")
	} else {
		for i, entry := range entries {
			cursor := "  "
			if m.focus == Left && i == m.cursor {
				cursor = "==> "
			}
			icon := ""
			if entry.IsDir {
				icon = "[DIR] "
			}
			left.WriteString(cursor)
			left.WriteString(icon)
			left.WriteString(entry.Name)
			left.WriteByte('\n')
		}
	}

	rightContent := m.ed.View()

	leftStyle := styles.Panel(true).Width(leftWidth).Height(panelHeight)
	rightStyle := styles.Panel(true).Width(rightWidth).Height(panelHeight)

	if m.focus == Left {
		leftStyle = leftStyle.BorderForeground(lipgloss.Color("39"))
	} else {
		rightStyle = rightStyle.BorderForeground(lipgloss.Color("39"))
	}

	header := styles.Title.Render(currentPath(m.current))

	body := lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftStyle.Render(left.String()),
		rightStyle.Render(rightContent),
	)

	return lipgloss.JoinVertical(lipgloss.Left, header, "", body)
}

func currentPath(n *exp.Node) string {
	if n == nil {
		return "/"
	}
	var parts []string
	for n.Parent != nil {
		parts = append(parts, n.Name)
		n = n.Parent
	}
	if len(parts) == 0 {
		return "/"
	}
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}
	return "/" + strings.Join(parts, "/")
}
