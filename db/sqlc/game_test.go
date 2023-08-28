package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomGame(t *testing.T) Game {
	user := createRandomUser(t)

	arg := CreateGameParams{
		OwnerID: user.ID,
	}

	game, err := testStore.CreateGame(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, game)

	require.Equal(t, arg.OwnerID, game.OwnerID)

	require.NotZero(t, game.ID)
	require.NotZero(t, game.CreatedAt)
	require.NotZero(t, game.StartsAt)
	require.NotZero(t, game.EndsAt)

	return game
}

func TestCreategame(t *testing.T) {
	createRandomGame(t)
}

func TestGetGame(t *testing.T) {
	game1 := createRandomGame(t)
	game2, err := testStore.GetGame(context.Background(), game1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, game2)

	require.Equal(t, game1.ID, game2.ID)
	require.Equal(t, game1.OwnerID, game2.OwnerID)
	require.WithinDuration(t, game1.CreatedAt, game2.CreatedAt, time.Second)
}

func TestUpdateGame(t *testing.T) {
	game1 := createRandomGame(t)

	arg := UpdateGameParams{
		ID: game1.ID,
	}

	game2, err := testStore.UpdateGame(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, game2)

	require.Equal(t, game1.ID, game2.ID)
	require.Equal(t, game1.OwnerID, game2.OwnerID)
	require.WithinDuration(t, game1.CreatedAt, game2.CreatedAt, time.Second)
}

func TestDeleteGame(t *testing.T) {
	game1 := createRandomGame(t)
	err := testStore.DeleteGame(context.Background(), game1.ID)
	require.NoError(t, err)

	game2, err := testStore.GetGame(context.Background(), game1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, game2)
}

func TestListGames(t *testing.T) {
	var lastGame Game
	for i := 0; i < 10; i++ {
		lastGame = createRandomGame(t)
	}

	arg := ListGamesParams{
		OwnerID: lastGame.OwnerID,
		Limit:   5,
		Offset:  0,
	}

	games, err := testStore.ListGames(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, games)

	for _, game := range games {
		require.NotEmpty(t, game)
		require.Equal(t, lastGame.OwnerID, game.OwnerID)
	}
}
