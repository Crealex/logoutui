package main

import (
	"charm.land/lipgloss/v2"
)

var (
	titleStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF79C6")).Bold(true).MarginBottom(1)
	itemStyle     = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#F8F8F2"))
	selectedStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#BD93F9")).Bold(true)
	boxStyle      = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#BD93F9")).Padding(1, 4)
)
