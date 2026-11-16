package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// model contain all app states
type model struct {
	cursor   int
	selected int
	width    int
	height   int
}

type command struct {
	label string   // display name
	argv  []string // command and arguements
	icon  string
}

var commands = []command{
	{"Logout", []string{"hyprctl", "dispatch", "exit"}, "󰍃"},
	{"Shutdown", []string{"systemctl", "poweroff"}, ""},
	{"Lockscreen", []string{"loginctl", "lock-session"}, ""},
	{"Suspend", []string{"systemctl", "suspend"}, "󰤄"},
	{"Reboot", []string{"systemctl", "reboot"}, "󰑐"},
	{"Reboot to UEFI", []string{"systemctl", "reboot", "--firmware-setup"}, ""},
}

func initModel() model {
	return model{cursor: 0, selected: -1}
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
			m.selected = -1
			return m, tea.Quit
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "j", "down":
			if m.cursor < len(commands)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = m.cursor
			return m, tea.Quit
		}
	}
	return m, nil
}

// View return the dipslay screen
func (m model) View() tea.View {
	title := titleStyle.Render("LOGOUTUI")

	var items []string
	for i, cmd := range commands {
		cursor := "  "
		style := itemStyle
		if i == m.cursor {
			cursor = "> "
			style = selectedStyle
		}
		items = append(items, style.Render(cursor+cmd.icon+" "+cmd.label))
	}

	list := lipgloss.JoinVertical(lipgloss.Left, items...)
	body := lipgloss.JoinVertical(lipgloss.Center, title, list)
	box := boxStyle.Render(body)

	centered := lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)

	v := tea.NewView(centered)
	v.AltScreen = true
	return v
}

func main() {
	p := tea.NewProgram(initModel())
	result, err := p.Run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	m := result.(model)
	if m.selected < 0 {
		return
	}
	a := commands[m.selected]
	cmd := exec.Command(a.argv[0], a.argv[1:]...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true} // launch a side process, independant to kitty
	if err := cmd.Start(); err != nil {
		fmt.Println("Error with the command: ", err)
		os.Exit(1)
	}
}
