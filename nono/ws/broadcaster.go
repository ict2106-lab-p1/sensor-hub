package ws

import (
	"time"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"

	"senkawa.moe/sensor-hub/nono/config"
)

type Action struct {
	DeviceName string `json:"device_name"`
	Event      string `json:"event"`
}

const (
	// LightStateFormat example: "1:light:state:1"
	LightStateFormat = "%s:light:state:%d"
	// LightTempFormat example: "1:light:state:2500"
	LightTempFormat = "%s:light:temp:%d"

	SpeakerActiveFormat = "%s:speaker:%d"

	MusicPauseFormat = "%s:speaker:pause:%d"
	MusicPlayFormat  = "%s:speaker:play:%d"
)

type Bしょうこ struct {
	Endpoint   string
	BaseClient *resty.Client
	Config     *config.Dispatch
	Log        *zap.SugaredLogger
	Incoming   <-chan Action
}

func Newしょうこ(cfg *config.Dispatch, log *zap.SugaredLogger) (*Bしょうこ, <-chan Action) {
	incoming := make(chan Action)

	client := resty.New()
	client.SetTimeout(2 * time.Second)

	return &Bしょうこ{BaseClient: client, Config: cfg, Log: log, Incoming: incoming}, incoming
}

func (s *Bしょうこ) StartDispatcher() {
	for incoming := range s.Incoming {
		go func(payload Action) {
			_, err := s.BaseClient.R().SetBody(payload).Get(s.Endpoint)
			if err != nil {
				s.Log.Warnw("automation backend returned error: %v", err)
			}
		}(incoming)
	}
}
