package main

import (
	"errors"
	"log"
	"net/http"
	"sync"
)

// Глобальный реестр для записи маршрутов
var R Register = Register{&sync.Mutex{}, make(chan Route), make(map[string]http.HandlerFunc)}

// Класс реестра маршрутов
type Register struct {
	*sync.Mutex
	ch       chan Route
	handlers map[string]http.HandlerFunc
}

// Класс маршрута
type Route struct {
	path    string
	handler http.HandlerFunc
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
