package Log

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func InfoLogStyle() *log.Styles {
	styles := log.DefaultStyles()
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("INFO").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("#0E49B5")).
		Foreground(lipgloss.Color("#FFFFFF"))

	return styles
}

func WarnLogStyle() *log.Styles {
	styles := log.DefaultStyles()
	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString("WARN").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("#FDE047")).
		Foreground(lipgloss.Color("#020617"))
	
	return styles
}

func ErrorLogStyle() *log.Styles {
	styles := log.DefaultStyles()
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERROR").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("#EF4444")).
		Foreground(lipgloss.Color("#FFFFFF"))
	
	return styles
}