package entities

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/jackc/pgx/v4"
)

type Card struct {
	CardId string `json:"card_id"`
	Text   string `json:"text"`
}

type Game struct {
	GameId    string `json:"game_id"`
	Cards     []Card `json:"cards"`
	Dimension int    `json:"dimension"`
}

// TODO: split this up into 2 functions
// when a user creates a game:
// 1. create the game
// 2. subsequent api calls will add cards to the game
// WARN: you already deleted the migrations to remove deferment for FKs
// create a game with associated cards
func CreateGame(
	ctx context.Context,
	dbConn *pgx.Conn,
	cardStrs []string,
	dimension int,
) (*Game, error) {
	tx, err := dbConn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// insert game
	gameSql, gameArgs, err := dialect.Insert("games").
		Rows(goqu.Record{"dimension": dimension}).
		Returning("game_id", "dimension").
		ToSQL()
	if err != nil {
		panic(err)
	}
	game := Game{}
	err = tx.QueryRow(ctx, gameSql, gameArgs...).Scan(&game.GameId, &game.Dimension)
	if err != nil {
		return nil, err
	}

	// insert cards
	var cardRows []interface{}
	for _, cardStr := range cardStrs {
		cardRows = append(cardRows, goqu.Record{"game_id": game.GameId, "text": cardStr})
	}
	cardSql, cardArgs, err := dialect.
		Insert("cards").
		Rows(cardRows...).
		Returning("card_id", "text").
		ToSQL()
	if err != nil {
		panic(err)
	}
	rows, err := tx.Query(ctx, cardSql, cardArgs...)
	if err != nil {
		return nil, err
	}
	var cardArr []Card
	for rows.Next() {
		card := Card{}
		err := rows.Scan(&card.CardId, &card.Text)
		if err != nil {
			return nil, err
		}
		cardArr = append(cardArr)
	}

	game.Cards = cardArr

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err

	}
	return &game, nil
}
