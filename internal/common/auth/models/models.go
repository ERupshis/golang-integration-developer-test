package models

import (
	pb_auth "github.com/erupshis/golang-integration-developer-test/pb/auth"
)

type User struct {
	ID       int64
	Login    string
	Password string
}

func ConvertUserFromGRPC(in *pb_auth.Creds) *User {
	return &User{
		Login:    in.GetLogin(),
		Password: in.GetPassword(),
	}
}

func ConvertUserToGRPC(in *User) *pb_auth.Creds {
	return &pb_auth.Creds{
		Login:    in.Login,
		Password: in.Password,
	}
}
