package cli

import tasksPresenter "todo-cli/presenters/tasksPresenter"

func (m model) View() string {
	switch m.viewContext {
	default:
		return m.tasksPresenter.Render(tasksPresenter.RenderProps{
			ViewContext: m.viewContext,
			Cursor:      m.cursor,
			Tasks:       m.getTasks(),
			TextInput:   m.textInput,
		})
	}
}
