package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// model contain all app states
type model struct {
	actions []string
	cursor  int
	width   int
	height  int
}

func initModel() model {
	return model{actions: []string{"Logout", "Shutdown", "Lockscreen", "Reboot", "Reboot to UEFI"}, cursor: 0}
}

// Call one time on startup, nothing to do more
func (m model) Init() tea.Cmd {
	return nil
}

// Update react to messages (key press, etc.)
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "j", "down":
			if m.cursor < len(m.actions)-1 {
				m.cursor++
			}
		case "enter":
			return m, tea.Quit
		}
	}
	return m, nil
}

// View return the dipslay screen
func (m model) View() tea.View {
	title := titleStyle.Render("LOGOUTUI")

	var items []string
	for i, action := range m.actions {
		cursor := "  "
		style := itemStyle
		if i == m.cursor {
			cursor = "> "
			style = selectedStyle
		}
		items = append(items, style.Render(cursor+action))
	}

	list := lipgloss.JoinVertical(lipgloss.Left, items...)
	body := lipgloss.JoinVertical(lipgloss.Left, title, list)
	box := boxStyle.Render(body)

	centered := lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)

	v := tea.NewView(centered)
	v.AltScreen = true
	return v
}

func main() {
	p := tea.NewProgram(initModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
