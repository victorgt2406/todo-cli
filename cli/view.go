package cli

func (m model) View() string {
	switch m.viewContext {
	default:
		return m.viewTasks()
	}
}
