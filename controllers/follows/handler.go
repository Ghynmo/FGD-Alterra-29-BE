package follows

import (
	"fgd-alterra-29/business/follows"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/follows/responses"
	"net/http"
	"strconv"

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
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	follows, err := handler.FollowUseCase.GetFollowers(ctx, id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ToListFollowerList(follows))
}
