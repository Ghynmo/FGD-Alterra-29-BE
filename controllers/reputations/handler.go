package reputations

import (
	"fgd-alterra-29/business/reputations"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/reputations/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ReputationController struct {
	ReputationUseCase reputations.UseCase
}

func NewReputationController(reputationsUseCase reputations.UseCase) *ReputationController {
	return &ReputationController{
		ReputationUseCase: reputationsUseCase,
	}
}

func (handler ReputationController) CreateReputationController(c echo.Context) error {
	var NewReputation = request.AddReputation{}
	c.Bind(&NewReputation)
	domain := NewReputation.ToDomain()

	ctx := c.Request().Context()

	reputations, err := handler.ReputationUseCase.CreateReputationController(ctx, domain)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NoDataSuccessResponse(c, reputations)
}
