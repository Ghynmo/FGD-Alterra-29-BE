package main

import (
	"fgd-alterra-29/app/routes"
	_mysqlDriver "fgd-alterra-29/drivers/mysql"
	"log"
	"time"

	_userUseCase "fgd-alterra-29/business/users"
	_userController "fgd-alterra-29/controllers/users"
	_userRepository "fgd-alterra-29/drivers/databases/users"

	_threadUseCase "fgd-alterra-29/business/threads"
	_threadController "fgd-alterra-29/controllers/threads"
	_threadRepository "fgd-alterra-29/drivers/databases/threads"

	_commentUseCase "fgd-alterra-29/business/comments"
	_commentController "fgd-alterra-29/controllers/comments"
	_commentRepository "fgd-alterra-29/drivers/databases/comments"

	_followUseCase "fgd-alterra-29/business/follows"
	_followController "fgd-alterra-29/controllers/follows"
	_followRepository "fgd-alterra-29/drivers/databases/follows"

	_userbadgeUseCase "fgd-alterra-29/business/user_badges"
	_userbadgeController "fgd-alterra-29/controllers/user_badges"
	_userbadgeRepository "fgd-alterra-29/drivers/databases/user_badges"

	_categoryUseCase "fgd-alterra-29/business/categories"
	_categoryController "fgd-alterra-29/controllers/categories"
	_categoryRepository "fgd-alterra-29/drivers/databases/categories"

	_userpointUseCase "fgd-alterra-29/business/user_points"
	_userpointController "fgd-alterra-29/controllers/user_points"
	_userpointRepository "fgd-alterra-29/drivers/databases/user_points"

	_badgeUseCase "fgd-alterra-29/business/badges"
	_badgeController "fgd-alterra-29/controllers/badges"
	_badgeRepository "fgd-alterra-29/drivers/databases/badges"

	_reputationRepository "fgd-alterra-29/drivers/databases/reputations"

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
	db.AutoMigrate(&_reputationRepository.Reputations{})
	db.AutoMigrate(&_badgeRepository.Badges{})
	db.AutoMigrate(&_categoryRepository.Categories{})
	db.AutoMigrate(&_followRepository.Follows{})
	db.AutoMigrate(&_threadRepository.Threads{})
	db.AutoMigrate(&_commentRepository.Comments{})
	db.AutoMigrate(&_userbadgeRepository.UserBadges{})
	db.AutoMigrate(&_userpointRepository.UserPoints{})
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

	userpointRepository := _userpointRepository.NewMysqlUserPointRepository(Conn)
	userpointUseCase := _userpointUseCase.NewUserPointUseCase(userpointRepository, timeoutContext)
	userpointController := _userpointController.NewUserPointController(userpointUseCase)

	threadRepository := _threadRepository.NewMysqlThreadRepository(Conn)
	threadUseCase := _threadUseCase.NewThreadUseCase(threadRepository, timeoutContext, userpointRepository)
	threadController := _threadController.NewThreadController(threadUseCase)

	userbadgeRepository := _userbadgeRepository.NewMysqlUserBadgeRepository(Conn)
	userbadgeUseCase := _userbadgeUseCase.NewUserBadgeUseCase(userbadgeRepository, timeoutContext)
	userbadgeController := _userbadgeController.NewUserBadgeController(userbadgeUseCase)

	commentRepository := _commentRepository.NewMysqlCommentRepository(Conn)
	commentUseCase := _commentUseCase.NewCommentUseCase(commentRepository, timeoutContext)
	commentController := _commentController.NewCommentController(commentUseCase)

	followRepository := _followRepository.NewMysqlFollowRepository(Conn)
	followUseCase := _followUseCase.NewFollowUseCase(followRepository, timeoutContext)
	followController := _followController.NewFollowController(followUseCase)

	categoryRepository := _categoryRepository.NewMysqlCategoryRepository(Conn)
	categoryUseCase := _categoryUseCase.NewCategoryUseCase(categoryRepository, timeoutContext)
	categoryController := _categoryController.NewCategoryController(categoryUseCase)

	badgeRepository := _badgeRepository.NewMysqlBadgeRepository(Conn)
	badgeUseCase := _badgeUseCase.NewBadgeUseCase(badgeRepository, timeoutContext)
	badgeController := _badgeController.NewBadgeController(badgeUseCase)

	userRepository := _userRepository.NewMysqlUserRepository(Conn)
	userUseCase := _userUseCase.NewUserUseCase(userRepository, timeoutContext)
	userController := _userController.NewUserController(userUseCase, threadUseCase, badgeUseCase, categoryUseCase)

	routesInit := routes.ControllerList{
		UserController:      *userController,
		UserBadgeController: *userbadgeController,
		ThreadController:    *threadController,
		CommentController:   *commentController,
		FollowController:    *followController,
		CategoryController:  *categoryController,
		UserPointController: *userpointController,
		BadgeController:     *badgeController,
	}

	routesInit.RouteRegister(*e)

	log.Fatal(e.Start((viper.GetString("server.address"))))
}
