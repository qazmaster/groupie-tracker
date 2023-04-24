package error

import (
	"fmt"
	"log"
	"runtime"
)

type ErrorHandler interface {
	HandleError(err error)
}

type LoggingErrorHandler struct{}

func getFileAndLine() (string, int) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return "unknown", 0
	}
	return fmt.Sprintf("%s:%d", file, line), int(pc)
}

func (h *LoggingErrorHandler) HandleError(err error) {
	file, line := getFileAndLine()
	log.Output(2, fmt.Sprintf("%s:%d: %v", file, line, err))
}
