package main

import (
	"github.com/Vova345345/psichoterapyclinic-app"
	"log"
)

func main() {
	srv := new(psichoterapyclinic.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
