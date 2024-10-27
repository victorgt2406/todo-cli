package load

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Process struct {
	Name     string
	Complete bool
}

type finishMsg string

func Load(processes []string) func(string) {
	done := make(chan struct{})
	p := tea.NewProgram(initialModel(processes))

	// Run the program in a goroutine
	go func() {
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error running program: %v", err)
		}

		// Signal that the program has finished
		close(done)
	}()

	// Return the function that will be used to finish processes
	finishFn := func(processName string) {
		p.Send(finishMsg(processName))

		// Wait a bit for the spinner to update
		time.Sleep(50 * time.Millisecond)
	}

	// Start a cleanup goroutine
	go func() {
		// Wait for the program to finish
		<-done

		// Give time for the last update to be visible
		time.Sleep(100 * time.Millisecond)

		// Clear the loader output
		fmt.Print("\033[2K\033[1A") // Clear current line and move up
		for range processes {
			fmt.Print("\033[2K\033[1A") // Clear each process line and move up
		}
		fmt.Print("\033[2K\033[1A") // Clear the extra newline
	}()

	return finishFn
}

type model struct {
	processes []Process
	spinners  map[string]spinner.Model
}

func initialModel(processNames []string) model {
	processes := make([]Process, len(processNames))
	spinners := make(map[string]spinner.Model)

	for i, name := range processNames {
		processes[i] = Process{Name: name, Complete: false}
		s := spinner.New()
		s.Spinner = spinner.Dot
		s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
		spinners[name] = s
	}

	return model{
		processes: processes,
		spinners:  spinners,
	}
}

func (m model) Init() tea.Cmd {
	var cmds []tea.Cmd
	for _, process := range m.processes {
		if !process.Complete {
			cmds = append(cmds, m.spinners[process.Name].Tick)
		}
	}
	return tea.Batch(cmds...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" || msg.String() == "esc" {
			return m, tea.Quit
		}

	case finishMsg:
		processName := string(msg)
		for i, process := range m.processes {
			if process.Name == processName {
				m.processes[i].Complete = true
			}
		}

		// Check if all processes are complete
		allComplete := true
		for _, process := range m.processes {
			if !process.Complete {
				allComplete = false
				break
			}
		}
		if allComplete {
			return m, tea.Quit
		}

	default:
		var cmds []tea.Cmd
		for name, s := range m.spinners {
			// Only update spinners for incomplete processes
			for _, process := range m.processes {
				if process.Name == name && !process.Complete {
					var cmd tea.Cmd
					s, cmd = s.Update(msg)
					m.spinners[name] = s
					cmds = append(cmds, cmd)
				}
			}
		}
		return m, tea.Batch(cmds...)
	}

	return m, nil
}

func (m model) View() string {
	var s string
	s += "\n"

	for _, process := range m.processes {
		if !process.Complete {
			spinner := m.spinners[process.Name]
			s += fmt.Sprintf(" %s Loading %s...\n", spinner.View(), process.Name)
		} else {
			s += fmt.Sprintf(" ✓ Completed %s\n", process.Name)
		}
	}

	s += "\n"
	return s
}
