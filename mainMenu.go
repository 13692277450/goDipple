package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	SignalString                string
	Init_Settings               bool = false
	Init_Logger                 bool = false
	Init_MySQL                  bool = false
	Init_Redis                  bool = false
	Init_MongoDB                bool = false
	Init_Kafka                  bool = false
	Init_RabbitMQ               bool = false
	Init_NATS                   bool = false
	Init_Consul                 bool = false
	Init_ETC                    bool = false
	Init_Zap                    bool = false
	Init_PrettyShutDown         bool = false
	Init_ProjectFolderStructure bool = false
)

var (
	cyan   = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFFF"))
	green  = lipgloss.NewStyle().Foreground(lipgloss.Color("#32CD32"))
	gray   = lipgloss.NewStyle().Foreground(lipgloss.Color("#696969"))
	gold   = lipgloss.NewStyle().Foreground(lipgloss.Color("#B8860B"))
	red    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	yellow = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFF00"))
	purple = lipgloss.NewStyle().Foreground(lipgloss.Color("#800080"))
	blue   = lipgloss.NewStyle().Foreground(lipgloss.Color("#0000FF"))
)

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

func initialModel() model {
	return model{
		// Our to-do list is a grocery list
		choices: []string{"Init Settings", "Init Project Folder Structure", "Init MySQL",
			"Init Redis", "Init MongoDB", "Init Kafka", "Init RabbitMQ", "Init Zap",
			"Init Consul", "Init Pretty Shutdown"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "enter":
			{
				fmt.Println("You choosed items: ")
				for i := range m.selected {
					fmt.Println("\n" + cyan.Render(m.choices[i]))
					switch m.choices[i] {
					case "Init Settings":
						Init_Settings = true
						SettingsCfg()
					case "Init Project Folder Structure":
						Init_ProjectFolderStructure = true
						ProjectFolderStructureCfg()
					case "Init MySQL":
						Init_MySQL = true
						MySQLCfg()
					case "Init Redis":
						Init_Redis = true
						RedisCfg()
					case "Init MongoDB":
						Init_MongoDB = true
						MongoDBCfg()
					case "Init Kafka":
						Init_Kafka = true
						KafkaCfg()
					case "Init RabbitMQ":
						Init_RabbitMQ = true
						RabbitMQCfg()
					case "Init NATS":
						Init_NATS = true
					case "Init Consul":
						Init_Consul = true
						ConsulCfg()
					case "Init ETC":
						Init_ETC = true
					case "Init Zap":
						Init_Zap = true
						ZapCfg()
					case "Init Pretty Shutdown":
						Init_PrettyShutDown = true
						PrettyShutDownCfg()
					}
				}

			}
			os.Exit(0)
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	SignalString := yellow.Render("\n\nGolang Framework Initializer Tool ver 0.1.(Author: Mang Zhang, m13692277450@outlook.com)\n")
	SignalString += yellow.Render("\nHomePage: www.pavogroup.top\n")
	SignalString += cyan.Render("\nWhich initiallize function you want to setup?")
	SignalString = SignalString + "\n"
	SignalString += gray.Render("Use arrow keys to navigate, spacebar to select.")
	SignalString = SignalString + "\n\n"
	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = green.Render(">") // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = cyan.Render("x") // selected!
		}

		// Render the row
		SignalString += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	SignalString += "\nPress q to quit, press enter to confirm.\n"
	SignalString += "\n"
	SignalString += gold.Render(NewVersionIsAvailable)
	SignalString += "\n\n"

	// Send the UI for rendering
	return SignalString
}

func moveCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}
func clearScreen1() {
	fmt.Print("\033[2J\033[H")
}
func MainMenu() {
	clearScreen()
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("BubbleBubbletea, there's been an error: %v", err)
		os.Exit(1)
	}
}
