package main

import (
    "github.com/gorilla/websocket"
    "net/http"
    "fmt"
)

type Handler func(*Client, interface{})

type Router struct {
    rules map[string]Handler
}

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
    CheckOrigin: func (r *http.Request) bool {return true},
}

func NewRouter() *Router {
    return &Router {
        rules : make(map[string]Handler),
    }
}

func (e *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    socket, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprint(w, err)
        return
    }

    client := NewClient(socket, e.FindHandler)
    go client.Write()
    client.Read()
}

func (r *Router) Handle(name string, fn Handler) {
    r.rules[name] = fn
}

func (r *Router) FindHandler(name string) (Handler, bool) {
    handler, ok := r.rules[name]
    return handler, ok
}
