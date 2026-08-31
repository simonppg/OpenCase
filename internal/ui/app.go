package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/simonppg/OpenCase/internal/hardware"
)

const (
	panelCategories = iota
	panelComponents
	panelDetails
)

type Model struct {
	width  int
	height int

	category int
	panel    int
	selected int

	system hardware.SystemInfo
}

func New(system hardware.SystemInfo) Model {
	return Model{
		category: 0,
		panel:    panelComponents,
		selected: 0,
		system:   system,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyPressMsg:
		switch msg.String() {

		case "q", "ctrl+c":
			return m, tea.Quit

		case "tab", "right", "l":
			m.panel = nextPanel(m.panel)

		case "shift+tab", "left", "h":
			m.panel = previousPanel(m.panel)

		case "up", "k":
			m.moveUp()

		case "down", "j":
			m.moveDown()

		case "enter":
			if m.panel == panelCategories {
				m.panel = panelComponents
			}

		case "esc":
			if m.panel == panelDetails {
				m.panel = panelComponents
			}
		}
	}

	return m, nil
}

func (m *Model) moveUp() {
	switch m.panel {
	case panelCategories:
		m.category = moveSelection(
			m.category,
			len(categories),
			-1,
		)
		m.selected = 0

	case panelComponents:
		m.selected = moveSelection(
			m.selected,
			len(m.components()),
			-1,
		)
	}
}

func (m *Model) moveDown() {
	switch m.panel {
	case panelCategories:
		m.category = moveSelection(
			m.category,
			len(categories),
			1,
		)
		m.selected = 0

	case panelComponents:
		m.selected = moveSelection(
			m.selected,
			len(m.components()),
			1,
		)
	}
}

func (m Model) View() tea.View {
	if m.width == 0 || m.height == 0 {
		view := tea.NewView("")
		view.AltScreen = true
		return view
	}

	bodyHeight := m.height - 4

	if bodyHeight < 5 {
		bodyHeight = 5
	}

	header := renderHeader(m.width)

	left := m.renderCategories(bodyHeight)
	middle := m.renderComponents(bodyHeight)
	right := m.renderDetails(bodyHeight)

	body := lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		" ",
		middle,
		" ",
		right,
	)

	footer := renderFooter(m.width)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
		footer,
	)

	view := tea.NewView(content)
	view.AltScreen = true
	view.BackgroundColor = bg
	view.WindowTitle = "OpenCase"

	return view
}

func (m Model) components() []string {
	switch categories[m.category] {

	case "System":
		return []string{
			"Overview",
			"CPU",
			"Memory",
		}

	case "Storage":
		result := make([]string, 0, len(m.system.Storage))

		for _, disk := range m.system.Storage {
			result = append(result, disk.Name)
		}

		return result

	case "Devices":
		result := make([]string, 0, len(m.system.PCI))

		for _, device := range m.system.PCI {
			result = append(result, device.Address)
		}

		return result

	case "Network":
		result := make([]string, 0, len(m.system.Network))

		for _, network := range m.system.Network {
			result = append(result, network.Name)
		}

		return result
	}

	return nil
}

func (m Model) selectedComponent() string {
	components := m.components()

	if len(components) == 0 {
		return ""
	}

	if m.selected >= len(components) {
		return components[len(components)-1]
	}

	return components[m.selected]
}

func formatBytes(bytes uint64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
		TB = GB * 1024
	)

	switch {
	case bytes >= TB:
		return fmt.Sprintf("%.2f TB", float64(bytes)/float64(TB))
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}

func formatUptime(seconds uint64) string {
	days := seconds / 86400
	seconds %= 86400

	hours := seconds / 3600
	seconds %= 3600

	minutes := seconds / 60

	if days > 0 {
		return fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
	}

	return fmt.Sprintf("%dh %dm", hours, minutes)
}
