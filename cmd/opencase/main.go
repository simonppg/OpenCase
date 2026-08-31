package main

import (
	"fmt"
	"os"

	"charm.land/bubbletea/v2"

	"github.com/simonppg/OpenCase/internal/hardware"
	"github.com/simonppg/OpenCase/internal/ui"
)

func main() {
	system, err := hardware.CollectSystem()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to collect system information:", err)
		os.Exit(1)
	}

	model := ui.New(system)

	program := tea.NewProgram(model)

	if _, err := program.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "failed to run OpenCase:", err)
		os.Exit(1)
	}
}
