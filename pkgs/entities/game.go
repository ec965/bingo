package entities

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/jackc/pgx/v4"
)

// when a user creates a game:
// 1. create the game
// 2. subsequent api calls will add cards to the game

type Card struct {
	CardId int    `json:"card_id"`
	Text   string `json:"text"`
}

type Game struct {
	GameId    int    `json:"game_id"`
	Cards     []Card `json:"cards"`
	Dimension int    `json:"dimension"`
}

// create cards for a game
func CreateCards(
	ctx context.Context,
	dbConn *pgx.Conn,
	gameId string,
	cardStrs []string,
) ([]Card, error) {
	// insert cards
	var cardRows []interface{}
	for _, cardStr := range cardStrs {
		cardRows = append(cardRows, goqu.Record{"game_id": gameId, "text": cardStr})
	}
	sql, args, err := dialect.
		Insert("cards").
		Rows(cardRows...).
		Returning("card_id", "text").
		ToSQL()
	if err != nil {
		panic(err)
	}
	rows, err := dbConn.Query(ctx, sql, args...)
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
	return cardArr, nil
}

// create a new blank game with dimension
func CreateGame(
	ctx context.Context,
	dbConn *pgx.Conn,
	dimension int,
) (*Game, error) {
	// insert game
	sql, args, err := dialect.Insert("games").
		Rows(goqu.Record{"dimension": dimension}).
		Returning("game_id", "dimension").
		ToSQL()
	if err != nil {
		panic(err)
	}
	// we want to send back an empty array
	cardArr := []Card{}
	game := Game{Cards: cardArr}
	err = dbConn.QueryRow(ctx, sql, args...).Scan(&game.GameId, &game.Dimension)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

// find a game and all its cards
func FindGame(
	ctx context.Context,
	dbConn *pgx.Conn,
	gameId int,
) (*Game, error) {
	sql, args := getFindGameSql(gameId)
	rows, err := dbConn.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	game := Game{}
	cardArr := []Card{}
	for rows.Next() {
		var cardId *int
		var cardText *string
		err := rows.Scan(&game.GameId, &game.Dimension, &cardId, &cardText)
		card := Card{}
		if cardId != nil && cardText != nil {
			card.Text = *cardText
			card.CardId = *cardId
			cardArr = append(cardArr)
		}
		if err != nil {
			return nil, err
		}
	}

	game.Cards = cardArr
	return &game, nil
}

// generate find game sql
func getFindGameSql(gameId int) (string, []interface{}) {
	sql, args, err := dialect.
		From("games").
		Select("games.game_id", "games.dimension", "cards.card_id", "cards.text").
		Where(goqu.Ex{"games.game_id": gameId}).
		LeftJoin(
			goqu.T("cards"),
			goqu.On(goqu.Ex{
				"games.game_id": goqu.I("cards.game_id"),
			}),
		).
		ToSQL()
	if err != nil {
		panic(err)
	}
	return sql, args
}
