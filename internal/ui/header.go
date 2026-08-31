package ui

import (
	"strings"

	"charm.land/lipgloss/v2"
)

func renderHeader(width int) string {
	logo := lipgloss.NewStyle().
		Bold(true).
		Foreground(purple).
		Render("◈ OpenCase")

	subtitle := normalStyle.Render(
		"  Explore your Linux computer",
	)

	right := normalStyle.Render("localhost")

	left := logo + subtitle

	spaces := width - lipgloss.Width(left) - lipgloss.Width(right)

	if spaces < 1 {
		spaces = 1
	}

	return left + strings.Repeat(" ", spaces) + right
}
