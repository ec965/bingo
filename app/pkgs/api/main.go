package api

import (
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/ec965/bingo/pkgs/entities"
	"github.com/ec965/bingo/pkgs/middleware"
	"github.com/ec965/bingo/pkgs/token"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v4"
)

var (
	validate     *validator.Validate
	dbConn       *pgx.Conn
	tokenManager *token.TokenManager
	secret       []byte
	authRoute    func(func(http.ResponseWriter, *http.Request, *entities.User)) http.HandlerFunc
)

func init() {
	tokenManager = &token.TokenManager{
		Secret: secret,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // 1 week
		},
	}
	authRoute = middleware.AuthUser(tokenManager)

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
func Start(conn *pgx.Conn, tokenSecret []byte) {
	dbConn = conn
	secret = tokenSecret
}
