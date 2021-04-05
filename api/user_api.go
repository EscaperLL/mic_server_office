package main

import (
	"github.com/micro/go-micro/v2"

	//"github.com/micro/go-micro/v2/client"
	_ "github.com/micro/go-micro/v2/registry/etcd"
	"encoding/json"
	"log"
"context"
	api "github.com/micro/micro/api/proto"
	mic_src "mic_srv_office/proto/mic_srv_office"
)

type IUser struct {
	Client mic_src.IUserService
}

//func (t *IUser)GetUsers(ctx context.Context,req *api.Request,rsp *api.Response) error{
//	log.Println("get Users API ---------------")
//	users :=mic_src.Users{}
//	_,err :=t.Client.GetUsers(ctx,&users)
//	if err != nil {
//		return err
//	}
//	rsp.StatusCode=200
//	user_bytes,_ := json.Marshal(users)
//	body ,_ := json.Marshal(map[string]string{
//		"users":string(user_bytes[:]),
//	})
//	rsp.Body=string(body)
//	return nil
//}

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
		&IUser{Client: mic_src.NewIUserService("go.micro.src.mic_srv_office",service.Client())})
	if err :=service.Run();err !=nil {
		log.Fatal(err)
	}
}