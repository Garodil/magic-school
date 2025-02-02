package main

import (
	"log"
	"net/http"
	"sync"
)

// Глобальный реестр для записи маршрутов
var R Register = Register{handlers: make(map[string]http.HandlerFunc)}

func main() {
	server := WebServer{http.NewServeMux(), "localhost:9000"}

	server.Handle("/", http.NotFoundHandler())
	server.HandleRoutes(R)
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
	log.Println("Server started at " + s.address)
	if err := http.ListenAndServe(s.address, s); err != nil {
		log.Fatalln(err)
	}
}

// Инициализирует все записанные маршруты
func (s *WebServer) HandleRoutes(r Register) {
	for path, handler := range r.handlers {
		s.HandleFunc(path, handler)
	}
}

// Записывает путь к функции
func (r *Register) Register(path string, handler http.HandlerFunc) {
	r.Lock()
	defer r.Unlock()
	r.handlers[path] = handler
}
