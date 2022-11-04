package models

import "time"

type Journal struct {
	ID              uint64    `json:"id,omitempty" bun:",pk,autoincrement"`
	Name            string    `json:"name,omitempty"`
	Description     string    `json:"description,omitempty"`
	CreatedAt       time.Time `json:"createdAt,omitempty"`
	StartDate       time.Time `json:"startDate,omitempty"`
	EndDate         time.Time `json:"endDate,omitempty"`
	CreatedBy       string    `json:"createdBy,omitempty"`
	UsersSubscribed []uint64  `json:"usersSubscribed,omitempty"`
}
