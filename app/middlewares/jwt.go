package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"fgd-alterra-29/app/configs"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	config, err := configs.LoadConfig("./") //My .env file stored
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	NewExpiredJWT, _ := strconv.Atoi(config.JWTExpired)

	claims := JWTClaims{
		id,
		IsAdmin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(NewExpiredJWT))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(config.JWTSecret))

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {

	config, err := configs.LoadConfig("./") //My .env file stored
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	hmacSecret := []byte(config.JWTSecret)
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

func ExtractID(c echo.Context) int {
	reqToken := c.Request().Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// convert map to json
	claims, _ := ExtractClaims(reqToken)
	jsonString, _ := json.Marshal(claims)

	// convert json to struct
	s := JWTClaims{}
	json.Unmarshal(jsonString, &s)

	return s.UserID
}

func ExtractAdmin(c echo.Context) (bool, error) {
	reqToken := c.Request().Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// convert map to json
	claims, _ := ExtractClaims(reqToken)
	jsonString, _ := json.Marshal(claims)

	// convert json to struct
	s := JWTClaims{}
	json.Unmarshal(jsonString, &s)

	return s.Admin, nil
}
