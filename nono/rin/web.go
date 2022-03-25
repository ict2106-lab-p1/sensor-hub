package rin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"go.uber.org/zap"
)

func UnderTheDesk(log *zap.SugaredLogger, debug bool) *Rin {
	server := NewRin(log)
	server.RegisterWebsocketRoutes()
	server.RegisterIndexPage(debug)

	d := server.App.Group("/api/v1/:device")
	d.Get("/light/state/:state", func(c *fiber.Ctx) error {
		device := utils.CopyString(c.Params("device"))
		server.Hub.Broadcast <- []byte("Lab " + device)

		log.Infow("💡 light:state",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/light/brightness/:level", func(c *fiber.Ctx) error {
		log.Infow("💡 light:brightness",
			"state", c.Params("level"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/locks/:state", func(c *fiber.Ctx) error {
		log.Infow("🔒 locks:state",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/hvac/state/:state", func(c *fiber.Ctx) error {
		log.Infow("🌬️hvac:state",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/hvac/temp/:temp", func(c *fiber.Ctx) error {
		log.Infow("🌬️hvac:temperature",
			"temp", c.Params("temp"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/camera/state/:state", func(c *fiber.Ctx) error {
		log.Infow("📷 camera:state",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/camera/swing/:state", func(c *fiber.Ctx) error {
		log.Infow("📷 camera:swing_state",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/camera/recording/:state", func(c *fiber.Ctx) error {
		log.Infow("📷 camera:recording_state",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/speaker/state/:state", func(c *fiber.Ctx) error {
		log.Infow("🔊 speaker:state",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/speaker/playback/:state", func(c *fiber.Ctx) error {
		log.Infow("🔊 speaker:playback_state",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/speaker/volume/:state", func(c *fiber.Ctx) error {
		log.Infow("🔊 speaker:volume",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	return server
}

func ok(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok", "path": c.Path()})
}
