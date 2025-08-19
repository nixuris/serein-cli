package todo

import (
	"fmt"
	"strings"
	"time"
	"strconv"
	"sort"
)

// Helper methods

func (m *Model) ShowInputDialog(mode InputMode, prompt string) {
	m.ViewMode = InputView // Corrected: Assign InputView constant
	m.InputMode = mode     // Corrected: Assign mode to InputMode
	m.InputPrompt = prompt
	m.TextInput.SetValue("")
	m.TextInput.Focus()
}

func (m *Model) ShowDateInputDialog() {
	m.ViewMode = DateInputView
	m.DateInputIndex = 0
	now := time.Now()
	m.DateInputs[0].SetValue(fmt.Sprintf("%02d", now.Day()))
	m.DateInputs[1].SetValue(fmt.Sprintf("%02d", now.Month()))
	m.DateInputs[2].SetValue(fmt.Sprintf("%d", now.Year()))
	for i := range m.DateInputs {
		m.DateInputs[i].Focus()
	}
}

func (m *Model) ShowRemoveTagDialog() {
	task := m.GetCurrentTask()
	if len(task.Tags) == 0 {
		m.ErrorMessage = "No tags to remove"
		return
	}
	m.ViewMode = RemoveTagView
	m.RemoveTagIndex = 0
	m.RemoveTagChecks = make([]bool, len(task.Tags))
}

func (m *Model) GetFilteredTasks() []Task {
	return m.GetTasksForContext(m.CurrentContext)
}

func (m *Model) GetTasksForContext(context string) []Task {
	var filtered []Task
	for _, task := range m.Tasks {
		if task.Context == context {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

func (m *Model) GetCurrentTask() Task {
	tasks := m.GetFilteredTasks()
	if len(tasks) == 0 || m.SelectedIndex >= len(tasks) {
		return Task{}
	}
	return tasks[m.SelectedIndex]
}

func (m *Model) MoveUp() {
	tasks := m.GetFilteredTasks()
	if len(tasks) > 0 {
		m.SelectedIndex = (m.SelectedIndex - 1 + len(tasks)) % len(tasks)
	}
}

func (m *Model) MoveDown() {
	tasks := m.GetFilteredTasks()
	if len(tasks) > 0 {
		m.SelectedIndex = (m.SelectedIndex + 1) % len(tasks)
	}
}

func (m *Model) MoveTaskUp() {
	tasks := m.GetFilteredTasks()
	if m.SelectedIndex > 0 {
		taskToMove := tasks[m.SelectedIndex]
		for i := range m.Tasks {
			if m.Tasks[i].ID == taskToMove.ID {
				m.Tasks[i], m.Tasks[i-1] = m.Tasks[i-1], m.Tasks[i]
				break
			}
		}
		m.SelectedIndex--
	}
}

func (m *Model) MoveTaskDown() {
	tasks := m.GetFilteredTasks()
	if m.SelectedIndex < len(tasks)-1 {
		taskToMove := tasks[m.SelectedIndex]
		for i := range m.Tasks {
			if m.Tasks[i].ID == taskToMove.ID {
				m.Tasks[i], m.Tasks[i+1] = m.Tasks[i+1], m.Tasks[i]
				break
			}
		}
		m.SelectedIndex++
	}
}

func (m *Model) NextContext() {
	if len(m.Contexts) > 0 {
		currentIdx := m.FindContextIndex(m.CurrentContext)
		nextIdx := (currentIdx + 1) % len(m.Contexts)
		m.CurrentContext = m.Contexts[nextIdx]
		m.SelectedIndex = 0
	}
}

func (m *Model) PreviousContext() {
	if len(m.Contexts) > 0 {
		currentIdx := m.FindContextIndex(m.CurrentContext)
		prevIdx := (currentIdx - 1 + len(m.Contexts)) % len(m.Contexts)
		m.CurrentContext = m.Contexts[prevIdx]
		m.SelectedIndex = 0
	}
}

func (m *Model) FindContextIndex(context string) int {
	for i, ctx := range m.Contexts {
		if ctx == context {
			return i
		}
	}
	return 0
}

func (m *Model) ToggleCurrentTask() {
	tasks := m.GetFilteredTasks()
	if len(tasks) == 0 {
		return
	}

	currentTask := tasks[m.SelectedIndex]
	for i := range m.Tasks {
		if m.Tasks[i].ID == currentTask.ID {
			m.Tasks[i].Checked = !m.Tasks[i].Checked
			break
		}
	}
}

func (m *Model) AddTask(taskText string) {
	newTask := Task{
		ID:      m.NextID,
		Task:    taskText,
		Checked: false,
		Context: m.CurrentContext,
	}
	m.Tasks = append(m.Tasks, newTask)
	m.NextID++
	
	// Move selection to new task
	filtered := m.GetFilteredTasks()
	m.SelectedIndex = len(filtered) - 1
}

func (m *Model) EditCurrentTask(newText string) {
	tasks := m.GetFilteredTasks()
	if len(tasks) == 0 {
		return
	}

	currentTask := tasks[m.SelectedIndex]
	for i := range m.Tasks {
		if m.Tasks[i].ID == currentTask.ID {
			m.Tasks[i].Task = newText
			break
		}
	}
}

func (m *Model) DeleteCurrentTask() {
	tasks := m.GetFilteredTasks()
	if len(tasks) == 0 {
		return
	}

	currentTask := tasks[m.SelectedIndex]
	for i := range m.Tasks {
		if m.Tasks[i].ID == currentTask.ID {
			m.Tasks = append(m.Tasks[:i], m.Tasks[i+1:]...)
			break
		}
	}

	// Adjust selection
	newTasks := m.GetFilteredTasks()
	if m.SelectedIndex >= len(newTasks) && len(newTasks) > 0 {
		m.SelectedIndex = len(newTasks) - 1
	}
}

func (m *Model) AddContext(contextName string) {
	// Check if context already exists
	for _, ctx := range m.Contexts {
		if ctx == contextName {
			m.ErrorMessage = "Context already exists"
			return
		}
	}

	m.Contexts = append(m.Contexts, contextName)
	m.CurrentContext = contextName
	m.SelectedIndex = 0
}

func (m *Model) RenameContext(newName string) {
	if newName == m.CurrentContext {
		return
	}

	// Check if new name already exists
	for _, ctx := range m.Contexts {
		if ctx == newName {
			m.ErrorMessage = "Context name already exists"
			return
		}
	}

	oldName := m.CurrentContext

	// Update context in Contexts list
	for i, ctx := range m.Contexts {
		if ctx == oldName {
			m.Contexts[i] = newName
			break
		}
	}

	// Update context in all tasks
	for i := range m.Tasks {
		if m.Tasks[i].Context == oldName {
			m.Tasks[i].Context = newName
		}
	}

	m.CurrentContext = newName
}

func (m *Model) DeleteContext() {
	if len(m.Contexts) <= 1 {
		m.ErrorMessage = "Cannot delete the only context"
		return
	}

	// Remove all tasks in this context
	var newTasks []Task
	for _, task := range m.Tasks {
		if task.Context != m.CurrentContext {
			newTasks = append(newTasks, task)
		}
	}
	m.Tasks = newTasks

	// Remove context from list
	var newContexts []string
	for _, ctx := range m.Contexts {
		if ctx != m.CurrentContext {
			newContexts = append(newContexts, ctx)
		}
	}
	m.Contexts = newContexts

	// Switch to first remaining context
	if len(m.Contexts) > 0 {
		m.CurrentContext = m.Contexts[0]
		m.SelectedIndex = 0
	}
}

func (m *Model) ToggleCurrentTaskPriority() {
	tasks := m.GetFilteredTasks()
	if len(tasks) == 0 {
		return
	}

	currentTask := tasks[m.SelectedIndex]
	for i := range m.Tasks {
		if m.Tasks[i].ID == currentTask.ID {
			priorities := []string{"", "low", "medium", "high"}
			currentIdx := 0
			for j, p := range priorities {
				if p == m.Tasks[i].Priority {
					currentIdx = j
					break
				}
			}
			nextIdx := (currentIdx + 1) % len(priorities)
			m.Tasks[i].Priority = priorities[nextIdx]
			break
		}
	}
}

func (m *Model) AddTagToCurrentTask(tag string) {
	tasks := m.GetFilteredTasks()
	if len(tasks) == 0 {
		return
	}

	currentTask := tasks[m.SelectedIndex]
	for i := range m.Tasks {
		if m.Tasks[i].ID == currentTask.ID {
			// Check if tag already exists
			for _, existingTag := range m.Tasks[i].Tags {
				if existingTag == tag {
					return
				}
			}
			m.Tasks[i].Tags = append(m.Tasks[i].Tags, tag)
			break
		}
	}
}

func (m *Model) RemoveTagsFromCurrentTask() {
	tasks := m.GetFilteredTasks()
	if len(tasks) == 0 {
		return
	}

	currentTask := tasks[m.SelectedIndex]
	for i := range m.Tasks {
		if m.Tasks[i].ID == currentTask.ID {
			var newTags []string
			for j, tag := range m.Tasks[i].Tags {
				if !m.RemoveTagChecks[j] {
					newTags = append(newTags, tag)
				}
			}
			m.Tasks[i].Tags = newTags
			break
		}
	}
}

func (m *Model) SetDueDateForCurrentTask(dateStr string) {
	tasks := m.GetFilteredTasks()
	if len(tasks) == 0 {
		return
	}

	currentTask := tasks[m.SelectedIndex]
	for i := range m.Tasks {
		if m.Tasks[i].ID == currentTask.ID {
			if strings.ToLower(dateStr) == "clear" {
				m.Tasks[i].DueDate = ""
			} else if dateStr != "" {
				// Basic date validation (YYYY-MM-DD format)
				parts := strings.Split(dateStr, "-")
				if len(parts) == 3 {
					if year, err := strconv.Atoi(parts[0]); err == nil && year > 1900 && year < 3000 {
						if month, err := strconv.Atoi(parts[1]); err == nil && month >= 1 && month <= 12 {
							if day, err := strconv.Atoi(parts[2]); err == nil && day >= 1 && day <= 31 {
								m.Tasks[i].DueDate = dateStr
								return
							}
						}
					}
				}
				m.ErrorMessage = "Invalid date format. Use YYYY-MM-DD"
			}
			break
		}
	}
}





func (m *Model) UpdateContexts() {
	// Collect contexts from tasks
	taskContexts := make(map[string]bool)
	for _, task := range m.Tasks {
		taskContexts[task.Context] = true
	}

	// Merge with existing contexts (loaded from config)
	mergedContexts := make(map[string]bool)
	for _, ctx := range m.Contexts { // Existing contexts from config
		mergedContexts[ctx] = true
	}
	for ctx := range taskContexts { // Contexts from tasks
		mergedContexts[ctx] = true
	}

	// Rebuild m.Contexts from the merged set
	m.Contexts = make([]string, 0, len(mergedContexts))
	for context := range mergedContexts {
		m.Contexts = append(m.Contexts, context)
	}
	sort.Strings(m.Contexts)

	// Set current context if not set or if current doesn't exist
	if m.CurrentContext == "" || !m.Contains(m.Contexts, m.CurrentContext) { // Use Contains helper
		if len(m.Contexts) > 0 {
			m.CurrentContext = m.Contexts[0]
		} else {
			m.CurrentContext = "Work" // Default context
			m.Contexts = []string{"Work"}
		}
	}
}

// Helper function to check if a string is in a slice
func (m *Model) Contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func (m *Model) SaveStateForUndo() {
	// Deep copy current tasks
	stateCopy := make([]Task, len(m.Tasks))
	copy(stateCopy, m.Tasks)
	
	m.History = append(m.History, stateCopy)
	
	// Limit history size
	if len(m.History) > m.MaxHistory {
		m.History = m.History[1:]
	}
}

func (m *Model) Undo() {
	if len(m.History) == 0 {
		m.ErrorMessage = "Nothing to undo"
		return
	}

	// Restore previous state
	m.Tasks = m.History[len(m.History)-1]
	m.History = m.History[:len(m.History)-1]
	
	// Update contexts and ensure current context is valid
	m.UpdateContexts()
	
	// Reset selection
	m.SelectedIndex = 0
}