package ui

import "strings"

var categories = []string{
	"System",
	"Storage",
	"Devices",
	"Network",
}

func (m Model) renderCategories(height int) string {
	var b strings.Builder

	b.WriteString(sectionStyle.Render("CATEGORIES"))
	b.WriteString("\n\n")

	for i, category := range categories {
		if i == m.category {
			b.WriteString(selectedStyle.Render("▶ " + category))
		} else {
			b.WriteString(normalStyle.Render("  " + category))
		}

		b.WriteString("\n")
	}

	return panelStyle(m.panel == panelCategories).
		Width(18).
		Height(height).
		Render(b.String())
}

func (m Model) renderComponents(height int) string {
	var b strings.Builder

	b.WriteString(sectionStyle.Render("COMPONENTS"))
	b.WriteString("\n\n")

	components := m.components()

	if len(components) == 0 {
		b.WriteString(normalStyle.Render("  No devices found"))
	} else {
		for i, component := range components {
			if i == m.selected {
				b.WriteString(selectedStyle.Render("▶ " + component))
			} else {
				b.WriteString(normalStyle.Render("  " + component))
			}

			b.WriteString("\n")
		}
	}

	return panelStyle(m.panel == panelComponents).
		Width(25).
		Height(height).
		Render(b.String())
}
