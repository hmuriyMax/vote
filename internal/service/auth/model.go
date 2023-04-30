package auth

import "errors"

var (
	ErrNoUser              = errors.New("no user found")
	ErrWrongPassword       = errors.New("wrong password")
	ErrAccountAlreadyExist = errors.New("account already exists")
)
