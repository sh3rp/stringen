package main

import (
	log2 "github.com/sh3rp/stringen/pkg/log"
	"github.com/sh3rp/stringen/pkg/server"
	"log"
	"net/http"
)

func main() {
	log2.LOGGER.Info().Msgf("StrinGen v1.0")
	svc := server.NewService(http.NewServeMux())
	log.Fatal(svc.Serve(":8888"))
}
