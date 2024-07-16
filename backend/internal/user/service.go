package user

import (
	"context"
	"strings"
	"time"

	"github.com/andikaraditya/horus-test/backend/internal/api"
	"github.com/andikaraditya/horus-test/backend/internal/db"
	"github.com/andikaraditya/horus-test/backend/internal/helper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
)

type UserService interface {
	CreateUser(user *User) error
	Login(req *LoginRequest) (string, error)
}

type srv struct {
	db db.DBService
}

var (
	Service UserService
)

func init() {
	Service = New(db.Service)
}

func New(db db.DBService) UserService {
	return &srv{db}
}

func (s *srv) CreateUser(user *User) error {
	var err error

	user.Password, err = helper.HashPassword(user.Password)
	if err != nil {
		return err
	}

	if err := s.db.Commit(nil, func(tx pgx.Tx) error {
		_, err := tx.Exec(
			context.Background(),
			`INSERT INTO "user" (
				nama,
				email,
				password,
				username
			) VALUES ($1, $2, $3, $4)`,
			user.Name,
			user.Email,
			user.Password,
			user.Username,
		)
		if err != nil {
			if strings.Contains(err.Error(), "23505") {
				return api.ErrPayload
			}
			return err
		}
		return nil
	}); err != nil {
		return err

	}
	return nil
}

func (s *srv) Login(req *LoginRequest) (string, error) {
	var hash string
	var userId string

	s.db.QueryRow(
		`SELECT password, id
			FROM "user"
			WHERE username = $1`,
		req.Username,
	).Scan(&hash, &userId)

	err := helper.ComparePassword(hash, req.Password)
	if err != nil {
		return "", api.ErrPayload
	}

	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}
