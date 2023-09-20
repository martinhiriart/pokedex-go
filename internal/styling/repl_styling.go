package styling

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func PrintStyledMessage(msg string, msgType string) {

	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00ff00")).
		Bold(true)
	errStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ff0000")).
		Bold(true)
	menuHeaderStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00ffff")).
		Bold(true)
	menuStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ff00ff"))

	var renderedMsg string
	var newLineString string
	switch msgType {
	case "Header":
		renderedMsg = headerStyle.Render(msg)
		newLineString = ""
	case "Menu":
		renderedMsg = menuStyle.Render(msg)
		newLineString = "\n"
	case "MenuHeader":
		renderedMsg = menuHeaderStyle.Render(msg)
		newLineString = "\n"
	case "StatDisplay":
		renderedMsg = menuStyle.Render(msg)
		newLineString = ""
	case "Error":
		renderedMsg = errStyle.Render(msg)
		newLineString = "\n"
	}
	fmt.Printf("%s%v", renderedMsg, newLineString)
}
