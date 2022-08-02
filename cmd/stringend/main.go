package main

import (
	"flag"
	"fmt"
	log2 "github.com/sh3rp/stringen/pkg/log"
	"github.com/sh3rp/stringen/pkg/server"
	"log"
	"net/http"
)

var port int

func main() {
	log2.LOGGER.Info().Msgf("StrinGen v1.1")
	flag.IntVar(&port, "p", 8888, "Port to run on")
	flag.Parse()
	svc := server.NewService(http.NewServeMux())
	log.Fatal(svc.Serve(fmt.Sprintf(":%d", port)))
}
