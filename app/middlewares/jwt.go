package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type JWTClaims struct {
	UserID int
	jwt.RegisteredClaims
}

type ConfigJWT struct {
	Secret    string
	ExpiresAt int64
}

func (JwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTClaims{},
		SigningKey: []byte(JwtConf.Secret),
	}
}

func GenerateToken(id int) (string, error) {
	claims := JWTClaims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(viper.GetInt64(`jwt.expired`)*3600, 0)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString([]byte(viper.GetString(`jwt.secret`)))

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
