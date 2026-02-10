package main

import (
	"fmt"
	"os"

	"github.com/cladamos/clawea/config"
	"github.com/cladamos/clawea/model"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	config.EnsureConfig()

	p := tea.NewProgram(model.InitialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
