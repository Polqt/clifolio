package main

import (
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
	flag.Parse()

	_ = styles.NewThemeFromName(*themeName)

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Oh no! env file not found.")
	}

	p := tea.NewProgram(ui.AppModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
