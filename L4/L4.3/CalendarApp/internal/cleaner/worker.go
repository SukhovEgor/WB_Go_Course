package cleaner

import (
	"calendar/internal/calendar"
	"calendar/pkg/logger"
	"time"
)

type Worker struct {
	service calendar.Service

	interval time.Duration
}

func New(
	service calendar.Service,
	interval time.Duration,
) *Worker {

	return &Worker{
		service: service,

		interval: interval,
	}
}

func (w *Worker) Start() {

	logger.InfoLog("Cleaner worker started")

	ticker := time.NewTicker(w.interval)

	defer ticker.Stop()

	for range ticker.C {

		w.service.ArchiveOldEvents()

		logger.InfoLog("Old events archived")
	}
}