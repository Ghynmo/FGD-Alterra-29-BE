package main

import (
	"fgd-alterra-29/app/routes"
	_mysqlDriver "fgd-alterra-29/drivers/mysqlconf"
	"log"
	"time"

	_userUseCase "fgd-alterra-29/business/users"
	_userController "fgd-alterra-29/controllers/users"
	_userRepository "fgd-alterra-29/drivers/databases/users"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/configs/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func DbMigrate(db *gorm.DB) {
	db.AutoMigrate(&_userRepository.Users{})
}

func main() {

	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	Conn := configDB.InitialDB()
	DbMigrate(Conn)

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := _userRepository.NewMysqlUserRepository(Conn)
	userUseCase := _userUseCase.NewUserUseCase(userRepository, timeoutContext)
	userController := _userController.NewUserController(userUseCase)

	routesInit := routes.ControllerList{
		UserController: *userController,
	}

	routesInit.RouteRegister(*e)

	log.Fatal(e.Start((viper.GetString("server.address"))))
}
