package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/ec965/bingo/pkgs/entities"
	"github.com/golang-jwt/jwt"
)

type tokenClaims struct {
	*jwt.StandardClaims
	user *entities.User
}

type TokenManager struct {
	Secret         []byte
	StandardClaims *jwt.StandardClaims
}

// createjwt from user entity
func (tm *TokenManager) CreateToken(user *entities.User) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		&tokenClaims{
			tm.StandardClaims,
			user,
		},
	)
	// signed string expects a byte array
	tStr, err := t.SignedString(tm.Secret)
	if err != nil {
		panic(err)
	}
	return tStr
}

// validate jwt
func (tm *TokenManager) ValidateToken(tokenString string) (*entities.User, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&tokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(
					"unexpected signing method: %v", token.Header["alg"],
				)
			}
			return tm.Secret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		isExpired := claims.VerifyExpiresAt(time.Now().Unix(), true)
		if isExpired {
			return nil, errors.New("token is expired")
		}
		return claims.user, nil
	}
	return nil, errors.New("invalid token")
}
