package explorer

import (
	"xerror/pages"
	"xerror/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type BackMsg struct{}

type Focus int

const (
	Left Focus = iota
	Right
)

type Model struct {
	folders []string
	files   []string

	leftCursor  int
	rightCursor int

	focus Focus
}

func New() pages.Page {
	return Model{
		folders: []string{
			"src",
			"assets",
			"include",
			"tests",
		},

		files: []string{
			"main.go",
			"lexer.go",
			"parser.go",
			"README.md",
		},

		focus: Left,
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

			if m.focus == Left {
				if m.leftCursor > 0 {
					m.leftCursor--
				}
			} else {
				if m.rightCursor > 0 {
					m.rightCursor--
				}
			}

		case "down":

			if m.focus == Left {
				if m.leftCursor < len(m.folders)-1 {
					m.leftCursor++
				}
			} else {
				if m.rightCursor < len(m.files)-1 {
					m.rightCursor++
				}
			}

		case "enter":

			if m.focus == Left {
				m.focus = Right
			}

		case "esc":

			if m.focus == Right {
				m.focus = Left
				return m, nil
			}

			return m, func() tea.Msg {
				return BackMsg{}
			}
		}
	}

	return m, nil
}

func (m Model) View(width, height int) string {

	left := "Folders\n\n"

	for i, folder := range m.folders {

		cursor := "  "

		if m.focus == Left && i == m.leftCursor {
			cursor = "▶ "
		}

		left += cursor + "📁 " + folder + "\n"
	}

	right := "Files\n\n"

	for i, file := range m.files {

		cursor := "  "

		if m.focus == Right && i == m.rightCursor {
			cursor = "▶ "
		}

		right += cursor + "📄 " + file + "\n"
	}

	leftStyle := styles.Panel.
		Width(30).
		Height(height - 4)

	rightStyle := styles.Panel.
		Width(width - 36).
		Height(height - 4)

	if m.focus == Left {
		leftStyle = leftStyle.BorderForeground(lipgloss.Color("39"))
	} else {
		rightStyle = rightStyle.BorderForeground(lipgloss.Color("39"))
	}

	leftPanel := leftStyle.Render(left)
	rightPanel := rightStyle.Render(right)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftPanel,
		rightPanel,
	)
}