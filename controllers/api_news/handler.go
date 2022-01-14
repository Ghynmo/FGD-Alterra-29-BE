package apinews

import (
	apinews "fgd-alterra-29/business/api_news"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/api_news/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type APINewsController struct {
	APINewsUseCase apinews.UseCase
}

func NewAPINewsController(badgeUseCase apinews.UseCase) *APINewsController {
	return &APINewsController{
		APINewsUseCase: badgeUseCase,
	}
}

func (handler APINewsController) GetAPINewsController(c echo.Context) error {
	apikey := c.QueryParam("apikey")
	ctx := c.Request().Context()

	news, err := handler.APINewsUseCase.GetAPINews(ctx, apikey)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToDomain(news))
}
