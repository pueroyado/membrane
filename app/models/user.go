package models

import "golang.org/x/crypto/bcrypt"

const UserStatusActive int = 1

type User struct {
	Id           int32  `json:"id" db:"user_id"`
	Email        string `json:"email" db:"user_email"`
	Password     string `json:"-"`
	HashPassword string `json:"_" db:"user_hash_password"`
	Status       int    `json:"status" db:"user_status"`
}

type UserRepository interface {
	FindOne(id int32) (*User, error)
	FindByEmail(email string) (*User, error)
	Reg(email string, password string) (*JwtPayload, error)
	Auth(email string, password string) (*JwtPayload, error)
}

func (u *User) GenerateHashPassword() {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	u.HashPassword = string(hashPassword)
}

func (u *User) CompareHashAndPassword(password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(u.HashPassword),
		[]byte(password),
	)
}
