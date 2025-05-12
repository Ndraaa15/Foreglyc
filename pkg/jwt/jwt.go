package jwt

import (
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/user/entity"
	"github.com/Ndraaa15/foreglyc-server/pkg/errx"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

var (
	ErrInvalidTokenExpired = errx.Unauthorized("invalid token expired")
	ErrFailedClaimJWT      = errx.Unauthorized("failed claim jwt")
	ErrInvalidSignature    = errx.Unauthorized("invalid signature")
	ErrSignJwt             = errx.Unauthorized("failed to sign jwt")
	ErrMalformedToken      = errx.BadRequest("malformed token")
)

const (
	TokenTypeBearer string = "Bearer"
)

func EncodeToken(user *entity.User) (string, error) {
	claims := &JWTClaims{
		ID: user.Id.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.expiration"))),
			Issuer:    viper.GetString("jwt.issuer"),
			Subject:   user.Id.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(viper.GetString("JWT_SECRET_KEY")))
	if err != nil {
		return "", ErrSignJwt
	}
	return signedToken, nil
}

func DecodeToken(token string) (*JWTClaims, error) {
	decoded, err := jwt.ParseWithClaims(token, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		token.Method = jwt.SigningMethodHS256
		return []byte(viper.GetString("jwt.secretkey")), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return &JWTClaims{}, ErrMalformedToken
			}
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return &JWTClaims{}, ErrInvalidTokenExpired
			}
		}
		return &JWTClaims{}, ErrInvalidSignature
	}

	if !decoded.Valid {
		return &JWTClaims{}, ErrInvalidTokenExpired
	}

	claims, ok := decoded.Claims.(*JWTClaims)
	if !ok {
		return &JWTClaims{}, ErrFailedClaimJWT
	}

	return claims, nil
}
