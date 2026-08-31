package ui

import (
	"strings"

	"charm.land/lipgloss/v2"
)

func renderFooter(width int) string {
	left := "↑↓ Navigate   ←→ Panel"
	right := "Tab Switch   q Quit"

	spaces := width - lipgloss.Width(left) - lipgloss.Width(right)

	if spaces < 1 {
		spaces = 1
	}

	return normalStyle.Render(
		left + strings.Repeat(" ", spaces) + right,
	)
}
