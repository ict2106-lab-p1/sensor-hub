package app

import (
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	Energy Energy `mapstructure:"energy"`
}

type Energy struct {
	Energy   []Device `mapstructure:"energy"`
	Active   bool     `mapstructure:"active"`
	Endpoint string   `mapstructure:"endpoint"`
	Range    struct {
		Start int `mapstructure:"start"`
		End   int `mapstructure:"end"`
	} `mapstructure:"range"`
	Devices []Device `mapstructure:"devices"`
}

type Device struct {
	Name string `mapstructure:"name"`
	Lab  int    `mapstructure:"lab"`
}

type EnergyRequest struct {
	LabID          int     `json:"lab_id"`
	DeviceSerialNo string  `json:"device_serial_no"`
	Interval       int     `json:"interval"`
	EnergyUsage    float64 `json:"energy_usage"`
}

func parseConfig() Config {
	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return config
}

func randDevice(devices []Device) Device {
	return devices[rand.Intn(len(devices))]
}

func randFloat(min, max int) float64 {
	return float64(min) + rand.Float64()*(float64(max)-float64(min))
}

func Run(log *zap.SugaredLogger) {
	config := parseConfig()
	devices := config.Energy.Devices

	client := resty.New()
	client.SetTimeout(1 * time.Second)

	ticker := time.NewTicker(1 * time.Second)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for {
			select {
			case <-c:
				os.Exit(1)
			case <-ticker.C:
				go func() {
					randValue := randFloat(config.Energy.Range.Start, config.Energy.Range.End)
					device := randDevice(devices)
					generated := &EnergyRequest{LabID: device.Lab, DeviceSerialNo: device.Name, Interval: 10, EnergyUsage: randValue}
					pending := client.R().
						SetBody(generated)

					log.Debugw("sent energy reading", "request", generated)
					_, err := pending.Post(viper.GetString("energy.endpoint"))
					if err == nil {
						log.Debugw("client responded", "request", generated)
					}
				}()
			}
		}
	}()
}

func init() {
	rand.Seed(time.Now().Unix())
}
