// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, id int64) (User, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (User, error)
	GetUserEmail(ctx context.Context, email string) (User, error)
	GetUsersByFirstName(ctx context.Context, firstName string) ([]User, error)
	GetUsersByLastName(ctx context.Context, lastName string) ([]User, error)
	UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (User, error)
	UpdateUserFirstName(ctx context.Context, arg UpdateUserFirstNameParams) (User, error)
	UpdateUserLastName(ctx context.Context, arg UpdateUserLastNameParams) (User, error)
	UpdateUserPhoneNumber(ctx context.Context, arg UpdateUserPhoneNumberParams) (User, error)
	UpdateUsersPassword(ctx context.Context, arg UpdateUsersPasswordParams) (User, error)
}

var _ Querier = (*Queries)(nil)
