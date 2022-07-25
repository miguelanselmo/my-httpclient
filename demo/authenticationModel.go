package demo

type Authentication struct {
	UserId   int    `json:"-"`
	UserName string `json:"name" validate:"required,min=3,max=255"`
	Password string `json:"password" validate:"required,password"`
}
