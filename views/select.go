package views

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func SelectFromList(items []string) (string, error) {
	indexedItems := make(map[int]string)
	for i, item := range items {
		indexedItems[i] = item
	}

	_, selectedItem, err := SelectFromMap(indexedItems)
	if err != nil {
		return "", fmt.Errorf("error selecting item: %v", err)
	}

	return selectedItem, nil
}

func SelectFromMap(items map[int]string) (int, string, error) {
	p := tea.NewProgram(initialModel(items))
	m, err := p.Run()
	if err != nil {
		return 0, "", fmt.Errorf("error running program: %v", err)
	}

	if m, ok := m.(model); ok && m.selected != -1 {
		selectedKey := m.keys[m.selected]
		return selectedKey, items[selectedKey], nil
	}
	return 0, "", fmt.Errorf("no item selected")
}

type model struct {
	items    map[int]string
	keys     []int
	cursor   int
	selected int
}

func initialModel(items map[int]string) model {
	keys := make([]int, 0, len(items))
	for k := range items {
		keys = append(keys, k)
	}
	return model{
		items:    items,
		keys:     keys,
		selected: -1,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.keys)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = m.cursor
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Select an item:\n\n"

	for i, key := range m.keys {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, m.items[key])
	}

	s += "\nPress q to quit.\n"
	return s
}
