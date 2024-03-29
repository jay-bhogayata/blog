// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Article struct {
	ArticleID   pgtype.UUID      `json:"article_id"`
	Title       string           `json:"title"`
	Content     string           `json:"content"`
	UserID      pgtype.UUID      `json:"user_id"`
	CategoryID  pgtype.Int4      `json:"category_id"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	IsPublished pgtype.Bool      `json:"is_published"`
}

type Category struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type User struct {
	UserID            pgtype.UUID      `json:"user_id"`
	Username          string           `json:"username"`
	Email             string           `json:"email"`
	PasswordHash      string           `json:"password_hash"`
	CreatedAt         pgtype.Timestamp `json:"created_at"`
	UpdatedAt         pgtype.Timestamp `json:"updated_at"`
	IsVerified        pgtype.Bool      `json:"is_verified"`
	VerificationToken pgtype.Text      `json:"verification_token"`
}
