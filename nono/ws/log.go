package ws

import (
	"go.uber.org/zap"

	"senkawa.moe/sensor-hub/nono"
)

var log *zap.SugaredLogger

func init() {
	log = nono.ConfigureLogger()
}
