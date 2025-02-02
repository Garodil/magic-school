package main

import (
	"log"
	"net/http"
)

// Сервер
var H host = host{http.NewServeMux(), "localhost:9000"}

// Массив с записанными маршрутами
var R map[string]http.HandlerFunc = make(map[string]http.HandlerFunc)

func main() {
	HandleRoutes()
	Host()
}

// Класс сервера
type host struct {
	*http.ServeMux        // Защита маршрутов
	ip             string // Адрес сервера
}

// Запускает сервер
func Host() {
	log.Println("Hosted")
	if err := http.ListenAndServe(H.ip, H); err != nil {
		log.Fatalln(err)
	}
}

// Задаёт все незаданные пути как 404
func NotFound() {
	Route("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("no method"))
	})
}

// Инициализирует все записанные маршруты
func HandleRoutes() {
	NotFound()
	for path, handler := range R {
		H.HandleFunc(path, handler)
	}
}

// Записывает путь к функции
func Route(path string, handler http.HandlerFunc) {
	R[path] = handler
}
