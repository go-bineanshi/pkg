package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTConfig struct {
	Secret   string   `json:"secret"`
	Expires  int      `json:"expires,default=3600"`
	Subject  string   `json:"subject,omitempty"`
	Issuer   string   `json:"issuer,omitempty"`
	Audience []string `json:"audience,omitempty"`
}

type claims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

func (jc *JWTConfig) GenerateToken(userID string) (string, error) {
	claims := claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jc.Expires) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    jc.Issuer,
			Subject:   jc.Subject,
			ID:        userID,
			Audience:  jc.Audience,
		},
		UserID: userID,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(jc.Secret))
}

func (jc *JWTConfig) GenerateRefreshToken(userID string) (string, error) {
	claims := claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jc.Expires*2) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    jc.Issuer,
			Subject:   jc.Subject,
			ID:        userID,
			Audience:  jc.Audience,
		},
		UserID: userID,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(jc.Secret))
}

func (jc *JWTConfig) ParseToken(tokenString string) (userID string, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jc.Secret), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*claims); ok && token.Valid {
		return claims.UserID, nil
	}
	return "", err
}
