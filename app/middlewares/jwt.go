package middlewares

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type JWTClaims struct {
	UserID int
	Admin  bool
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

func (JwtConf *ConfigJWT) GenerateToken(id int, admin bool) (string, error) {
	var IsAdmin = false
	if admin {
		IsAdmin = true
	} else {
		IsAdmin = false
	}

	claims := JWTClaims{
		id,
		IsAdmin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(viper.GetInt64(`jwt.expired`)))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(viper.GetString(`jwt.secret`)))

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecret := []byte(viper.GetString(`jwt.secret`))
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return claims, true
	} else {
		fmt.Println("Invalid JWT Token")
		return nil, false
	}
}
