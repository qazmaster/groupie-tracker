package main

import (
	"github.com/qazmaster/groupie-tracker/internal/app"
	"log"
)

func init() {
	if err := app.Run(); err != nil {
		log.Println(err.Error())
		return
	}
}

func main() {

}
