package rin

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"go.uber.org/zap"

	"senkawa.moe/sensor-hub/nono/ws"
)

type RinConfig struct {
	Log              *zap.SugaredLogger
	Hub              *ws.Hub
	Syoko            *ws.Bしょうこ
	Debug            bool
	OutboundMessages chan<- ws.Action
}

func UnderTheDesk(appConfig *RinConfig) *Rin {
	server := NewRin(appConfig.Log, appConfig.Hub)
	server.RegisterWebsocketRoutes()
	server.RegisterIndexPage(appConfig.Debug)

	log := appConfig.Log
	outbound := appConfig.OutboundMessages
	wsHub := server.Hub

	e := server.App.Group("/api/outbound")
	e.Get("/:payload", func(c *fiber.Ctx) error {
		payload := utils.CopyString(c.Params("payload"))

		outbound <- ws.Action{FiredAction: payload}

		return c.JSON(fiber.Map{"status": "ok"})
	})

	d := server.App.Group("/api/v1/:device")
	d.Use(func(c *fiber.Ctx) error {
		wsHub.LogBroadcast(c.Path())

		return c.Next()
	})

	d.Get("/light/state/:state", func(c *fiber.Ctx) error {
		device := utils.CopyString(c.Params("device"))
		state := utils.CopyString(c.Params("state"))
		wsHub.ActionBroadcast(
			fmt.Sprintf("light:%s:%s", device, state),
		)

		log.Infow("💡 light:state",
			"state", c.Params("state"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/light/brightness/:level", func(c *fiber.Ctx) error {
		device := utils.CopyString(c.Params("device"))
		level := utils.CopyString(c.Params("level"))
		wsHub.ActionBroadcast(
			fmt.Sprintf("brightness:%s:%s", device, level),
		)

		log.Infow("💡 light:brightness",
			"level", c.Params("level"),
			"device", c.Params("device"),
		)
		return ok(c)
	})

	d.Get("/light/temp/:level", func(c *fiber.Ctx) error {
		device := utils.CopyString(c.Params("device"))
		level := utils.CopyString(c.Params("level"))
		wsHub.ActionBroadcast(
			fmt.Sprintf("temp:%s:%s", device, level),
		)

		log.Infow("💡 light:temperature",
			"level", c.Params("level"),
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

	server.App.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("404.")
	})

	return server
}

func ok(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"status": "ok", "path": c.Path()})
}
