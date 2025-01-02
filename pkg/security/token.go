package security

import (
	"jwtgo/pkg/security/schema"
	"time"

	"github.com/dgrijalva/jwt-go"
	customErr "jwtgo/pkg/security/error"
)

type TokenManager struct {
	secretKey       string
	accessLifetime  int
	refreshLifetime int
}

func NewTokenManager(secretKey string, accessLifetime, refreshLifetime int) *TokenManager {
	return &TokenManager{
		secretKey:       secretKey,
		accessLifetime:  accessLifetime,
		refreshLifetime: refreshLifetime,
	}
}

func (tm *TokenManager) GenerateTokens(id string) (string, string, error) {
	accessClaims := &schema.Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Minute * time.Duration(tm.accessLifetime)).Unix(),
		},
	}

	refreshClaims := &schema.Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Minute * time.Duration(tm.refreshLifetime)).Unix(),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(tm.secretKey))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(tm.secretKey))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (tm *TokenManager) ValidateToken(signedToken string) (*schema.Claims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&schema.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(tm.secretKey), nil
		},
	)

	if err != nil {
		return nil, customErr.NewInvalidTokenError("Token is invalid")
	}

	claims, ok := token.Claims.(*schema.Claims)
	if !ok {
		return nil, customErr.NewInvalidTokenError("Token is invalid")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, customErr.NewExpiredTokenError("Token is expired")
	}

	return claims, nil
}
