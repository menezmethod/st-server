package models

import (
	"time"
)

type Journal struct {
	ID              uint64    `json:"id,omitempty" bun:",pk,autoincrement"`
	Name            string    `json:"name" validate:"required,min=1,max=255"`
	Description     string    `json:"description" validate:"required,min=1,max=1000"`
	CreatedAt       time.Time `json:"createdAt"`
	StartDate       string    `json:"startDate" validate:"required,datetime=2006-01-02"`
	EndDate         string    `json:"endDate" validate:"required,datetime=2006-01-02"`
	CreatedBy       uint64    `json:"createdBy" validate:"required"`
	UsersSubscribed []uint64  `json:"usersSubscribed"`
}
