package models

import (
	"time"
)

type User struct {
	Id        uint64    `json:"id" bun:",pk,autoincrement"`
	Bio       string    `json:"bio" validate:"max=1024"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"CreatedBy"`
	Email     string    `json:"email" validate:"required,email"`
	FirstName string    `json:"firstName" validate:"required,alpha"`
	LastName  string    `json:"lastName" validate:"required,alpha"`
	Password  string    `json:"password" validate:"required,gte=6"`
	Role      string    `json:"role" validate:"required,oneof=ADMIN USER TRADER"`
}
