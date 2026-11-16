package main

import (
	"charm.land/lipgloss/v2"
)

var (
	titleStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#cba6f7")).Bold(true).MarginBottom(1).PaddingLeft(2)
	itemStyle     = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#cdd6f4"))
	selectedStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#89b4fa")).Bold(true)
	boxStyle      = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#89b4fa")).Padding(1, 4)
)
