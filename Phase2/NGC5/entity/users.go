package entity

type User struct {
	Id         int    `json:"id"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=8"`
	FullName   string `json:"fullname" validate:"required,min=6,max=15"`
	Age        int    `json:"age" validate:"required,min=17"`
	Occupation string `json:"occupation" validate:"required"`
	Role       string `json:"role"`
}
