package demo

import (
	"time"
)

//Entity User
type User struct {
	Id        int       `json:"id" uri:"id"`
	Name      string    `json:"name" validate:"required,min=3,max=255"`
	Password  string    `json:"-"`
	Email     string    `json:"email" validate:"required,email"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	GroupName string    `json:"group_name"`
}
