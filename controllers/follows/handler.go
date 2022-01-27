package follows

import (
	"fgd-alterra-29/app/middlewares"
	"fgd-alterra-29/business/follows"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/follows/request"
	"fgd-alterra-29/controllers/follows/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FollowController struct {
	FollowUseCase follows.UseCase
}

func NewFollowController(followUseCase follows.UseCase) *FollowController {
	return &FollowController{
		FollowUseCase: followUseCase,
	}
}

func (handler FollowController) GetFollowers(c echo.Context) error {
	GetFollows := request.GetFollows{}
	c.Bind(&GetFollows)

	target_id := GetFollows.Target_id
	my_id := middlewares.ExtractID(c)

	ctx := c.Request().Context()

	follows, err := handler.FollowUseCase.GetFollowers(ctx, target_id, my_id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ToListFollowerList(follows))
}

func (handler FollowController) GetFollowing(c echo.Context) error {
	GetFollows := request.GetFollows{}
	c.Bind(&GetFollows)

	target_id := GetFollows.Target_id
	my_id := middlewares.ExtractID(c)

	ctx := c.Request().Context()

	follows, err := handler.FollowUseCase.GetFollowing(ctx, target_id, my_id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ToListFollowingList(follows))
}

func (handler FollowController) FollowsController(c echo.Context) error {
	id := middlewares.ExtractID(c)
	NewFollow := request.Follow{}

	c.Bind(&NewFollow)
	domain := NewFollow.ToDomain()

	ctx := c.Request().Context()

	follows, err := handler.FollowUseCase.FollowsController(ctx, domain, id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NoDataSuccessResponse(c, follows)
}
