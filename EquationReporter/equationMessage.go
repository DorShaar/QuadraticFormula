package EquationReporter

type EquationMessage struct {
	correlationId     string
	originalEquation  string
	organizedEquation string
	a                 int
	b                 int
	c                 int
	root1             float32
	root2             float32
	processTimeList   []processTime
}

func (equationMessage *EquationMessage) TryAddEvent(newEvent event.Event) bool {
	var shouldAddEvent bool = true
	for _, event := range schedule.eventsList {
		if event.EventTime.AreCoincide(newEvent.EventTime) {
			shouldAddEvent = false
			break
		}
	}

	if shouldAddEvent {
		schedule.eventsList = append(schedule.eventsList, newEvent)
	}

	return shouldAddEvent
}
