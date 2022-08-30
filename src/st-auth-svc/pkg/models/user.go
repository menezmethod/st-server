package models

import "time"

type User struct {
	Id             uint64    `json:"id" gorm:"primaryKey"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	FullName       string    `json:"fullName"`
	TimeRegistered time.Time `json:"timeRegistered"`
}
