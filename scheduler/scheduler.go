package scheduler

import (
	"TracerExample/internal"
	"github.com/FedosOnGIT/TracerLib/uploadBatch"
	"time"
)

func StartBackgroundTask(logger *uploadBatch.Logger) {
	backgroundLogger := logger.WithTag("worker", "background")
	backgroundLogger.Warn("Starting background task")

	ticker := time.NewTimer(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			backgroundLogger.Warn("Background task")
			err := internal.FailingFunction(backgroundLogger)
			backgroundLogger.Errorf("Background task failed: %v", err)
			ticker.Reset(1 * time.Minute)
		}
	}
}
