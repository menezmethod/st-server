package utils

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"st-auth-svc/pkg/models"
	"st-auth-svc/pkg/pb"
)

func ToUser(in models.User) *pb.User {

	out := new(pb.User)
	out.Id = in.Id
	if in.Email.Valid {
		out.Email = in.Email.String
	}
	if in.Password.Valid {
		out.Password = in.Password.String
	}
	if in.FullName.Valid {
		out.FullName = in.FullName.String
	}
	if in.Role.Valid {
		out.Role = in.Role.String
	}
	if in.TimeRegistered.Valid {
		out.TimeRegistered = timestamppb.New(in.TimeRegistered.Time)
	}
	return out
}
