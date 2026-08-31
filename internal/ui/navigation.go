package ui

func nextPanel(panel int) int {
	return (panel + 1) % 3
}

func previousPanel(panel int) int {
	if panel == 0 {
		return 2
	}

	return panel - 1
}

func moveSelection(current, total, delta int) int {
	if total <= 0 {
		return 0
	}

	next := current + delta

	if next < 0 {
		next = 0
	}

	if next >= total {
		next = total - 1
	}

	return next
}
