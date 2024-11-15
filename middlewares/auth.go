package middlewares

import (
	"go-carbon-tracker/constants"
	"go-carbon-tracker/handlers/base"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWTCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type JWTConfig struct {
	SecretKey       string
	ExpiresDuration int
}

func (jwtConfig *JWTConfig) Init() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTCustomClaims)
		},
		SigningKey: []byte(jwtConfig.SecretKey),
		ErrorHandler: func(c echo.Context, err error) error {
			if err != nil {
				return base.ErrorResponse(c, err)
			}

			return nil
		},
	}
}

func (jwtConfig *JWTConfig) GenerateToken(userID string) (string, error) {
	expire := jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConfig.ExpiresDuration))))

	claims := &JWTCustomClaims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: expire,
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte(jwtConfig.SecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUser(c echo.Context) (*JWTCustomClaims, error) {
	user := c.Get("user").(*jwt.Token)

	if user == nil {
		return nil, constants.Err_INVALID_TOKEN
	}

	claims := user.Claims.(*JWTCustomClaims)

	return claims, nil
}
