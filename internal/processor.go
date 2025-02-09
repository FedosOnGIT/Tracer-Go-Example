package internal

import (
	"errors"
	"github.com/FedosOnGIT/TracerLib/uploadBatch"
)

func FailingFunction(logger *uploadBatch.Logger) error {
	logger.Error("Failing function")
	return errors.New("failing function")
}
