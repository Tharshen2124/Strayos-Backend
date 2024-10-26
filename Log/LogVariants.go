package Log

import (
	"net/http"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

// log your routes with the url path and method showing
func Route(request *http.Request) {
	styles := InfoLogStyle()

	styles.Keys["route"] = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	styles.Values["route"] = lipgloss.NewStyle().Bold(true)

	styles.Keys["method"] = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	styles.Values["method"] = lipgloss.NewStyle().Foreground(lipgloss.Color("#4ADE80")).Bold(true)
	
	logger := log.New(os.Stderr)
	logger.SetStyles(styles)
	logger.Info("Router has been hit at", "route", request.URL.Path, "method", request.Method)		
}

// logs the error message in neat format
func Error(errorMessage string) {
	styles := ErrorLogStyle()

	styles.Keys["Details"] = lipgloss.NewStyle().Foreground(lipgloss.Color("#dc2626"))
	styles.Values["Details"] = lipgloss.NewStyle().Foreground(lipgloss.Color("#f87171"))

	logger := log.New(os.Stderr)
	logger.SetStyles(styles)
	logger.Error("Error detected! \n", "Details", errorMessage)		
}
