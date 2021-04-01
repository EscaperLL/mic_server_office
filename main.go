package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"mic_srv_office/handler"
	"mic_srv_office/subscriber"

	_ "github.com/micro/go-micro/v2/registry/etcd"
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
	mic_srv_office.RegisterMicSrvOfficeHandler(service.Server(), new(handler.Mic_srv_office))
	mic_srv_office.RegisterIUserServiceHandler(service.Server(),new(handler.Mic_src_user))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.mic_srv_office", service.Server(), new(subscriber.Mic_srv_office))
	micro.RegisterSubscriber("go.micro.service.mic_srv_office_user", service.Server(), new(subscriber.Mic_srv_office))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
