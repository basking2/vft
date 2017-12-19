package Server

import (
	"fmt"
	"github.com/bbriggs/vft"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type customClaims struct {
	ClientId string `json:"client_id"`
	jwt.StandardClaims
}

func (s *Server) authenticate(m *vft.Message) (string, error) {
	if m.Secret == "SharedSecret" {
		return s.newJWT(m)
	} else {
		err := fmt.Errorf("Authentication failed!")
		return "", err
	}
}

// Create the Claims
func (s *Server) newJWT(m *vft.Message) (string, error) {
	signingKey := []byte(s.db.GetSigningKey())
	claims := customClaims{
		m.ClientId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "vft",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

func (s *Server) validateJWT(tokenString string) bool {
	signingKey := []byte(s.db.GetSigningKey())
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	// I definitely stole this from https://github.com/dgrijalva/jwt-go/blob/master/example_test.go
	if token.Valid {
		return true
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			s.log.Error("Malformed token")
			return false
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			s.log.Error("Token is either not valid yet or has expired")
			return false
		} else {
			s.log.Error("Unable to handle token " + err.Error())
			return false
		}
	} else {
		s.log.Error("Unable to handle token " + err.Error())
		return false
	}
}