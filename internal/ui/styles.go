package ui

import "charm.land/lipgloss/v2"

var (
	white  = lipgloss.Color("#FFFFFF")
	gray   = lipgloss.Color("#888888")
	dark   = lipgloss.Color("#444444")
	purple = lipgloss.Color("#7C3AED")
	cyan   = lipgloss.Color("#22D3EE")
	bg     = lipgloss.Color("#111111")

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(white)

	normalStyle = lipgloss.NewStyle().
			Foreground(gray)

	selectedStyle = lipgloss.NewStyle().
			Foreground(white).
			Background(purple).
			Bold(true)

	sectionStyle = lipgloss.NewStyle().
			Foreground(cyan).
			Bold(true)
)

func panelStyle(active bool) lipgloss.Style {
	borderColor := dark

	if active {
		borderColor = purple
	}

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor).
		Padding(1, 2)
}
