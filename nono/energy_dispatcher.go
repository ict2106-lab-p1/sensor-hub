package nono

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"

	"senkawa.moe/sensor-hub/nono/config"
)

type まゆ struct {
	BaseClient *resty.Client
	Config     *config.Energy
	Log        *zap.SugaredLogger
}

func Newまゆ(config *config.Energy, log *zap.SugaredLogger) *まゆ {
	client := resty.New()
	client.SetTimeout(1 * time.Second)

	return &まゆ{BaseClient: client, Config: config, Log: log}
}

func (n *まゆ) RunEnergyDispatcher() {
	if !n.Config.Active {
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	ticker := shouldBeBetweenInterval(n.Config.IntervalInMs)
	n.Log.Infof("will dispatch requests to %v every %vms", n.Config.Endpoint, n.Config.IntervalInMs)
	for {
		select {
		case <-quit:
			fmt.Print("\n")
			n.Log.Info("exiting ticker")
			os.Exit(1)
		case <-ticker.C:
			go n.dispatchRequestNow()
		}
	}
}

func (n *まゆ) dispatchRequestNow() {
	device := randDevice(n.Config.Devices)
	requestToDispatch := &config.EnergyRequest{
		LabLocation:    device.Lab,
		DeviceSerialNo: device.Name,
		Interval:       n.Config.IntervalInMs,
		EnergyUsage:    randFloat(n.Config.Range.Start, n.Config.Range.End),
	}

	pending := n.BaseClient.R().SetBody(requestToDispatch)

	n.Log.Infow("sent energy reading", "request", requestToDispatch)
	_, err := pending.Post(n.Config.Endpoint)
	if err == nil {
		n.Log.Debugw("client responded", "request", requestToDispatch)
	} else {
		n.Log.Debugw("client err'd", "error", err)
	}
}

// shouldBeBetweenInterval clamps the interval to at least 1000ms (1s)
func shouldBeBetweenInterval(intervalInMs int) *time.Ticker {
	if intervalInMs < 1000 {
		intervalInMs = 1000
	}

	return time.NewTicker(time.Duration(intervalInMs) * time.Millisecond)
}

func randDevice(devices []config.Device) config.Device {
	return devices[rand.Intn(len(devices))]
}

func randFloat(min, max int) float64 {
	return float64(min) + rand.Float64()*(float64(max)-float64(min))
}

func init() {
	rand.Seed(time.Now().Unix())
}
