package todo

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
)

// Initialize creates a new model
func Initialize() Model {
	homeDir, _ := os.UserHomeDir()
	configPath := filepath.Join(homeDir, ".config", "tuido")

	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 200
	ti.Width = 50

	dateInputs := make([]textinput.Model, 3)
	for i := range dateInputs {
		dateInputs[i] = textinput.New()
		dateInputs[i].Focus()
		dateInputs[i].CharLimit = 4
		dateInputs[i].Width = 10
	}

	m := Model{
		TextInput:      ti,
		DateInputs:     dateInputs,
		KeyMap:         DefaultKeyMap(),
		Help:           help.New(),
		ConfigPath:     configPath,
		MaxHistory:     50,
		ViewMode:       NormalView,
	}

	m.LoadConfig()
	m.UpdateContexts()

	return m
}

// Init implements tea.Model
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}
