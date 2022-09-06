package models

import "time"

type User struct {
	Id             uint64    `json:"id" gorm:"primaryKey"`
	Email          string    `json:"email"`
	FullName       string    `json:"fullName"`
	Password       string    `json:"password"`
	Role           string    `json:"role"`
	TimeRegistered time.Time `json:"timeRegistered"`
}
