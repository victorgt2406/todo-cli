package tdc

func Check(m model) {
	m.tasks[m.cursor].IsDone = !m.tasks[m.cursor].IsDone
	m.db.Save(&m.tasks[m.cursor])
}
