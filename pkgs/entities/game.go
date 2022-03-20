package entities

/* import (
	"context"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/jackc/pgx/v4"
) */

type Card struct {
	CardId string `json:"card_id"`
	Text   string `json:"text"`
}

type Game struct {
	GameId    string `json:"game_id"`
	Cards     []Card `json:"cards"`
	Dimension int    `json:"dimension"`
}

/* func CreateGame(
	ctx context.Context,
	dbConn *pgx.Conn,
	cardStrs []string,
	dimension int
) (*Game, error) {
} */
