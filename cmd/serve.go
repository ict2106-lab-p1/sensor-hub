package cmd

import (
	"github.com/spf13/cobra"

	"senkawa.moe/sensor-hub/nono/router"

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

		go nono.RunEnergyDispatcher(nono.ParseConfig(), log)
		web := router.WebBuilder(log, debug)

		log.Infof("app listening on %v", host)
		log.Fatal(web.Listen(host))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("listen", "l", "localhost:8000", "Address to listen on")
	serveCmd.Flags().BoolP("debug", "d", false, "Serve index.html from local dir (instead of embedded)")
}
