package router

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/websocket/v2"
	"go.uber.org/zap"

	"senkawa.moe/sensor-hub/nono/ws"
)

// Embed a single file
//go:embed index.html
var frontPage embed.FS

const outBoundBufferSize = 256

func RegisterWebsocketRoutes(app fiber.Router, hub *ws.Hub) {
	app.Get("/ws",
		func(c *fiber.Ctx) error {
			if websocket.IsWebSocketUpgrade(c) {
				return c.Next()
			}

			return fiber.ErrUpgradeRequired
		},
		websocket.New(func(c *websocket.Conn) {
			client := &ws.Client{Hub: hub, Conn: c.Conn, Send: make(chan []byte, outBoundBufferSize)}
			client.Hub.Register <- client

			go client.WritePump()
			client.ReadPump()
		}),
	)
}

func WebBuilder(log *zap.SugaredLogger, debug bool) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	go ws.Bus.Run()
	RegisterWebsocketRoutes(app, ws.Bus)

	if debug {
		log.Info("debugged index.html")
		app.Use("/", filesystem.New(filesystem.Config{
			Root: http.Dir("./nono/router"),
		}))
	} else {
		app.Use("/", filesystem.New(filesystem.Config{
			Root: http.FS(frontPage),
		}))
	}

	d := app.Group("/api/v1/:device")
	d.Get("/light/state/:state", func(c *fiber.Ctx) error {
		device := utils.CopyString(c.Params("device"))
		ws.Bus.Broadcast <- []byte("Lab " + device)

		log.Infow("light:state",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/light/brightness/:level", func(c *fiber.Ctx) error {
		log.Infow("light:brightness",
			"state", c.Params("level"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/locks/:state", func(c *fiber.Ctx) error {
		log.Infow("locks:state",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/hvac/state/:state", func(c *fiber.Ctx) error {
		log.Infow("hvac:state",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/hvac/temp/:temp", func(c *fiber.Ctx) error {
		log.Infow("hvac:temperature",
			"temp", c.Params("temp"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/camera/state/:state", func(c *fiber.Ctx) error {
		log.Infow("camera:state",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/camera/swing/:state", func(c *fiber.Ctx) error {
		log.Infow("camera:swing_state",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/camera/recording/:state", func(c *fiber.Ctx) error {
		log.Infow("camera:recording_state",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/speaker/state/:state", func(c *fiber.Ctx) error {
		log.Infow("speaker:state",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/speaker/playback/:state", func(c *fiber.Ctx) error {
		log.Infow("speaker:playback_state",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/speaker/volume/:state", func(c *fiber.Ctx) error {
		log.Infow("speaker:volume",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	return app
}

func ok(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok", "path": c.Path()})
}
