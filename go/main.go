package main

import (
	"net/http"
	"time"
)

var startTimeStamp time.Time = time.Now()

func main() {
	server := webServer{http.NewServeMux(), "localhost:9000"}
	go server.handleRoutes(&R)
	server.listenAndServe()
}
