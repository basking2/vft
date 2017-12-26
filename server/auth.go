package Server

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/madurosecurity/vft"
	"time"
)

type customClaims struct {
	ClientId string `json:"client_id"`
	jwt.StandardClaims
}

func (s *Server) authenticate(m *vft.Message) (string, error) {
	if m.Secret == s.secret {
		return s.newJWT(m)
	} else {
		s.log.Info("Bad secret")
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
			ExpiresAt: time.Now().Unix() + 86400, // 24 hours
			IssuedAt:  time.Now().Unix(),
			Issuer:    "vft",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

func (s *Server) validateJWT(tokenString string) (isValid bool, shouldRenew bool) {
	isValid = false
	shouldRenew = true

	signingKey := []byte(s.db.GetSigningKey())
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	// I definitely stole this from https://github.com/dgrijalva/jwt-go/blob/master/example_test.go
	if token.Valid {
		isValid = true
		claims, _ := token.Claims.(*customClaims)
		if time.Now().Unix()+300 > claims.StandardClaims.ExpiresAt {
			shouldRenew = true
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			s.log.Error("Malformed token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			s.log.Error("Token is either not valid yet or has expired")
		} else {
			s.log.Error("Unable to handle token " + err.Error())
		}
	} else {
		s.log.Error("Unable to handle token " + err.Error())
	}
	return
}
