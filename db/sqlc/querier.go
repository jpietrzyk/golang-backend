// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateGame(ctx context.Context, arg CreateGameParams) (Game, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteGame(ctx context.Context, id int64) error
	GetGame(ctx context.Context, id int64) (Game, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetUser(ctx context.Context, email string) (User, error)
	ListGames(ctx context.Context, arg ListGamesParams) ([]Game, error)
	UpdateGame(ctx context.Context, arg UpdateGameParams) (Game, error)
}

var _ Querier = (*Queries)(nil)
