// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories (name) VALUES ($1) RETURNING id, name
`

func (q *Queries) CreateCategory(ctx context.Context, name string) (Category, error) {
	row := q.db.QueryRow(ctx, createCategory, name)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, email, password_hash, is_verified, verification_token) VALUES ($1, $2, $3, $4, $5) RETURNING user_id, username, email, password_hash, created_at, updated_at, is_verified, verification_token
`

type CreateUserParams struct {
	Username          string      `json:"username"`
	Email             string      `json:"email"`
	PasswordHash      string      `json:"password_hash"`
	IsVerified        pgtype.Bool `json:"is_verified"`
	VerificationToken pgtype.Text `json:"verification_token"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.PasswordHash,
		arg.IsVerified,
		arg.VerificationToken,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsVerified,
		&i.VerificationToken,
	)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteCategory, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, userID pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, userID)
	return err
}

const getAllCategories = `-- name: GetAllCategories :many
SELECT id, name FROM categories
`

func (q *Queries) GetAllCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.Query(ctx, getAllCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT user_id, username, email, password_hash, created_at, updated_at, is_verified, verification_token FROM users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.Email,
			&i.PasswordHash,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IsVerified,
			&i.VerificationToken,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCategoryById = `-- name: GetCategoryById :one
SELECT id, name FROM categories WHERE id = $1
`

func (q *Queries) GetCategoryById(ctx context.Context, id int32) (Category, error) {
	row := q.db.QueryRow(ctx, getCategoryById, id)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT user_id, username, email, password_hash, created_at, updated_at, is_verified, verification_token FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsVerified,
		&i.VerificationToken,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT user_id, username, email, password_hash, created_at, updated_at, is_verified, verification_token FROM users WHERE user_id = $1
`

func (q *Queries) GetUserById(ctx context.Context, userID pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsVerified,
		&i.VerificationToken,
	)
	return i, err
}

const getUserByVerificationToken = `-- name: GetUserByVerificationToken :one
SELECT user_id, username, email, password_hash, created_at, updated_at, is_verified, verification_token FROM users WHERE verification_token = $1
`

func (q *Queries) GetUserByVerificationToken(ctx context.Context, verificationToken pgtype.Text) (User, error) {
	row := q.db.QueryRow(ctx, getUserByVerificationToken, verificationToken)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsVerified,
		&i.VerificationToken,
	)
	return i, err
}

const updateCategory = `-- name: UpdateCategory :one
UPDATE categories SET name = $1 WHERE id = $2 RETURNING id, name
`

type UpdateCategoryParams struct {
	Name string `json:"name"`
	ID   int32  `json:"id"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error) {
	row := q.db.QueryRow(ctx, updateCategory, arg.Name, arg.ID)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET username = $1, email = $2, password_hash = $3, is_verified = $4, verification_token = $5 WHERE user_id = $6 RETURNING user_id, username, email, password_hash, created_at, updated_at, is_verified, verification_token
`

type UpdateUserParams struct {
	Username          string      `json:"username"`
	Email             string      `json:"email"`
	PasswordHash      string      `json:"password_hash"`
	IsVerified        pgtype.Bool `json:"is_verified"`
	VerificationToken pgtype.Text `json:"verification_token"`
	UserID            pgtype.UUID `json:"user_id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Username,
		arg.Email,
		arg.PasswordHash,
		arg.IsVerified,
		arg.VerificationToken,
		arg.UserID,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsVerified,
		&i.VerificationToken,
	)
	return i, err
}

const verifyUser = `-- name: VerifyUser :exec
UPDATE users SET is_verified = true, verification_token = null , updated_at = now() WHERE verification_token = $1  RETURNING user_id, username, email, password_hash, created_at, updated_at, is_verified, verification_token
`

func (q *Queries) VerifyUser(ctx context.Context, verificationToken pgtype.Text) error {
	_, err := q.db.Exec(ctx, verifyUser, verificationToken)
	return err
}
