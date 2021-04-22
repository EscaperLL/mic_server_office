package handler

import (
	"context"
	"encoding/json"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/v2/logger"
	mic_src "mic_srv_office/proto/mic_srv_office"
	"strconv"
)

type IUser struct {
	Client mic_src.IUserService
}


type Say struct {
	Client mic_src.IUserService
}

//http://localhost:8080/user/say/hello?name=sdgfh
func (s *Say) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	logger.Info("Received Say.Hello API request")

	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Name cannot be blank")
	}



	rsp.StatusCode = 200

	rsp.Body = string("12123")

	return nil
}



func (t *IUser)GetUsers(ctx context.Context,req *api.Request,rsp *api.Response) error{
	logger.Info("Received iuser.GetUsers request")

	users ,err:=t.Client.GetAllUser(ctx,&mic_src.ProtoRequest{})

	if err != nil {
		rsp.StatusCode=500
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

func (t *IUser)GetUserByID(ctx context.Context,req *api.Request,rsp *api.Response) error{
	id,ok :=req.Get["ID"]
	if !ok {
		return errors.BadRequest("id can not be empty","")
	}
	idNum,_ :=strconv.Atoi(id.Values[0])
	ids :=mic_src.UserIDs{}
	ids.ID=append(ids.ID,int32(idNum))
	user,err :=t.Client.GetUserByID(ctx,&ids)
	if err !=nil {
		rsp.StatusCode=500
		return err
	}
	rsp.StatusCode=200
	user_bytes,_ :=json.Marshal(user)
	body ,_ := json.Marshal(map[string]string{
		"user":string(user_bytes[:]),
	})
	rsp.Body=string(body)
	return nil
}

func (t *IUser)SetUser(ctx context.Context,req *api.Request,rsp *api.Response) error{
	in_user,ok := req.Get["user"]
	if !ok {
		return errors.BadRequest("user can not be empty","")
	}
 	user:=mic_src.User{}
	err :=json.Unmarshal([]byte(in_user.Values[0]),&user)
	if err != nil{
		logger.Info(in_user.Values)
		return errors.BadRequest("json Unmarshal failed","")
	}
	_,err_setuser := t.Client.SetUser(ctx,&user)
	if  err_setuser!=nil{
		logger.Info(err_setuser)
		return errors.BadRequest("set user failed","")
	}
	rsp.StatusCode=200
	return nil
}


func (t *IUser)DelUser(ctx context.Context,req *api.Request,rsp *api.Response) error{
	ids,ok := req.Get["delete_ids"]
	if !ok {
		return errors.BadRequest("user can not be empty","")
	}
	users :=mic_src.UserIDs{}
	for _,v :=range ids.Values{
		idNum,_ :=strconv.Atoi(v)
		users.ID=append(users.ID,int32(idNum))
	}
	_,err:=t.Client.DelUser(ctx,&users)
	if  err!=nil{
		logger.Info(err)
		return errors.BadRequest("set user failed","")
	}
	rsp.StatusCode=200
	return nil
}