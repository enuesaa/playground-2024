package main

import (
	"context"
	"log"
	"os"
	"runtime/trace"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// websocket のコネクションたちを保持する
func NewWsConns() WsConns {
	return WsConns {
		conns: make(map[string]*websocket.Conn),
	}
}

type WsConns struct {
	conns map[string]*websocket.Conn
}
func (c *WsConns) Add(conn *websocket.Conn) string {
	id := uuid.NewString()
	c.conns[id] = conn
	return id
}
func (c *WsConns) Remove(id string) {
	if conn, ok := c.conns[id]; ok {
		delete(c.conns, id)
		conn.Close()
	}
}
func (c *WsConns) Send(message string) {
	for id, conn := range c.conns {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			c.Remove(id)
		}
	}
}

func main() {
	// setup tracer
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	defer f.Close()
	if err := trace.Start(f); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	defer trace.Stop()
	_, task := trace.NewTask(context.Background(), "main")
	defer task.End()

	// fiber app
	app := fiber.New()

	// websocket
	wsconns := NewWsConns()
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		requestId := wsconns.Add(c)
		defer wsconns.Remove(requestId)

		for {
			messageType, _, err := c.ReadMessage()
			if err != nil {
				break
			}
			if messageType == websocket.CloseGoingAway {
				break
			}
		}
	}))
	go func() {
		for {
			time.Sleep(1 * time.Second)
			wsconns.Send("hello")
		}
	}()

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

