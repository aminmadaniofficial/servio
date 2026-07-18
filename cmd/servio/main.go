package main

import (
	"fmt"
	"os"

	"github.com/aminmadaniofficial/servio/internal/systemd"
	"github.com/aminmadaniofficial/servio/internal/ui"
	"github.com/charmbracelet/bubbletea"
)

type model struct {
	services []systemd.Service
	cursor   int
	offset   int // برای مدیریت اسکرول
	maxRows  int // تعداد ردیف‌های قابل نمایش در ترمینال
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
				if m.cursor < m.offset { m.offset-- }
			}
		case "down", "j":
			if m.cursor < len(m.services)-1 {
				m.cursor++
				if m.cursor >= m.offset+m.maxRows { m.offset++ }
			}
		case "window size": // این مورد برای تشخیص ابعاد ترمینال عالیه
			// اینجا می‌تونی قد ترمینال رو بگیری
		}
	}
	return m, nil
}

func (m model) View() string {
	m.maxRows = 15 // تعداد خطوطی که می‌خوای همیشه نشون بده
	s := ui.TitleStyle.Render(" Servio: Systemd Manager ") + "\n\n"
	
	end := m.offset + m.maxRows
	if end > len(m.services) { end = len(m.services) }

	for i := m.offset; i < end; i++ {
		cursor := "  "
		if m.cursor == i { cursor = ui.CursorStyle.Render("> ") }
		
		status := ui.RunningStyle.Render(m.services[i].Status)
		if m.services[i].Status != "running" { status = ui.ExitedStyle.Render(m.services[i].Status) }

		s += fmt.Sprintf("%s %-30s [%s]\n", cursor, m.services[i].Name, status)
	}
	s += fmt.Sprintf("\n(j/k move, q quit | Showing %d-%d of %d)\n", m.offset+1, end, len(m.services))
	return s
}

func main() {
	svcs, _ := systemd.GetServices()
	p := tea.NewProgram(model{services: svcs, maxRows: 15})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}