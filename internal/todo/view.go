package todo

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Styles
var (
	// Base styles
	baseStyle = lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1)

	// Title styles
	titleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFDF5")).
		Background(lipgloss.Color("#25A065")).
		Padding(0, 1).
		Bold(true)

	// Task styles
	taskStyle = lipgloss.NewStyle().
		PaddingLeft(2)

	selectedTaskStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#EE6FF8")).
		Background(lipgloss.Color("#313244")).
		PaddingLeft(2)

	completedTaskStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#A6E3A1")).
		Strikethrough(true)

	// Priority styles
	highPriorityStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F38BA8"))

	mediumPriorityStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAB387"))

	lowPriorityStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F9E2AF"))

	// Context styles
	contextStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#89B4FA")).
		Bold(true)

	// Error style
	errorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F38BA8")).
		Bold(true)

	// Help style
	helpStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6C7086"))

	// Input styles
	inputStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1).
		Margin(1)
)

// View implements tea.Model
func (m Model) View() string {
	switch m.ViewMode {
	case InputView:
		return m.RenderInputView()
	case DateInputView:
		return m.RenderDateInputView()
	case RemoveTagView:
		return m.RenderRemoveTagView()
	case KanbanView:
		return m.RenderKanbanView()
	case StatsView:
		return m.RenderStatsView()
	default:
		return m.RenderNormalView()
	}
}

// RenderNormalView renders the main task list view
func (m Model) RenderNormalView() string {
	var content strings.Builder

	// Header
	contextText := fmt.Sprintf("Context: %s", m.CurrentContext)
	content.WriteString(titleStyle.Render(contextText) + "\n\n")
	// Tasks
	tasks := m.GetFilteredTasks()
	if len(tasks) == 0 {
		if len(m.Contexts) == 0 {
			content.WriteString("No contexts exist. Press 'n' to create one.\n")
		} else {
			content.WriteString("No tasks in this context. Press 'a' to add one.\n")
		}
	} else {
		for i, task := range tasks {
			taskLine := m.RenderTask(task, i == m.SelectedIndex, i == m.MovingTaskIndex && m.MovingMode)
			content.WriteString(taskLine + "\n")
		}
	}

	// Error message
	if m.ErrorMessage != "" {
		content.WriteString("\n" + errorStyle.Render(m.ErrorMessage) + "\n")
	}

	// Help
	m.Help.ShowAll = true
	content.WriteString("\n" + helpStyle.Render(m.Help.View(m.KeyMap)))

	return baseStyle.Render(content.String())
}

// RenderTask renders a single task
func (m Model) RenderTask(task Task, selected, moving bool) string {
	// Checkbox
	checkbox := "[ ]"
	if task.Checked {
		checkbox = "[✓]"
	}

	// Priority indicator
	priority := ""
	switch task.Priority {
	case "high":
		priority = highPriorityStyle.Render("!!! ")
	case "medium":
		priority = mediumPriorityStyle.Render("!! ")
	case "low":
		priority = lowPriorityStyle.Render("! ")
	}

	// Task text
	taskText := task.Task

	// Tags
	tags := ""
	if len(task.Tags) > 0 {
		tags = " > " + strings.Join(task.Tags, ", ")
	}

	// Due date
	dueDate := ""
	if task.DueDate != "" {
		dueDate = fmt.Sprintf(" [Due: %s]", task.DueDate)
	}

	// Combine text
	text := fmt.Sprintf("%s %s%s%s", checkbox, taskText, tags, dueDate)

	// Apply styles
	style := taskStyle
	if task.Checked {
		style = completedTaskStyle
	}

	if selected {
		style = style.Copy().Background(lipgloss.Color("#313244"))
	}

	if moving {
		style = style.Copy().Bold(true)
	}

	return priority + style.Render(text)
}

// RenderInputView renders input dialogs
func (m Model) RenderInputView() string {
	return inputStyle.Render(
		fmt.Sprintf("%s\n\n%s", m.InputPrompt, m.TextInput.View()),
	)
}

// RenderDateInputView renders due date input dialog
func (m Model) RenderDateInputView() string {
	var content strings.Builder
	content.WriteString("Set due date (YYYY-MM-DD):\n\n")
	inputs := []string{
		fmt.Sprintf("Day: %s", m.DateInputs[0].View()),
		fmt.Sprintf("Month: %s", m.DateInputs[1].View()),
		fmt.Sprintf("Year: %s", m.DateInputs[2].View()),
	}
	for i, input := range inputs {
		if i == m.DateInputIndex {
			content.WriteString(selectedTaskStyle.Render(input) + "\n")
		} else {
			content.WriteString(input + "\n")
		}
	}
	return inputStyle.Render(content.String())
}

// RenderRemoveTagView renders remove tag view
func (m Model) RenderRemoveTagView() string {
	var content strings.Builder
	content.WriteString("Select tags to remove:\n\n")
	task := m.GetCurrentTask()
	for i, tag := range task.Tags {
		checkbox := "[ ]"
		if m.RemoveTagChecks[i] {
			checkbox = "[✓]"
		}
		line := fmt.Sprintf("%s %s", checkbox, tag)
		if i == m.RemoveTagIndex {
			content.WriteString(selectedTaskStyle.Render(line) + "\n")
		} else {
			content.WriteString(line + "\n")
		}
	}
	return inputStyle.Render(content.String())
}

// RenderKanbanView renders the kanban board
func (m Model) RenderKanbanView() string {
	var content strings.Builder
	
	content.WriteString(titleStyle.Render("Kanban View (ESC to return)") + "\n\n")

	if len(m.Contexts) == 0 {
		content.WriteString("No contexts available.\n")
		return baseStyle.Render(content.String())
	}

	// Calculate column width
	colWidth := (m.WindowWidth - 4) / len(m.Contexts)
	if colWidth < 20 {
		colWidth = 20
	}

	// Render columns
	var columns []string
	for _, context := range m.Contexts {
		var column strings.Builder
		
		// Column header
		header := contextStyle.Render(context)
		column.WriteString(header + "\n")
		column.WriteString(strings.Repeat("─", colWidth) + "\n")

		// Tasks in this context
		tasks := m.GetTasksForContext(context)
		for _, task := range tasks {
			taskText := task.Task
			if len(taskText) > colWidth-4 {
				taskText = taskText[:colWidth-7] + "..."
			}

			tags := ""
			if len(task.Tags) > 0 {
				tags = " > " + strings.Join(task.Tags, ", ")
			}

			dueDate := ""
			if task.DueDate != "" {
				dueDate = fmt.Sprintf(" [Due: %s]", task.DueDate)
			}

			if task.Checked {
				column.WriteString(completedTaskStyle.Render(fmt.Sprintf("✓ %s%s%s", taskText, tags, dueDate)) + "\n")
			} else {
				column.WriteString(taskStyle.Render(fmt.Sprintf("• %s%s%s", taskText, tags, dueDate)) + "\n")
			}
		}

		columns = append(columns, column.String())
	}

	// Combine columns side by side (simplified - in real implementation you'd use lipgloss.JoinHorizontal)
	for i, col := range columns {
		if i > 0 {
			content.WriteString(" | ")
		}
		content.WriteString(col)
	}

	return baseStyle.Render(content.String())
}

// RenderStatsView renders the statistics view
func (m Model) RenderStatsView() string {
	var content strings.Builder
	
	content.WriteString(titleStyle.Render("Statistics (ESC to return)") + "\n\n")

	// Overall stats
	total := len(m.Tasks)
	completed := 0
	for _, task := range m.Tasks {
		if task.Checked {
			completed++
		}
	}

	completionRate := 0.0
	if total > 0 {
		completionRate = float64(completed) / float64(total) * 100
	}

	content.WriteString(fmt.Sprintf("Total Tasks: %d\n", total))
	content.WriteString(fmt.Sprintf("Completed: %d (%.1f%%)\n\n", completed, completionRate))

	// Context stats
	content.WriteString("Context Statistics:\n")
	for _, context := range m.Contexts {
		tasks := m.GetTasksForContext(context)
		ctxTotal := len(tasks)
		ctxCompleted := 0
		for _, task := range tasks {
			if task.Checked {
				ctxCompleted++
			}
		}

		ctxRate := 0.0
		if ctxTotal > 0 {
			ctxRate = float64(ctxCompleted) / float64(ctxTotal) * 100
		}

		content.WriteString(fmt.Sprintf("  %s: %d/%d (%.1f%%)\n", 
			contextStyle.Render(context), ctxCompleted, ctxTotal, ctxRate))
	}

	return baseStyle.Render(content.String())
}