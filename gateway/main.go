package main

import (
	"log"
	"net/http"

	"go-kit-demo/gateway/internal/service"
	"go-kit-demo/gateway/internal/transport"
)

func main() {
	var (
		svc           = service.NewAddService()
		sumHandler    = transport.NewSumHandler(svc)
		concatHandler = transport.NewConcatHandler(svc)
	)
	http.Handle("/sum", sumHandler)
	http.Handle("/concat", concatHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
