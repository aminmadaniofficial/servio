package ui

import "github.com/charmbracelet/lipgloss"

var (
	TitleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(0, 1)

	CursorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF5F87"))

	RunningStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FF00"))

	ExitedStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF5F87"))
)