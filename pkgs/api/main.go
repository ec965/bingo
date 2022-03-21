package api

import (
	"reflect"
	"strings"
	"time"

	"github.com/ec965/bingo/pkgs/token"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v4"
)

var (
	validate     *validator.Validate
	dbConn       *pgx.Conn
	tokenManager *token.TokenManager
)

func init() {
	tokenManager = &token.TokenManager{
		Secret: []byte("very_secret"), // FIXME: uhhh not very secret lol
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // 1 week
		},
	}

	validate = validator.New()
	// get the json encoded name on error
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

}

// start function to connect to the database
func DbConnect(conn *pgx.Conn) {
	dbConn = conn
}
