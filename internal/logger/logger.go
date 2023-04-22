package logger

import (
	"fmt"
	"github.com/qazmaster/groupie-tracker/internal/config"
	"log"
	"os"
)

type Logger struct {
	log *log.Logger
}

func New(cfg *config.Config) (*Logger, error) {
	file, err := os.OpenFile(cfg.Logger.OutputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening log file: %v", err)
	}

	logger := log.New(file, "", log.LstdFlags)
	logger.SetPrefix(cfg.Logger.Level + " ")

	return &Logger{log: logger}, nil
}
