package models

import (
	"time"
)

type User struct {
	Id        uint64    `json:"id" bun:",pk,autoincrement"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Bio       string    `json:"bio"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}
