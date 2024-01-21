package api

import (
	"Test-Task2/pkg/api/jwt"
	"fmt"
	j "github.com/dgrijalva/jwt-go/v4"
)

func GetMapClaims(token string, jwt *jwt.JWT) (j.MapClaims, error) {
	if token == "" {
		return nil, fmt.Errorf("authorization header required")
	}

	mapClaims, err := jwt.Parse(token)
	if err != nil {
		return nil, err
	}

	return mapClaims, nil
}
