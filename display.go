package display

import (
	"fmt"
	"golang.org/x/term"
)

const (
	ansiClearScreen = "\033[2J"
	ansiResetColor  = "\033[0m"
	ansiCyan        = "\033[36m"
)

func LoginPrompt() string {
	// Get the terminal size
	size, _ := term.GetSize(0)
	rows, cols := size.Height, size.Width

	// Calculate the center position
	x := (cols - len("Please log in:")) / 2
	y := rows / 2

	// Create the login prompt
	loginPrompt := fmt.Sprintf("%s\033[%d;%dH%s%sUsername: %s%s",
		ansiClearScreen, y, x, ansiCyan, ansiResetColor, ansiCyan, ansiResetColor)

	return loginPrompt
}
