package main

import (
	"errors"
	"log"
	"net/http"
	"sync"
	"time"
)

// Глобальный реестр для записи маршрутов
var R Register = Register{&sync.Mutex{}, make(chan Route), make(map[string]http.HandlerFunc)}

var startTimeStamp time.Time = time.Now()

func main() {
	server := webServer{http.NewServeMux(), "localhost:9000"}
	go server.handleRoutes(&R)
	server.listenAndServe()
}

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

// Класс реестра маршрутов
type Register struct {
	*sync.Mutex
	ch       chan Route
	handlers map[string]http.HandlerFunc
}

// Записывает путь к функции
func (r *Register) Register(path string, handler http.HandlerFunc) error {
	r.Lock()
	defer r.Unlock()
	if r.handlers[path] != nil {
		log.Println("path was tried to be registered twice: " + path)
		return errors.New("the path has been already registered")
	}
	r.handlers[path] = handler
	log.Println("path is registered: " + path)
	r.ch <- Route{path, handler}
	return nil
}

// Класс маршрута
type Route struct {
	path    string
	handler http.HandlerFunc
}
