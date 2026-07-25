package editor

import (
	"os"
	"xerror/tools"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/ionut-t/goeditor"
)

// Model wraps goeditor for use as the right-hand panel.
type Model struct {
	ed       goeditor.Model
	filePath string
	width    int
	height   int
	focused  bool
}

func New(width, height int) Model {
	ed := goeditor.New(width, height)
	// Default to Vim mode. Call DisableVimMode(true) for simple mode.
	ed.DisableVimMode(false)
	return Model{
		ed:     ed,
		width:  width,
		height: height,
	}
}

func (m *Model) OpenFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	m.filePath = path
	lang := tools.Lang(path)
	if lang != "" {
		m.ed.SetLanguage(lang, "monokai")
	} else {
		m.ed.SetLanguage("", "")
	}
	m.ed.SetBytes(data)
	m.ed.Focus()
	m.focused = true
	return nil
}

func (m *Model) SetSize(width, height int) {
	m.width = width
	m.height = height
	m.ed.SetSize(width, height)
}

func (m *Model) Focus() {
	m.ed.Focus()
	m.focused = true
}

func (m *Model) Blur() {
	m.ed.Blur()
	m.focused = false
}

func (m Model) IsFocused() bool  { return m.focused }
func (m Model) FilePath() string { return m.filePath }

func (m Model) Init() tea.Cmd {
	return m.ed.CursorBlink()
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if !m.focused {
		return m, nil
	}

	var cmd tea.Cmd
	m.ed, cmd = m.ed.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.filePath == "" {
		style := lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Width(m.width).
			Height(m.height).
			Align(lipgloss.Center, lipgloss.Center)
		return style.Render(tools.Logo() + "\nSelect a file and press Enter\nto open it here")
	}
	return m.ed.View()
}

func (m *Model) Clear() {
	m.filePath = ""
	m.ed.SetContent("") // or SetBytes(nil) depending on the API
	m.ed.Blur()
	m.focused = false
}
