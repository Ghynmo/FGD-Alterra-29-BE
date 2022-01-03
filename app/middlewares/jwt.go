package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type JWTClaims struct {
	UserID int
	jwt.RegisteredClaims
}

//Initialize Viper just for get data from JSON file
func init() {
	viper.SetConfigFile(`app/configs/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
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
