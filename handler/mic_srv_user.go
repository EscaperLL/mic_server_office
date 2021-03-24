package handler

import (
	mic_srv_office "mic_srv_office/proto/mic_srv_office"

	"context"
)

type mic_src_user struct {
}



func (u *mic_src_user)GetUsers(ctx context.Context, in *mic_srv_office.Users) (*mic_srv_office.Reply, error)  {

}

func (u *mic_src_user)GetUserByID(ctx context.Context, in  *mic_srv_office.UserIDs) (*mic_srv_office.User, error)  {

}

func (u *mic_src_user)SetUser(ctx context.Context, in *mic_srv_office.Users) (*mic_srv_office.Reply, error)  {

}

func (u *mic_src_user)SetUsers(ctx context.Context, in *mic_srv_office.Users) (*mic_srv_office.Reply, error)  {

}

func (u *mic_src_user)DelUser(ctx context.Context, in *mic_srv_office.Users) (*mic_srv_office.Reply, error)  {

}