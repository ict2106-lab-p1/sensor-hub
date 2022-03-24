package nono

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

func randDevice(devices []Device) Device {
	return devices[rand.Intn(len(devices))]
}

func randFloat(min, max int) float64 {
	return float64(min) + rand.Float64()*(float64(max)-float64(min))
}

func RunEnergyDispatcher(config Config, log *zap.SugaredLogger) {
	if !config.Energy.Active {
		return
	}

	devices := config.Energy.Devices

	client := resty.New()
	client.SetTimeout(1 * time.Second)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	ticker := shouldBeBetweenInterval(config.Energy.IntervalInMs)
	log.Infof("will dispatch requests to %v every %vms", config.Energy.Endpoint, config.Energy.IntervalInMs)
	for {
		select {
		case <-quit:
			fmt.Print("\n")
			log.Info("exiting ticker")
			os.Exit(1)
		case <-ticker.C:
			go dispatchRequestNow(config, devices, client, log)
		}
	}
}

// shouldBeBetweenInterval clamps the interval to at least 1000ms (1s)
func shouldBeBetweenInterval(intervalInMs int) *time.Ticker {
	if intervalInMs < 1000 {
		intervalInMs = 1000
	}

	return time.NewTicker(time.Duration(intervalInMs) * time.Millisecond)
}

func dispatchRequestNow(config Config, devices []Device, client *resty.Client, log *zap.SugaredLogger) {
	randValue := randFloat(config.Energy.Range.Start, config.Energy.Range.End)
	device := randDevice(devices)
	generated := &EnergyRequest{LabID: device.Lab, DeviceSerialNo: device.Name, Interval: config.Energy.IntervalInMs, EnergyUsage: randValue}
	pending := client.R().SetBody(generated)

	log.Infow("sent energy reading", "request", generated)
	_, err := pending.Post(config.Energy.Endpoint)
	if err == nil {
		log.Debugw("client responded", "request", generated)
	} else {
		log.Debugw("client err'd", "error", err)
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}
