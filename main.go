package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"senkawa.moe/sensor-hub/cmd"
)

const defaultConfig = `energy:
  interval_in_ms: 5000
  active: true
  endpoint: http://localhost:5000/api/energylog/log
  range:
    start: 100
    end: 1000
  devices:
    - name: SC1001
      lab: NYP-SR7C
    - name: R1001
      lab: NYP-SR7C
    - name: S1001
      lab: NYP-SR7C
dispatch:
  endpoint: http://localhost:5050/sensor/dispatch
`

func main() {
	if _, err := os.Stat("config.yaml"); err != nil {
		fmt.Println("config generated, edit your config and re-run the application")
		if err := os.WriteFile("config.yaml", []byte(defaultConfig), 0o644); err != nil {
			panic(err)
			return
		}

		os.Exit(0)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w \n", err))
	}

	cmd.Execute()
}
