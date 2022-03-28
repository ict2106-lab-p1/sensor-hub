package nono

import (
	"github.com/spf13/viper"
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
	Devices      []Device `mapstructure:"devices"`
	IntervalInMs int      `mapstructure:"interval_in_ms"`
}

type Device struct {
	Name string `mapstructure:"name"`
	Lab  string `mapstructure:"lab"`
}

type EnergyRequest struct {
	LabLocation    string  `json:"lab_location"`
	DeviceSerialNo string  `json:"device_serial_no"`
	Interval       int     `json:"interval"`
	EnergyUsage    float64  `json:"energy_usage"`
}

func ParseConfig() Config {
	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return config
}
