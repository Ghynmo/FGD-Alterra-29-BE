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

	_threadlikeUseCase "fgd-alterra-29/business/thread_likes"
	_threadlikeController "fgd-alterra-29/controllers/thread_likes"
	_threadlikeRepository "fgd-alterra-29/drivers/databases/thread_likes"

	_commentlikeUseCase "fgd-alterra-29/business/comment_likes"
	_commentlikeController "fgd-alterra-29/controllers/comment_likes"
	_commentlikeRepository "fgd-alterra-29/drivers/databases/comment_likes"

	_threadsaveUseCase "fgd-alterra-29/business/thread_saves"
	_threadsaveController "fgd-alterra-29/controllers/thread_saves"
	_threadsaveRepository "fgd-alterra-29/drivers/databases/thread_saves"

	_badgeRepository "fgd-alterra-29/drivers/databases/badges"
	_reputationRepository "fgd-alterra-29/drivers/databases/reputations"
	_threadfollowRepository "fgd-alterra-29/drivers/databases/thread_follows"

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
	db.AutoMigrate(&_threadlikeRepository.ThreadLikes{})
	db.AutoMigrate(&_threadfollowRepository.ThreadFollows{})
	db.AutoMigrate(&_commentlikeRepository.CommentLikes{})
	db.AutoMigrate(&_threadsaveRepository.ThreadSaves{})
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

	commentRepository := _commentRepository.NewMysqlCommentRepository(Conn)
	commentUseCase := _commentUseCase.NewCommentUseCase(commentRepository, timeoutContext)
	commentController := _commentController.NewCommentController(commentUseCase)

	threadRepository := _threadRepository.NewMysqlThreadRepository(Conn)
	threadUseCase := _threadUseCase.NewThreadUseCase(threadRepository, timeoutContext)
	threadController := _threadController.NewThreadController(threadUseCase)

	userbadgeRepository := _userbadgeRepository.NewMysqlUserBadgeRepository(Conn)
	userbadgeUseCase := _userbadgeUseCase.NewUserBadgeUseCase(userbadgeRepository, timeoutContext)
	userbadgeController := _userbadgeController.NewUserBadgeController(userbadgeUseCase)

	followRepository := _followRepository.NewMysqlFollowRepository(Conn)
	followUseCase := _followUseCase.NewFollowUseCase(followRepository, timeoutContext)
	followController := _followController.NewFollowController(followUseCase)

	categoryRepository := _categoryRepository.NewMysqlCategoryRepository(Conn)
	categoryUseCase := _categoryUseCase.NewCategoryUseCase(categoryRepository, timeoutContext)
	categoryController := _categoryController.NewCategoryController(categoryUseCase)

	threadlikeRepository := _threadlikeRepository.NewMysqlThreadLikeRepository(Conn)
	threadlikeUseCase := _threadlikeUseCase.NewThreadLikeUseCase(threadlikeRepository, timeoutContext)
	threadlikeController := _threadlikeController.NewThreadLikeController(threadlikeUseCase)

	commentlikeRepository := _commentlikeRepository.NewMysqlCommentLikeRepository(Conn)
	commentlikeUseCase := _commentlikeUseCase.NewCommentLikeUseCase(commentlikeRepository, timeoutContext)
	commentlikeController := _commentlikeController.NewCommentLikeController(commentlikeUseCase)

	threadsaveRepository := _threadsaveRepository.NewMysqlThreadSaveRepository(Conn)
	threadsaveUseCase := _threadsaveUseCase.NewThreadSaveUseCase(threadsaveRepository, timeoutContext)
	threadsaveController := _threadsaveController.NewThreadSaveController(threadsaveUseCase)

	userRepository := _userRepository.NewMysqlUserRepository(Conn)
	userUseCase := _userUseCase.NewUserUseCase(userRepository, timeoutContext)
	userController := _userController.NewUserController(userUseCase, threadUseCase, userbadgeUseCase, categoryUseCase)

	routesInit := routes.ControllerList{
		UserController:        *userController,
		UserBadgeController:   *userbadgeController,
		ThreadController:      *threadController,
		CommentController:     *commentController,
		FollowController:      *followController,
		CategoryController:    *categoryController,
		ThreadLikeController:  *threadlikeController,
		CommentLikeController: *commentlikeController,
		ThreadSaveController:  *threadsaveController,
	}

	routesInit.RouteRegister(*e)

	log.Fatal(e.Start((viper.GetString("server.address"))))
}
