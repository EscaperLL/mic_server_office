package handler

import (
	"mic_srv_office/model"
	mic_srv_office "mic_srv_office/proto/mic_srv_office"

	"context"
)

type Mic_src_user struct {
}



func (u *Mic_src_user)GetUsers(ctx context.Context, in *mic_srv_office.Users, reply *mic_srv_office.Reply) error  {
	usermodels :=model.GetUsers()
	model.ModelsConvert2ProtoModels(&usermodels,in.GetUsersInfo())
	reply=&mic_srv_office.Reply{Code: 0,Msg: ""}
	return nil
}

func (u *Mic_src_user)GetUserByID( ctx context.Context,in *mic_srv_office.UserIDs,micuser *mic_srv_office.User) error {
	usermodel :=model.GetUserByID(int64(in.GetID()[0]))
	mic_user:=mic_srv_office.User{}
	model.ModelConver2ProtoModel(*usermodel,&mic_user)
	micuser=&mic_user
	return nil
}

func (u *Mic_src_user)SetUser(ctx context.Context,in *mic_srv_office.User,re *mic_srv_office.Reply) error  {
	usermodel := model.UserModel{}

	model.ProtoModel2UserModel(in,&usermodel)
	err :=model.AddUser(usermodel)
	if err != nil{
		return  err
	}
	return nil
}

func (u *Mic_src_user)SetUsers(cxt context.Context,in *mic_srv_office.Users,re *mic_srv_office.Reply) error {
	var usermodels []*model.UserModel
	proInNum := len(in.UsersInfo)
	usermodels =make([]*model.UserModel,proInNum)
	model.ProtoModels2Usermodes(in.UsersInfo,usermodels)
	err := model.AddUsers(usermodels)
	if nil != err{
		 return err
	}
	return nil

}

func (u *Mic_src_user)DelUser(ctx context.Context,in *mic_srv_office.UserIDs,re *mic_srv_office.Reply) error  {
	var usermodels []*model.UserModel
	proInNum := len(in.ID)
	usermodels =make([]*model.UserModel,proInNum)
	userprotos :=make([]*mic_srv_office.User,proInNum)
	for _,id := range in.ID{
		userprotos=append(userprotos,&mic_srv_office.User{ID: id})
	}
	model.ProtoModels2Usermodes(userprotos,usermodels)
	err := model.DeleteUsers(usermodels)
	return err
}