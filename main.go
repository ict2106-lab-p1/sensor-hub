package main

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	deflog, _ := config.Build()
	defer deflog.Sync()
	log := deflog.Sugar()

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/state", func(ctx *fiber.Ctx) error {
		return nil
	})

	d := app.Group("/:device")
	d.Get("/light/state/:state", func(c *fiber.Ctx) error {
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

	log.Info("app listening on :3000")
	log.Fatal(app.Listen("127.0.0.1:8000"))
}

func ok(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok", "path": c.Path()})
}
