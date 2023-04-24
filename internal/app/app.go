package app

import (
	"fmt"
	"github.com/qazmaster/groupie-tracker/pkg/config"
	"github.com/qazmaster/groupie-tracker/pkg/logger"
	"log"
)

func Run() error {
	// config
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("cannot init config: %w", err)
	}
	// logger
	l, err := logger.New(cfg)
	if err != nil {
		return fmt.Errorf("cannot init logger: %w", err)
	}
	defer func() {
		err = log.Output(2, "")
		if err != nil {
			log.Println(err)

		}
	}()
	l = l
	// TODO: continue here

	return nil
}
