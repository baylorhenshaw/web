package web

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

// Logger is the global logger for the server
var Logger = log.New(os.Stderr)

// Configure logging format in the console
func ConfigureLogs() {
	styles := log.DefaultStyles()

	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("INFO").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("63")).
		Foreground(lipgloss.Color("16")).
		Bold(true)
	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString("WARN").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("166")).
		Foreground(lipgloss.Color("16")).
		Bold(true)
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERRO").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("124")).
		Foreground(lipgloss.Color("16")).
		Bold(true)
	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString("FATAL").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("124")).
		Foreground(lipgloss.Color("16")).
		Bold(true)
	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString("DEBU").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("134")).
		Foreground(lipgloss.Color("16")).
		Bold(true)

	Logger.SetStyles(styles)
	Logger.SetReportTimestamp(false)
}

// Enables debug mode for the logger
func EnableDebug() {
	Logger.SetLevel(log.DebugLevel)
	Logger.Debug("Debug mode has been enabled.")
}
