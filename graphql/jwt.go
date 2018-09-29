package graphql

import (
	"github.com/Tinee/go-graphql-chat/domain"
	jwt "github.com/dgrijalva/jwt-go"
)

func (r *Resolver) claimJWT(u domain.User) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["username"] = u.Username

	t, _ := token.SignedString([]byte(r.secret))

	return t
}

func (r *Resolver) validateAndExtractId(token string) (id string, err error) {
	t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte(r.secret), nil
	})
	if err != nil || !t.Valid {
		return "", ErrInvalidToken
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return "", ErrInvalidToken
	}

	id, ok = claims["id"].(string)
	if !ok {
		return "", ErrInvalidToken
	}

	return id, nil
}
