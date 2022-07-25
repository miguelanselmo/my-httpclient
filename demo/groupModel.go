package demo

import (
	"time"
)

//Entity Group
type Group struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Users     []User    `json:"usrs"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
