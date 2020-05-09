package main

import (
	"log"
	"net/http"

	"github.com/sh3rp/stringen"
)

func main() {
	stringen.LOGGER.Info().Msgf("StrinGen v1.0")
	svc := stringen.NewService(http.NewServeMux())
	log.Fatal(svc.Serve(":8888"))
}
