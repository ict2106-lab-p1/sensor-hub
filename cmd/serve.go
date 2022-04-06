package cmd

import (
	"github.com/spf13/cobra"

	"senkawa.moe/sensor-hub/nono/config"
	"senkawa.moe/sensor-hub/nono/rin"
	"senkawa.moe/sensor-hub/nono/ws"

	"senkawa.moe/sensor-hub/nono"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Boot up the web server. Also starts up the runners in the background.",
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("listen")
		debug, _ := cmd.Flags().GetBool("debug")

		log := nono.ConfigureLogger()
		cfg := config.ParseConfig()

		hub := ws.NewHub(log)
		go hub.Run()

		energy := nono.Newまゆ(&cfg.Energy, log)
		go energy.RunEnergyDispatcher()

		smartDeviceEvents, outboundMessages := ws.Newしょうこ(&cfg.Dispatch, log)
		go smartDeviceEvents.StartDispatcher()

		server := rin.UnderTheDesk(&rin.RinConfig{
			Log:              log,
			Hub:              hub,
			Syoko:            smartDeviceEvents,
			Debug:            debug,
			OutboundMessages: outboundMessages,
		})

		log.Infof("app listening on %v", host)
		log.Fatal(server.App.Listen(host))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("listen", "l", "localhost:8000", "Address to listen on")
	serveCmd.Flags().BoolP("debug", "d", false, "Serve index.html from local dir (instead of embedded)")
}
