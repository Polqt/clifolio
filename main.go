package main

import (
	"clifolio/internal/services"
	"clifolio/internal/styles"
	"clifolio/internal/ui"
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
)



func main() {
	themeName := flag.String("theme", "default", "theme name (hacker|dracula|default)")
	sshMode := flag.Bool("ssh-mode", false, "run as SSH server instead of local TUI")
	flag.Parse()

	_ = styles.NewThemeFromName(*themeName)

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Oh no! env file not found.")
	}

	if *sshMode {
		fmt.Println("Starting SSH server mode...")
		services.StartSSHServer(func() tea.Model {
			return ui.AppModel()
		})
	} else {
		p := tea.NewProgram(ui.AppModel(), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	}
}
