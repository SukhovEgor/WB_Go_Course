package reminder

import (
	"calendar/internal/calendar"
	"calendar/pkg/logger"
	"fmt"
	"time"
)

type Worker struct {
	events <-chan calendar.Event
}

func New(events <-chan calendar.Event) *Worker {

	return &Worker{
		events: events,
	}
}

func (w *Worker) Start() {

	logger.InfoLog("Reminder worker started")

	for event := range w.events {

		go w.process(event)
	}
}

func (w *Worker) process(event calendar.Event) {

	wait := time.Until(event.Date)

	if wait > 0 {
		time.Sleep(wait)
	}

	logger.InfoLog(
		fmt.Sprintf(
			`Reminder: event "%s" for user %d`,
			event.Title,
			event.UserID,
		),
	)
}
