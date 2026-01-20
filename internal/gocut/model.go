package gocut

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/stanislav-zeman/gocut/internal/navigation"
)

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1).
			MarginBottom(1)

	breadcrumbStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			MarginBottom(1)

	itemStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFDF5")).
				Background(lipgloss.Color("#25A065")).
				PaddingLeft(2)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			MarginTop(1)
)

type Model struct {
	rootNodes       []*navigation.TreeNode
	currentNodes    []*navigation.TreeNode
	cursor          int
	navigationPath  []*navigation.TreeNode
	selectedCommand string
	width           int
	height          int
}

func NewModel(rootNodes []*navigation.TreeNode) Model {
	return Model{
		rootNodes:      rootNodes,
		currentNodes:   rootNodes,
		cursor:         0,
		navigationPath: []*navigation.TreeNode{},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			// Quit without running command
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.currentNodes)-1 {
				m.cursor++
			}

		case "enter":
			if len(m.currentNodes) == 0 {
				return m, nil
			}

			selected := m.currentNodes[m.cursor]
			if selected.IsFolder {
				// Navigate into folder
				m.navigationPath = append(m.navigationPath, selected)
				m.currentNodes = selected.Children
				m.cursor = 0
			} else {
				// Execute command and quit
				m.selectedCommand = selected.Command
				return m, tea.Quit
			}

		case "backspace", "left", "h", "esc":
			// Navigate back up
			if len(m.navigationPath) > 0 {
				// Pop the last folder from navigation path
				m.navigationPath = m.navigationPath[:len(m.navigationPath)-1]

				// Restore the previous level's nodes
				if len(m.navigationPath) > 0 {
					m.currentNodes = m.navigationPath[len(m.navigationPath)-1].Children
				} else {
					m.currentNodes = m.rootNodes
				}

				// Reset cursor to top
				m.cursor = 0
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	if len(m.currentNodes) == 0 {
		return titleStyle.Render("Gocut") + "\n\n" +
			"No commands available.\n\n" +
			helpStyle.Render("Press q to quit")
	}

	var b strings.Builder

	// Title
	b.WriteString(titleStyle.Render("Gocut"))
	b.WriteString("\n")

	// Breadcrumb trail
	breadcrumb := "📁 Home"
	for _, node := range m.navigationPath {
		breadcrumb += " > " + node.Name
	}
	b.WriteString(breadcrumbStyle.Render(breadcrumb))
	b.WriteString("\n")

	// List current nodes
	for i, node := range m.currentNodes {
		var icon string
		if node.IsFolder {
			icon = "📁"
		} else {
			icon = "⚡"
		}

		line := icon + " " + node.Name

		if i == m.cursor {
			b.WriteString(selectedItemStyle.Render(line))
		} else {
			b.WriteString(itemStyle.Render(line))
		}
		b.WriteString("\n")
	}

	// Help text
	b.WriteString("\n")
	helpText := "↑/↓: Navigate | ↵: Select"
	if len(m.navigationPath) > 0 {
		helpText += " | ←/Backspace: Back"
	}
	helpText += " | q: Quit"
	b.WriteString(helpStyle.Render(helpText))

	return b.String()
}

// GetSelectedCommand returns the command that was selected for execution
func (m Model) GetSelectedCommand() string {
	return m.selectedCommand
}
