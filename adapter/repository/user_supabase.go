package repository

import (
	"context"
	"database/sql"
	"go_cleanArchitecture_study/domain"
)

type UserSupabase struct {
	db *sql.DB
}

func NewUserSupabase(db *sql.DB) *UserSupabase {
	return &UserSupabase{
		db: db,
	}
}

func (r *UserSupabase) Create(user domain.User) (domain.User, error) {

	query := `INSERT INTO users (id, name, email, created_at) VALUES ($1, $2, $3, $4)`

	_, err := r.db.ExecContext(
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

func (r *UserSupabase) FindByID(id domain.UserID) (domain.User, error) {
	row := r.db.QueryRow(
		`SELECT id, name, email, created_at FROM users WHERE id = $1`,
		id,
	)

	var (
		uid 		string
		name		string
		email		string
		createdAt	sql.NullTime
	)

	if err := row.Scan(&uid, &name, &email, &createdAt); err != nil {
		return domain.User{}, err
	}

	user := domain.NewUser(
		domain.UserID(uid),
		name,
		email,
		createdAt.Time,
	)

	return user, nil


}