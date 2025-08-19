package todo

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
)

// Update implements tea.Model  
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.WindowWidth = msg.Width
		m.WindowHeight = msg.Height
		m.Help.Width = msg.Width
		return m, tea.ClearScreen

	case tea.KeyMsg:
		// Clear error message on any key press
		m.ErrorMessage = ""

		// Handle input mode
		if m.ViewMode == InputView {
			return m.UpdateInputMode(msg)
		} else if m.ViewMode == DateInputView {
			return m.UpdateDateInputMode(msg)
		} else if m.ViewMode == RemoveTagView {
			return m.UpdateRemoveTagMode(msg)
		}

		// Handle different view modes
		switch m.ViewMode {
		case NormalView: // Removed SearchView
			return m.UpdateNormalView(msg)
		case KanbanView:
			return m.UpdateKanbanView(msg)
		case StatsView:
			return m.UpdateStatsView(msg)
		}
	}

	return m, nil
}

// UpdateInputMode handles input dialog updates
func (m Model) UpdateInputMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch {
	case key.Matches(msg, m.KeyMap.Back):
		m.ViewMode = NormalView
		return m, nil

	case key.Matches(msg, m.KeyMap.Enter):
		input := strings.TrimSpace(m.TextInput.Value())
		m.TextInput.SetValue("")
		
		switch m.InputMode {
		case AddTaskInput:
			if input != "" {
				m.SaveStateForUndo()
				m.AddTask(input)
			}
		case EditTaskInput:
			if input != "" {
				m.SaveStateForUndo()
				m.EditCurrentTask(input)
			}
		case AddContextInput:
			if input != "" {
				m.AddContext(input)
			}
		case RenameContextInput:
			if input != "" && input != m.CurrentContext {
				m.RenameContext(input)
			}
		case AddTagInput:
			if input != "" {
				m.SaveStateForUndo()
				m.AddTagToCurrentTask(input)
			}
		// Removed SearchInput case
		// Removed SearchInput case
		case DeleteConfirmInput:
			if strings.ToLower(input) == "y" {
				m.SaveStateForUndo()
				m.DeleteContext()
			}
		}
		
		m.ViewMode = NormalView
		return m, nil
	}

	m.TextInput, cmd = m.TextInput.Update(msg)
	return m, cmd
}

// UpdateDateInputMode handles due date input updates
func (m Model) UpdateDateInputMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch {
	case key.Matches(msg, m.KeyMap.Back):
		m.ViewMode = NormalView
		return m, nil

	case key.Matches(msg, m.KeyMap.Enter):
		day := m.DateInputs[0].Value()
		month := m.DateInputs[1].Value()
		year := m.DateInputs[2].Value()
		dateStr := fmt.Sprintf("%s-%s-%s", year, month, day)
		m.SaveStateForUndo()
		m.SetDueDateForCurrentTask(dateStr)
		m.ViewMode = NormalView
		return m, nil

	case key.Matches(msg, m.KeyMap.Up):
		m.DateInputs[m.DateInputIndex].Blur()
		m.DateInputIndex = (m.DateInputIndex - 1 + 3) % 3
		m.DateInputs[m.DateInputIndex].Focus()

	case key.Matches(msg, m.KeyMap.Down):
		m.DateInputs[m.DateInputIndex].Blur()
		m.DateInputIndex = (m.DateInputIndex + 1) % 3
		m.DateInputs[m.DateInputIndex].Focus()
	}

	m.DateInputs[m.DateInputIndex], cmd = m.DateInputs[m.DateInputIndex].Update(msg)
	return m, cmd
}

// UpdateRemoveTagMode handles remove tag view updates
func (m Model) UpdateRemoveTagMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.KeyMap.Back):
		m.ViewMode = NormalView
		return m, nil

	case key.Matches(msg, m.KeyMap.Enter):
		m.SaveStateForUndo()
		m.RemoveTagsFromCurrentTask()
		m.ViewMode = NormalView
		return m, nil

	case key.Matches(msg, m.KeyMap.Up):
		if m.RemoveTagIndex > 0 {
			m.RemoveTagIndex--
		}

	case key.Matches(msg, m.KeyMap.Down):
		task := m.GetCurrentTask()
		if m.RemoveTagIndex < len(task.Tags)-1 {
			m.RemoveTagIndex++
		}

	case key.Matches(msg, m.KeyMap.Toggle):
		m.RemoveTagChecks[m.RemoveTagIndex] = !m.RemoveTagChecks[m.RemoveTagIndex]
	}

	return m, nil
}

// UpdateNormalView handles normal view updates
func (m Model) UpdateNormalView(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.KeyMap.Quit):
		m.SaveConfig()
		return m, tea.Quit

	case key.Matches(msg, m.KeyMap.Back):
		// Removed search exit logic
		return m, nil

	case key.Matches(msg, m.KeyMap.Up):
		if m.MovingMode {
			m.MoveTaskUp()
		} else {
			m.MoveUp()
		}

	case key.Matches(msg, m.KeyMap.Down):
		if m.MovingMode {
			m.MoveTaskDown()
		} else {
			m.MoveDown()
		}

	case key.Matches(msg, m.KeyMap.Left):
		m.PreviousContext()

	case key.Matches(msg, m.KeyMap.Right):
		m.NextContext()

	case key.Matches(msg, m.KeyMap.Toggle):
		if len(m.GetFilteredTasks()) > 0 {
			m.SaveStateForUndo()
			m.ToggleCurrentTask()
		}

	case key.Matches(msg, m.KeyMap.Add):
		m.ShowInputDialog(AddTaskInput, "Add new task:")

	case key.Matches(msg, m.KeyMap.Edit):
		if len(m.GetFilteredTasks()) > 0 {
			task := m.GetCurrentTask()
			m.ShowInputDialog(EditTaskInput, "Edit task:")
			m.TextInput.SetValue(task.Task)
		}

	case key.Matches(msg, m.KeyMap.Delete):
		if len(m.GetFilteredTasks()) > 0 {
			m.SaveStateForUndo()
			m.DeleteCurrentTask()
		}

	case key.Matches(msg, m.KeyMap.AddContext):
		m.ShowInputDialog(AddContextInput, "New context name:")

	case key.Matches(msg, m.KeyMap.RenameContext):
		m.ShowInputDialog(RenameContextInput, "Rename context to:")
		m.TextInput.SetValue(m.CurrentContext)

	case key.Matches(msg, m.KeyMap.DeleteContext):
		if len(m.Contexts) > 1 {
			m.ShowInputDialog(DeleteConfirmInput, fmt.Sprintf("Delete context '%s'? (y/n):", m.CurrentContext))
		} else {
			m.ErrorMessage = "Cannot delete the only context"
		}

	case key.Matches(msg, m.KeyMap.TogglePriority):
		if len(m.GetFilteredTasks()) > 0 {
			m.SaveStateForUndo()
			m.ToggleCurrentTaskPriority()
		}

	case key.Matches(msg, m.KeyMap.AddTag):
		if len(m.GetFilteredTasks()) > 0 {
			m.ShowInputDialog(AddTagInput, "Add tag:")
		}

	case key.Matches(msg, m.KeyMap.RemoveTag):
		if len(m.GetFilteredTasks()) > 0 {
			m.ShowRemoveTagDialog()
		}

	case key.Matches(msg, m.KeyMap.SetDueDate):
		if len(m.GetFilteredTasks()) > 0 {
			m.ShowDateInputDialog()
		}

	case key.Matches(msg, m.KeyMap.ClearDueDate):
		if len(m.GetFilteredTasks()) > 0 {
			m.SaveStateForUndo()
			m.SetDueDateForCurrentTask("clear")
		}

	// Removed Search key binding
	case key.Matches(msg, m.KeyMap.KanbanView):
		m.ViewMode = KanbanView

	case key.Matches(msg, m.KeyMap.StatsView):
		m.ViewMode = StatsView

	case key.Matches(msg, m.KeyMap.Undo):
		m.Undo()

	case key.Matches(msg, m.KeyMap.Move):
		if len(m.GetFilteredTasks()) > 0 {
			m.MovingMode = !m.MovingMode
			if m.MovingMode {
				m.MovingTaskIndex = m.SelectedIndex
			} else {
				m.SaveStateForUndo()
			}
		}
	}

	return m, nil
}

// UpdateKanbanView handles kanban view updates
func (m Model) UpdateKanbanView(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.KeyMap.Back), key.Matches(msg, m.KeyMap.Quit), key.Matches(msg, m.KeyMap.KanbanView):
		m.ViewMode = NormalView
	}
	return m, nil
}

// UpdateStatsView handles stats view updates  
func (m Model) UpdateStatsView(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.KeyMap.Back), key.Matches(msg, m.KeyMap.Quit), key.Matches(msg, m.KeyMap.StatsView):
		m.ViewMode = NormalView
	}
	return m, nil
}