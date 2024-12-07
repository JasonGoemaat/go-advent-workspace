package main

// An example program demonstrating the pager component from the Bubbles
// component library.

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// You generally won't need this unless you're processing stuff with
// complicated ANSI escape sequences. Turn it on if you notice flickering.
//
// Also keep in mind that high performance rendering only works for programs
// that use the full size of the terminal. We're enabling that below with
// tea.EnterAltScreen().
const useHighPerformanceRenderer = false

// JASON: Ok, loke like these are special 'lipgloss' components for the header and footer
// The 'Right' and 'Left are for the boxes around
var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		// b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		// b.Left = "┤"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()
)

type model struct {
	content    string
	ready      bool
	viewport   viewport.Model
	fullWidth  int
	fullHeight int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			m.fullWidth = msg.Width
			m.fullHeight = msg.Height
			m.viewport = viewport.New(msg.Width/2-6, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.viewport.SetContent(m.content)
			m.ready = true

			// This is only necessary for high performance rendering, which in
			// most cases you won't need.
			//
			// Render the viewport one line below the header.
			m.viewport.YPosition = headerHeight + 1
		} else {
			m.viewport.Width = msg.Width/2 - 6
			m.viewport.Height = msg.Height - verticalMarginHeight
			m.fullWidth = msg.Width
			m.fullHeight = msg.Height
		}

		if useHighPerformanceRenderer {
			// Render (or re-render) the whole viewport. Necessary both to
			// initialize the viewport and when the window is resized.
			//
			// This is needed for high-performance rendering only.
			cmds = append(cmds, viewport.Sync(m.viewport))
		}
	}

	// Handle keyboard and mouse events in the viewport
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// JASON:  Ok, the view renders a single string, joining header, view, and footer
// with `\n` for newlines.   We could add more components apparently
func (m model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	middle := m.viewport.View()
	middle = lipgloss.JoinHorizontal(lipgloss.Center, middle, " <--- ", middle)
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), middle, m.footerView())
}

var titleText = "this is a\n very long title that will require either going to a new\n line or cutting off the right part of\n what I'm trying to say"

// JASON: Ok, here we go.
func (m model) headerView() string {
	title := titleStyle.Render(titleText) // defined above, rounded box with '├' on the right
	// line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title))) // fill the rest of the viewport with a line
	line := strings.Repeat("─", max(0, m.fullWidth-lipgloss.Width(title))) // fill the rest of the viewport with a line
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)           // and center it
}

// JASON: See above, footer uses rounded style 'infoStyle' with left ” and prepends a horizontal line
func (m model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	// line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	line := strings.Repeat("─", max(0, m.fullWidth-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// move this into my main
// func main() {
// 	// Load some text for our viewport
// 	content, err := os.ReadFile("artichoke.md")
// 	if err != nil {
// 		fmt.Println("could not load file:", err)
// 		os.Exit(1)
// 	}

// 	p := tea.NewProgram(
// 		model{content: string(content)},
// 		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
// 		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
// 	)

// 	if _, err := p.Run(); err != nil {
// 		fmt.Println("could not run program:", err)
// 		os.Exit(1)
// 	}
// }
