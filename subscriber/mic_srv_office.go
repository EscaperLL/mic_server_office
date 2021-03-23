package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	mic_srv_office "mic_srv_office/proto/mic_srv_office"
)

type Mic_srv_office struct{}

func (e *Mic_srv_office) Handle(ctx context.Context, msg *mic_srv_office.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *mic_srv_office.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
