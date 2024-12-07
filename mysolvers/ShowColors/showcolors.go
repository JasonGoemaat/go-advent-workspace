package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	showcolors()
}

func showcolors() {
	for i := range 16 {
		colorString := lipgloss.Color(fmt.Sprintf("%d", i))
		text := fmt.Sprintf("Color %s", colorString)
		fmt.Printf("%s: %s or %s\n", text, lipgloss.NewStyle().Foreground(colorString).Render(text), lipgloss.NewStyle().Background(colorString).Render(text))
	}
	fmt.Printf("\nNow bold:\n\n")
	for i := range 16 {
		colorString := lipgloss.Color(fmt.Sprintf("%d", i))
		text := fmt.Sprintf("Color %s", colorString)
		fmt.Printf("%s: %s or %s\n", text, lipgloss.NewStyle().Foreground(colorString).Bold(true).Render(text), lipgloss.NewStyle().Background(colorString).Bold(true).Render(text))
	}
}

func showcolors2() {
	for i := range 256 {
		colorString := lipgloss.Color(fmt.Sprintf("%d", i))
		text := fmt.Sprintf("%03d", i)
		fmt.Printf("%s:%s %s  ", text, lipgloss.NewStyle().Foreground(colorString).Render(text), lipgloss.NewStyle().Background(colorString).Render(text))
		if ((i + 1) % 8) == 0 {
			fmt.Printf("\n")
		}
	}
	// fmt.Printf("\nNow bold:\n\n")
	// for i := range 16 {
	// 	colorString := lipgloss.Color(fmt.Sprintf("%d", i))
	// 	text := fmt.Sprintf("Color %s", colorString)
	// 	fmt.Printf("%s: %s or %s\n", text, lipgloss.NewStyle().Foreground(colorString).Bold(true).Render(text), lipgloss.NewStyle().Background(colorString).Bold(true).Render(text))
	// }
}

func showascii() {
	for r := range 16 {
		if r < 2 {
			continue
		}
		for c := range 16 {
			ascii := (r << 4) | c
			fmt.Printf(" %02x '%c'  ", ascii, ascii)
		}
		fmt.Printf("\n")
	}

	// 2589 ▉ Left Seven Eighths Block

	// chars for robot: 25BC down 25c0 left, 25B2 up, 25B6 right

	// other:
	// 25A0 'Black Square' ■■
	//                     ■■■■
	// 25AB small white square: ▫
	// 228F, 2290 - square image of and square original of: ⊏⊐ ⊏⊐

}
