package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {

	var insertedId int

	query := `INSERT INTO users(username, password, email) VALUES($1,$2,$3) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&insertedId)
	if err != nil {
		return &User{}, err
	}

	user.ID = int64(insertedId)
	return user, err
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := &User{}

	query := "SELECT * FROM users WHERE email=$1"
	err := r.db.QueryRowContext(ctx, query).Scan(&u.ID, &u.Email, &u.Username, &u.Password)
	if err != nil {
		return &User{}, nil
	}

	return u, err
}