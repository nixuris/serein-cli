package todo

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Configuration and persistence

func (m *Model) LoadConfig() {
	// Ensure config directory exists
	os.MkdirAll(m.ConfigPath, 0755)
	
	configFile := filepath.Join(m.ConfigPath, "config.json")
	
	// Try to load existing config
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		// Create default config
		m.CreateDefaultConfig()
		return
	}

	var config struct {
		Tasks    []Task   `json:"tasks"`
		NextID   int      `json:"next_id"`
		Contexts []string `json:"contexts"`
	}

	if err := json.Unmarshal(data, &config); err != nil {
		m.CreateDefaultConfig()
		return
	}

	m.Tasks = config.Tasks
	m.NextID = config.NextID
	m.Contexts = config.Contexts
	
	// Ensure we have a valid next ID
	if m.NextID == 0 {
		maxID := 0
		for _, task := range m.Tasks {
			if task.ID > maxID {
				maxID = task.ID
			}
		}
		m.NextID = maxID + 1
	}
}

func (m *Model) CreateDefaultConfig() {
	m.Tasks = []Task{
		{ID: 1, Task: "Welcome to your todo app!", Checked: false, Context: "Work"},
		{ID: 2, Task: "Press 'a' to add a new task", Checked: false, Context: "Work"},
		{ID: 3, Task: "Press space to toggle completion", Checked: true, Context: "Personal"},
		{ID: 4, Task: "Use arrow keys to navigate", Checked: false, Context: "Personal"},
	}
	m.NextID = 5
}

func (m *Model) SaveConfig() {
	configFile := filepath.Join(m.ConfigPath, "config.json")
	
	config := struct {
		Tasks    []Task   `json:"tasks"`
		NextID   int      `json:"next_id"`
		Contexts []string `json:"contexts"`
	}{
		Tasks:    m.Tasks,
		NextID:   m.NextID,
		Contexts: m.Contexts,
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return
	}

	ioutil.WriteFile(configFile, data, 0644)
}
