package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	dispatcher "senkawa.moe/sensor-hub/app"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Boot up the web server. Also starts up the runners in the background.",
	Run: func(cmd *cobra.Command, args []string) {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		deflog, _ := config.Build()
		defer deflog.Sync()
		log := deflog.Sugar()

		dispatcher.Run(log)

		app := fiber.New(fiber.Config{
			DisableStartupMessage: true,
		})

		app.Get("/state", func(ctx *fiber.Ctx) error {
			return nil
		})

		d := app.Group("/api/v1/:device")
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

		log.Info("app listening on :8000")
		log.Fatal(app.Listen("127.0.0.1:8000"))
	},
}

func ok(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok", "path": c.Path()})
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("endpoint", "e", "", "Endpoint to POST data to.")
}
