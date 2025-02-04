package main

import (
	"errors"
	"log"
	"net/http"
	"sync"
	"time"
)

// Глобальный реестр для записи маршрутов
var R Register = Register{&sync.Mutex{}, make(map[string]http.HandlerFunc)}

var startTimeStamp time.Time = time.Now()

func main() {
	server := WebServer{http.NewServeMux(), "localhost:9000"}

	go server.HandleRoutes(R)
	server.ListenAndServe()
}

// Класс веб-сервера
type WebServer struct {
	*http.ServeMux        // Защита маршрутов
	address        string // Адрес сервера
}

// Класс реестра маршрутов
type Register struct {
	*sync.Mutex
	handlers map[string]http.HandlerFunc
}

// Запускает сервер
func (s *WebServer) ListenAndServe() {
	log.Printf("Server started at "+s.address+" in %s\n", time.Since(startTimeStamp))
	if err := http.ListenAndServe(s.address, s); err != nil {
		log.Fatalln(err)
	}
}

// Инициализирует все записанные маршруты
func (s *WebServer) HandleRoutes(r Register) {
	for path, handler := range r.handlers {
		go func() {
			s.HandleFunc(path, handler)
			log.Println(path + " path handled")
		}()
	}
}

// Записывает путь к функции
func (r *Register) Register(path string, handler http.HandlerFunc) error {
	r.Lock()
	defer r.Unlock()
	if r.handlers[path] != nil {
		log.Println(path + " was registered twice!")
		return errors.New(path + " path has been already registered!")
	}
	r.handlers[path] = handler
	log.Println(path + " path has registered")
	return nil
}
