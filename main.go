package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"mic_srv_office/handler"
	"mic_srv_office/subscriber"

	mic_srv_office "mic_srv_office/proto/mic_srv_office"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.mic_srv_office"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	mic_srv_office.RegisterMic_srv_officeHandler(service.Server(), new(handler.Mic_srv_office))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.mic_srv_office", service.Server(), new(subscriber.Mic_srv_office))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
