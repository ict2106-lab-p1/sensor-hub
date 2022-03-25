package rin

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/websocket/v2"
	"go.uber.org/zap"

	"senkawa.moe/sensor-hub/nono/ws"
)

//go:embed index.html
var chaika embed.FS

type Rin struct {
	App *fiber.App
	Log *zap.SugaredLogger
	Hub *ws.Hub
}

func NewRin(log *zap.SugaredLogger) *Rin {
	return &Rin{
		App: fiber.New(fiber.Config{
			DisableStartupMessage: true,
		}),
		Log: log,
		Hub: ws.NewHub(),
	}
}

const outBoundBufferSize = 256

func (r *Rin) RegisterWebsocketRoutes() {
	go r.Hub.Run()

	r.App.Get("/ws",
		func(c *fiber.Ctx) error {
			if websocket.IsWebSocketUpgrade(c) {
				return c.Next()
			}

			return fiber.ErrUpgradeRequired
		},
		websocket.New(func(c *websocket.Conn) {
			client := &ws.Client{Hub: r.Hub, Conn: c.Conn, Send: make(chan []byte, outBoundBufferSize), Log: r.Log}
			client.Hub.Register <- client

			go client.WritePump()
			client.ReadPump()
		}),
	)
}

func (r *Rin) RegisterIndexPage(debug bool) {
	if debug {
		r.Log.Info("debugged index.html")
		r.App.Use("/", filesystem.New(filesystem.Config{
			Root: http.Dir("./nono/rin"),
		}))
	} else {
		r.App.Use("/", filesystem.New(filesystem.Config{
			Root: http.FS(chaika),
		}))
	}
}
