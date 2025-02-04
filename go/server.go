package main

import (
	"log"
	"net/http"
	"time"
)

// Класс веб-сервера
type webServer struct {
	*http.ServeMux        // Защита маршрутов
	address        string // Адрес сервера
}

// Запускает сервер
func (s *webServer) listenAndServe() {
	since := time.Since(startTimeStamp)
	log.Printf("Server started at "+s.address+" in %s\n", since)
	if err := http.ListenAndServe(s.address, s); err != nil {
		log.Fatalln(err)
	}
}

// Инициализирует маршруты, запускать как Горутину
func (s *webServer) handleRoutes(r *Register) {
	for {
		route := <-r.ch
		go func() {
			s.HandleFunc(route.path, route.handler)
			log.Println("[MAIN] a path is set on: " + route.path)
		}()
	}
}
