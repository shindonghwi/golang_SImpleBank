// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccounts(ctx context.Context, arg CreateAccountsParams) (Account, error)
}

var _ Querier = (*Queries)(nil)
