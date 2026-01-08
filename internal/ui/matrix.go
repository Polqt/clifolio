package ui

import (
	"clifolio/internal/styles"
	"clifolio/internal/ui/state"
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type matrixModel struct {
	columns []column
	width   int
	height  int
	frame   int
}

type column struct {
	chars  []rune
	yPos   int
	speed  int
	length int
}

type matrixTickMsg struct{}

func MatrixModel() *matrixModel {
	return &matrixModel{
		width:  80,
		height: 24,
	}
}

func (m *matrixModel) Init() tea.Cmd {
	return tickMatrix()
}

func tickMatrix() tea.Cmd {
	return tea.Tick(50*time.Millisecond, func(t time.Time) tea.Msg {
		return matrixTickMsg{}
	})
}

func (m *matrixModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.initColumns()

	case matrixTickMsg:
		m.frame++
		m.updateColumns()
		return m, tickMatrix()

	case tea.KeyMsg:
		return m, func() tea.Msg { return state.ScreenMenu }
	}

	return m, nil
}

func (m *matrixModel) updateColumns() {
	for i := range m.columns {
		col := &m.columns[i]

		// Only update at the column's speed intereval
		if m.frame%col.speed == 0 {
			col.yPos++

			// Reset column when it goes off screen
			if col.yPos > m.height+col.length {
				col.yPos = -col.length
				col.chars = generateMatrixChars(m.height + 20)
				col.speed = rand.Intn(3) + 1
				col.length = rand.Intn(15) + 5
			}
		}
	}
}

func (m *matrixModel) initColumns() {
	m.columns = make([]column, m.width)

	for i := range m.columns {
		m.columns[i] = column{
			chars:  generateMatrixChars(m.height + 20),
			yPos:   rand.Intn(m.height) - m.height,
			speed:  rand.Intn(3) + 1,
			length: rand.Intn(15) + 5,
		}
	}
}

func generateMatrixChars(count int) []rune {
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%^&*()-_=+[]{}|;:',.<>?/`~")

	result := make([]rune, count)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return result
}

func (m *matrixModel) View() string {
	if len(m.columns) == 0 {
		m.initColumns()
	}

	theme := styles.NewThemeFromName("default")

	// 2D grid to hold characters
	grid := make([][]rune, m.height)
	for i := range grid {
		grid[i] = make([]rune, m.width)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	// Draw each column into the grid
	for x, col := range m.columns {
		if x >= m.width {
			continue
		}

		// Draw the tail of chars
		for i := 0; i < col.length; i++ {
			y := col.yPos - i
			if y >= 0 && y < m.height {
				charIndex := (y + i) % len(col.chars)
				grid[y][x] = col.chars[charIndex]
			}
		}
	}

	// Convert grid to string with color gradient
	var output string
	for y, row := range grid {
		for x, char := range row {
			if char == ' ' {
				output += " "
			} else {
				// Brightness based on column position
				col := m.columns[x]
				distFromHead := col.yPos - y

				var style lipgloss.Style
				if distFromHead == 0 {
					style = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))
				} else if distFromHead < col.length/3 {
					style = lipgloss.NewStyle().Foreground(theme.Primary)
				} else if distFromHead < col.length*2/3 {
					style = lipgloss.NewStyle().Foreground(theme.Secondary)
				} else {
					style = lipgloss.NewStyle().Foreground(lipgloss.Color("#003300"))
				}

				output += style.Render(string(char))
			}
		}
		if y < m.height-1 {
			output += "\n"
		}
	}

	return output
}
