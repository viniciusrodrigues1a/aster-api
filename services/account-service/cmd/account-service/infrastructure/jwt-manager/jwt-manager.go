package jwtmanager

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTManager struct {
	privateKey []byte
	publicKey  []byte
}

func NewJWTManager(privateKey []byte, publicKey []byte) *JWTManager {
	return &JWTManager{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func (j *JWTManager) Sign(email string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 4).Unix()

	key, err := jwt.ParseRSAPrivateKeyFromPEM(j.privateKey)
	if err != nil {
		return "", fmt.Errorf("sign: parse key: %w", err)
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("sign: sign token: %w", err)
	}

	return token, nil
}

func (j *JWTManager) Verify(token string) (interface{}, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(j.publicKey)
	if err != nil {
		return "", fmt.Errorf("verify: parse key %w", err)
	}

	parsedToken, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("verify: unexpected method: %s", jwtToken.Header["alg"])
		}

		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("verify: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("verify: invalid")
	}

	return claims["email"], nil
}
