package models

import (
	"time"
)

type User struct {
	Id             uint64    `json:"id" bun:",pk,autoincrement"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	FullName       string    `json:"fullName"`
	Role           string    `json:"role"`
	TimeRegistered time.Time `json:"timeRegistered"`
}
