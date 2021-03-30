package handler

import (
	"mic_srv_office/model"
	mic_srv_office "mic_srv_office/proto/mic_srv_office"

	"context"
)

type mic_src_user struct {
}



func (u *mic_src_user)GetUsers(ctx context.Context, in *mic_srv_office.Users) (*mic_srv_office.Reply, error)  {
	usermodels :=model.GetUsers()
	model.ModelConvert2ProtoModel(&usermodels,in.GetUsersInfo())
	return &mic_srv_office.Reply{Code: 0,Msg: ""},nil
}

func (u *mic_src_user)GetUserByID(ctx context.Context, in  *mic_srv_office.UserIDs) (*mic_srv_office.User, error)  {
	usermodel :=model.GetUserByID(int64(in.GetID()[0]))
	users :=make([]model.UserModel,2)
	prtos :=make([]*mic_srv_office.User,2)
	users = append(users,*usermodel)
	model.ModelConvert2ProtoModel(&users,prtos)

	return prtos[0],nil
}

func (u *mic_src_user)SetUser(ctx context.Context, in *mic_srv_office.Users) (*mic_srv_office.Reply, error)  {

}

func (u *mic_src_user)SetUsers(ctx context.Context, in *mic_srv_office.Users) (*mic_srv_office.Reply, error)  {

}

func (u *mic_src_user)DelUser(ctx context.Context, in *mic_srv_office.Users) (*mic_srv_office.Reply, error)  {

}