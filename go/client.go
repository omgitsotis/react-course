package main

import (
    "github.com/gorilla/websocket"
)

type Message struct {
    Name string `json:"name"`
    Data interface{} `json:"data"`
}

type FindHandler func(string) (Handler, bool)

type Client struct {
    send chan Message
    socket *websocket.Conn
    findHandler FindHandler
}

func (c *Client) Read() {
    var message Message
    for {
        if err := c.socket.ReadJSON(&message); err != nil {
            break
        }
        if fn, ok := c.findHandler(message.Name); ok {
            fn(c, message.Data)
        }
    }

    c.socket.Close()
}

func (c *Client) Write() {
    for msg := range c.send {
        if err := c.socket.WriteJSON(msg); err != nil {
            break
        }
    }

    c.socket.Close()
}

func NewClient(conn *websocket.Conn, fn FindHandler) *Client {
    return &Client {
        send: make(chan Message),
        socket: conn,
        findHandler: fn,
    }
}
