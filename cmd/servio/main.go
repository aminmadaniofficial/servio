package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/aminmadaniofficial/servio/internal/systemd"
	"github.com/aminmadaniofficial/servio/internal/ui"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	modalStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#FF5F87")).Padding(1, 2).Margin(1, 2)
	logBoxStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("#00BFFF")).Padding(0, 1)
	statusBar = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFF")).Background(lipgloss.Color("#7D56F4")).Padding(0, 1).MarginTop(1)
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).MarginTop(1)
	alertStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Bold(true)
)

const (
	stateList = iota
	stateConfirm
	stateLogs
)

type model struct {
	allServices      []systemd.Service
	filteredServices []systemd.Service
	cursor           int
	offset           int
	maxRows          int
	state            int
	
	actionType  string
	lastMessage string

	searchInput textinput.Model
	isSearching bool
	logsView    viewport.Model
}

func initialModel(svcs []systemd.Service) model {
	// تنظیمات نوار جستجو
	ti := textinput.New()
	ti.Placeholder = "Type to filter services..."
	ti.Prompt = "🔍 Search: "
	ti.CharLimit = 50

	// تنظیمات اولیه
	return model{
		allServices:      svcs,
		filteredServices: svcs,
		state:            stateList,
		searchInput:      ti,
	}
}

func (m model) Init() tea.Cmd { return textinput.Blink }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.maxRows = msg.Height - 8 
		if m.maxRows < 5 { m.maxRows = 5 }
		
		m.logsView.Width = msg.Width - 4
		m.logsView.Height = msg.Height - 6

	case tea.KeyMsg:
		if m.state == stateLogs {
			switch msg.String() {
			case "q", "esc", "l":
				m.state = stateList
				return m, nil
			default:
				m.logsView, cmd = m.logsView.Update(msg)
				return m, cmd
			}
		}

		// === حالت تایید عملیات ===
		if m.state == stateConfirm {
			switch msg.String() {
			case "y", "enter":
				targetSvc := m.filteredServices[m.cursor].Name
				err := executeSystemctl(m.actionType, targetSvc)
				if err != nil {
					m.lastMessage = "❌ Error: " + err.Error()
				} else {
					m.lastMessage = "✅ " + targetSvc + " " + m.actionType + "ed!"
					m.allServices, _ = systemd.GetServices()
					m.filterServices()
				}
				m.state = stateList
			case "n", "esc", "q":
				m.state = stateList
			}
			return m, nil
		}

		if m.isSearching {
			switch msg.String() {
			case "esc", "enter":
				m.isSearching = false
				m.searchInput.Blur()
			default:
				m.searchInput, cmd = m.searchInput.Update(msg)
				m.filterServices()
			}
			return m, cmd
		}

		m.lastMessage = ""
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
				if m.cursor < m.offset { m.offset-- }
			}
		case "down", "j":
			if m.cursor < len(m.filteredServices)-1 {
				m.cursor++
				if m.cursor >= m.offset+m.maxRows { m.offset++ }
			}
		case "/":
			m.isSearching = true
			m.searchInput.Focus()
			return m, textinput.Blink
		case "r":
			m.triggerAction("restart")
		case "s":
			m.triggerAction("stop")
		case "S":
			m.triggerAction("start")
		case "e":
			m.triggerAction("enable")
		case "d":
			m.triggerAction("disable")
		case "l":
			if len(m.filteredServices) > 0 {
				logs := fetchLogs(m.filteredServices[m.cursor].Name)
				m.logsView.SetContent(logs)
				m.logsView.GotoBottom() 
				m.state = stateLogs
			}
		}
	}
	return m, nil
}

func (m *model) filterServices() {
	query := strings.ToLower(m.searchInput.Value())
	if query == "" {
		m.filteredServices = m.allServices
	} else {
		var filtered []systemd.Service
		for _, s := range m.allServices {
			if strings.Contains(strings.ToLower(s.Name), query) {
				filtered = append(filtered, s)
			}
		}
		m.filteredServices = filtered
	}
	m.cursor = 0
	m.offset = 0
}

func (m *model) triggerAction(action string) {
	if len(m.filteredServices) > 0 {
		m.actionType = action
		m.state = stateConfirm
	}
}

func (m model) View() string {
	if m.state == stateConfirm {
		question := fmt.Sprintf("Execute %s on '%s'?\n\n[y] Yes   [n] No", 
			alertStyle.Render(strings.ToUpper(m.actionType)), m.filteredServices[m.cursor].Name)
		return modalStyle.Render(question)
	}

	if m.state == stateLogs {
		header := ui.TitleStyle.Render(" Logs: " + m.filteredServices[m.cursor].Name + " ")
		footer := helpStyle.Render("Press 'q' or 'esc' to back | Use Up/Down to scroll")
		return fmt.Sprintf("\n%s\n%s\n%s", header, logBoxStyle.Render(m.logsView.View()), footer)
	}

	s := ui.TitleStyle.Render(" 🚀 Servio: Systemd Manager ") + "\n"
	
	if m.isSearching || m.searchInput.Value() != "" {
		s += "\n" + m.searchInput.View() + "\n"
	} else if m.lastMessage != "" {
		s += "\n" + m.lastMessage + "\n"
	} else {
		s += "\n"
	}

	end := m.offset + m.maxRows
	if end > len(m.filteredServices) { end = len(m.filteredServices) }

	for i := m.offset; i < end; i++ {
		cursor := "  "
		if m.cursor == i { cursor = ui.CursorStyle.Render("▶ ") }

		status := ui.RunningStyle.Render(m.filteredServices[i].Status)
		if m.filteredServices[i].Status != "running" { status = ui.ExitedStyle.Render(m.filteredServices[i].Status) }

		s += fmt.Sprintf("%s %-40s [%s]\n", cursor, m.filteredServices[i].Name, status)
	}

	// Status Bar & Help
	statusText := fmt.Sprintf(" Total: %d | Filtered: %d ", len(m.allServices), len(m.filteredServices))
	s += "\n" + statusBar.Render(statusText)
	s += helpStyle.Render("\n(/) Search | (r) Restart | (s/S) Stop/Start | (e/d) Enable/Disable | (l) Logs | (q) Quit")
	
	return s
}

func executeSystemctl(action, service string) error {
	cmd := exec.Command("sudo", "systemctl", action, service)
	return cmd.Run()
}

func fetchLogs(service string) string {
	cmd := exec.Command("journalctl", "-u", service, "-n", "100", "--no-pager")
	out, err := cmd.CombinedOutput()
	if err != nil || len(out) == 0 { return "No logs available or access denied (run with sudo)." }
	return string(out)
}

func main() {
	svcs, err := systemd.GetServices()
	if err != nil {
		fmt.Println("Error: Systemd not found or accessible.", err)
		os.Exit(1)
	}

	p := tea.NewProgram(initialModel(svcs), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Fatal error: %v", err)
		os.Exit(1)
	}
}