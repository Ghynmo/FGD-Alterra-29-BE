package users

import (
	"fgd-alterra-29/business/badges"
	"fgd-alterra-29/business/categories"
	commentreport "fgd-alterra-29/business/comment_report"
	"fgd-alterra-29/business/comments"
	"fgd-alterra-29/business/threads"
	userbadges "fgd-alterra-29/business/user_badges"
	"fgd-alterra-29/business/users"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/users/request"
	"fgd-alterra-29/controllers/users/responses"
	editprofile "fgd-alterra-29/controllers/users/responses/edit_profile"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase          users.UseCase
	ThreadUseCase        threads.UseCase
	UserBadgeUseCase     userbadges.UseCase
	CategoryUseCase      categories.UseCase
	CommentReportUseCase commentreport.UseCase
	CommentUseCase       comments.UseCase
	BadgeUseCase         badges.UseCase
}

func NewUserController(userUC users.UseCase, threadUC threads.UseCase, userbadgeUC userbadges.UseCase, categoryUC categories.UseCase, commentreportUC commentreport.UseCase, commentUC comments.UseCase, badgeUC badges.UseCase) *UserController {
	return &UserController{
		UserUseCase:          userUC,
		ThreadUseCase:        threadUC,
		UserBadgeUseCase:     userbadgeUC,
		CategoryUseCase:      categoryUC,
		CommentReportUseCase: commentreportUC,
		CommentUseCase:       commentUC,
		BadgeUseCase:         badgeUC,
	}
}

func (handler UserController) RegisterController(c echo.Context) error {
	var NewRegister request.Register
	c.Bind(&NewRegister)

	domain := NewRegister.FromRegister()

	ctx := c.Request().Context()

	user, err := handler.UserUseCase.RegisterController(ctx, domain)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToLoginResponse(user))
}

func (handler UserController) LoginController(c echo.Context) error {
	var NewLogin request.Login
	c.Bind(&NewLogin)

	domain := NewLogin.FromLogin()

	ctx := c.Request().Context()

	user, err := handler.UserUseCase.LoginController(ctx, domain)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToLoginResponse(user))
}

func (handler UserController) GetUsersController(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := handler.UserUseCase.GetUsersController(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromListDomain(user))
}

func (handler UserController) GetUsersByName(c echo.Context) error {
	name := c.Param("name")

	ctx := c.Request().Context()

	user, err := handler.UserUseCase.GetUsersByNameController(ctx, name)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromListDomain(user))
}

func (handler UserController) GetProfileController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ctx := c.Request().Context()

	user, err := handler.UserUseCase.GetProfileController(ctx, id)
	thread, err1 := handler.ThreadUseCase.GetProfileThreads(ctx, id)
	badges, err2 := handler.BadgeUseCase.GetBadgesByUserController(ctx, id)
	catthreads, err3 := handler.CategoryUseCase.GetUserActiveInCategory(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if err1 != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if err2 != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if err3 != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToProfile(user, badges, catthreads, thread))
}

func (handler UserController) GetDashboardController(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := handler.UserUseCase.GetUsersController(ctx)
	commentreport, err1 := handler.CommentReportUseCase.GetCommentReportStat(ctx)
	threadqty, err2 := handler.ThreadUseCase.GetThreadQuantity(ctx)
	userqty, err3 := handler.UserUseCase.GetUsersQuantity(ctx)
	postqty, err4 := handler.CommentUseCase.GetPostQuantityController(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if err1 != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err1)
	}
	if err2 != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err2)
	}
	if err3 != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err3)
	}
	if err4 != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err3)
	}
	return controllers.NewSuccessResponse(c, responses.ToDashboard(user, userqty, threadqty, postqty, commentreport))
}

func (handler UserController) GetAdminSettingController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	user, err := handler.UserUseCase.GetProfileSetting(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, editprofile.ToAdminEdit(user))
}

func (handler UserController) GetUserSettingController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	user, err := handler.UserUseCase.GetProfileSetting(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, editprofile.ToUserEdit(user))
}

func (handler UserController) UpdateAdminProfile(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	Updateadmin := request.UpdateProfile{}
	c.Bind(&Updateadmin)

	Domain := Updateadmin.ToDomain()

	ctx := c.Request().Context()

	user, err := handler.UserUseCase.UpdateProfile(ctx, Domain, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, editprofile.ToAdminEdit(user))
}

func (handler UserController) UpdateUserProfile(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	Updateuser := request.UpdateProfile{}
	c.Bind(&Updateuser)

	Domain := Updateuser.ToDomain()

	ctx := c.Request().Context()

	user, err := handler.UserUseCase.UpdateProfile(ctx, Domain, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, editprofile.ToUserEdit(user))
}

func (handler UserController) BannedUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	user, err := handler.UserUseCase.BannedUser(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.BannedSuccessResponse(c, user)
}

func (handler UserController) UnbannedUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	user, err := handler.UserUseCase.UnbannedUser(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.UnbannedSuccessResponse(c, user)
}
