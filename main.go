package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/reaperhero/go-kit-microservice/napodate"
)

var (
	httpAddr = flag.String("http", ":8080", "http listen address")
)

func main() {

	flag.Parse()
	ctx := context.Background()
	srv := napodate.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := napodate.Endpoints{
		GetEndpoint:      napodate.MakeGetEndpoint(srv),
		StatusEndpoint:   napodate.MakeStatusEndpoint(srv),
		ValidateEndpoint: napodate.MakeValidateEndpoint(srv),
	}

	go func() {
		log.Println("napodate is listening on port:", *httpAddr)
		handler := napodate.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
