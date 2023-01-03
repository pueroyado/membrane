package repositories

import (
	"demo/models"
	"errors"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) FindOne(id int32) (*models.User, error) {
	user := &models.User{
		Id:       1,
		Email:    "email",
		Password: "password",
	}

	return user, nil
}

func (r *UserRepo) FindByEmail(email string) (*models.User, error) {
	rowSql := "SELECT * FROM user WHERE user_email = ?"
	row := r.db.QueryRowx(rowSql, email)

	u := &models.User{}
	err := row.StructScan(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepo) Reg(email string, password string) (*models.JwtPayload, error) {
	isUser, _ := r.FindByEmail(email)
	if isUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	user := &models.User{
		Email:    email,
		Password: password,
		Status:   models.UserStatusActive,
	}
	user.GenerateHashPassword()

	row, dbErr := r.db.Exec(
		"INSERT INTO user (user_email, user_hash_password, user_status) VALUES (?, ?, ?)",
		user.Email, user.HashPassword, user.Status,
	)
	if dbErr != nil {
		return nil, errors.New("db exec, err: " + dbErr.Error())
	}
	if count, _ := row.RowsAffected(); count != 1 {
		return nil, errors.New("error inserting row value")
	}

	userId, _ := row.LastInsertId()
	jwtPayload := models.CreateJwt(int32(userId))

	return jwtPayload, nil
}

func (r *UserRepo) Auth(email string, password string) (*models.JwtPayload, error) {
	user, errFindUser := r.FindByEmail(email)
	if errFindUser != nil {
		return nil, errors.New("incorrect login or password")
	}

	errCompare := user.CompareHashAndPassword(password)
	if errCompare != nil {
		return nil, errors.New("incorrect login or password")
	}

	jwtPayload := models.CreateJwt(user.Id)

	return jwtPayload, nil
}
