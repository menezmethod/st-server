package models

import (
	"time"
)

type Journal struct {
	CreatedAt       time.Time `json:"createdAt"`
	CreatedBy       uint64    `json:"createdBy" validate:"required"`
	Description     string    `json:"description" validate:"required,min=1,max=1000"`
	EndDate         string    `json:"endDate" validate:"required,datetime=2006-01-02"`
	ID              uint64    `json:"id,omitempty" bun:",pk,autoincrement"`
	LastUpdatedBy   uint64    `json:"lastUpdatedBy"`
	Name            string    `json:"name" validate:"required,min=1,max=255"`
	StartDate       string    `json:"startDate" validate:"required,datetime=2006-01-02"`
	UsersSubscribed []uint64  `json:"usersSubscribed"`
}
