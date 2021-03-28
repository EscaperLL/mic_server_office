package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/v2"

	//"github.com/micro/go-micro"
	api "github.com/micro/go-micro/v2/api/proto"
	//"github.com/micro/go-micro/v2/client"
	"log"

	mic_src "mic_srv_office/proto/mic_srv_office"
)

type IUser struct {
	Client mic_src.IUserService
}

func (t *IUser)GetUsers(ctx context.Context,req *api.Request,rsp *api.Response) error{
	log.Println("get Users API ---------------")
	users :=mic_src.Users{}
	_,err :=t.Client.GetUsers(ctx,&users)
	if err != nil {
		return err
	}
	rsp.StatusCode=200
	user_bytes,_ := json.Marshal(users)
	body ,_ := json.Marshal(map[string]string{
		"users":string(user_bytes[:]),
	})
	rsp.Body=string(body)
	return nil
}

func main()  {
	service :=micro.NewService(
		micro.Name("go.micro.api.user"))
	service.Init()
	service.Server().NewHandler(
		&IUser{Client: mic_src.NewIUserService("go.micro.srv.users",service.Client())})
	if err :=service.Run();err !=nil {
		log.Fatal(err)
	}
}