package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

// model contain all app states
type model struct {
	actions []string
	cursor  int
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
	s := "What do you want?\n"

	for i, action := range m.actions {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, action)
	}
	v := tea.NewView(s)
	v.AltScreen = true // Display on fullscreen, like vim
	return v
}

func main() {
	p := tea.NewProgram(initModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
