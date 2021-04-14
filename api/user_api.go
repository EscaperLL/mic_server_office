package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	//"github.com/micro/go-micro/v2/client"
	_ "github.com/micro/go-micro/v2/registry/etcd"
	micro_srv_user "mic_srv_office/proto/mic_srv_office"
	"mic_srv_office/api/handler"
)





func main()  {
	service :=micro.NewService(
		micro.Name("go.micro.api.user"))
	service.Init()
	service.Server().Handle(
		service.Server().NewHandler(
			&handler.IUser{Client: micro_srv_user.NewIUserService("go.micro.srv.user",service.Client())},
		),
	)
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}