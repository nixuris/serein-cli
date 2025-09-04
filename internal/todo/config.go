package todo

import (
	"encoding/json"
	"os"
	"fmt"
	"path/filepath"
)

// Configuration and persistence

func (m *Model) LoadConfig() {
    if err := os.MkdirAll(m.ConfigPath, 0755); err != nil {
        fmt.Println("Error creating config directory:", err)
        m.CreateDefaultConfig()
        return
    }

    configFile := filepath.Join(m.ConfigPath, "config.json")

    data, err := os.ReadFile(configFile)
    if err != nil {
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
        fmt.Println("Error marshaling config:", err)
        return
    }

    if err := os.WriteFile(configFile, data, 0644); err != nil {
        fmt.Println("Error saving config file:", err)
    }
}

func (m *Model) CreateDefaultConfig() {
    m.Tasks = []Task{
        {ID: 1, Task: "Welcome to your todo app!", Checked: false, Context: "Work"},
        {ID: 2, Task: "Press 'a' to add a new task", Checked: false, Context: "Work"},
        {ID: 3, Task: "Press space to toggle completion", Checked: true, Context: "Personal"},
        {ID: 4, Task: "Use arrow keys to navigate", Checked: false, Context: "Personal"},
    }
    m.Contexts = []string{"Work", "Personal"}
    m.NextID = 5
}
