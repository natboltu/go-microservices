// mainn entry point broker service
package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	webPort = "80"
)

type Config struct{}

// one root , response (to accepted json payload, make sure connnect to frontend)
func main() {
	app := Config{}

	log.Printf("Starting broker servier on port %s\n", webPort)

	//define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	//start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
