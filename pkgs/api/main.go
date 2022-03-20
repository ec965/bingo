package api

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4"
)

var (
	validate *validator.Validate
	dbConn   *pgx.Conn
)

func init() {
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

// example handler
func ApiHandler(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	j, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
