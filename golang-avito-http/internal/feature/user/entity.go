package user

import "github.com/google/uuid"

type DBUser struct {
	ID      uuid.UUID `db:"id"`
	Email   string    `db:"email"`
	Version uint64    `db:"version"`
	Role    string    `db:"role"`
}

type RequestRegister struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"required"`
}
type ResponseRegister struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"required"`
	Token    string `json:"token"`
}
