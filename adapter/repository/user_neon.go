package repository

import (
	"context"
	"go_cleanArchitecture_study/domain"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserNeon struct {
	db *pgxpool.Pool
}

func NewUserNeon(db *pgxpool.Pool) *UserNeon {
	return &UserNeon{
		db: db,
	}
}

func (r *UserNeon) Create(user domain.User) (domain.User, error) {

	query := `INSERT INTO users (id, name, email, created_at) VALUES ($1, $2, $3, $4)`

	_, err := r.db.Exec(
		context.Background(),
		query,
		user.ID(),
		user.Name(),
		user.Email(),
		user.CreatedAt(),
	)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *UserNeon) FindByID(id domain.UserID) (domain.User, error) {
	row := r.db.QueryRow(
		context.Background(),
		`SELECT id, name, email, created_at FROM users WHERE id = $1`,
		id,
	)

	var (
		uid 		string
		name		string
		email		string
		createdAt	time.Time
	)

	if err := row.Scan(&uid, &name, &email, &createdAt); err != nil {
		return domain.User{}, err
	}

	user := domain.NewUser(
		domain.UserID(uid),
		name,
		email,
		createdAt,
	)

	return user, nil
}