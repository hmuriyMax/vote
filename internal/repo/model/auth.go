package model

import "time"

type UserAuth struct {
	ID            int64     `db:"id"`
	Name          string    `db:"username"`
	Password      string    `db:"password"`
	IsBlocked     bool      `db:"is_blocked"`
	LastLoginTime time.Time `db:"last_login"`
}

type Token struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	Token     string    `db:"token"`
	CreatedAt time.Time `db:"created_at"`
}
