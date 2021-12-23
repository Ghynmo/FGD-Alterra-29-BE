package categories

import (
	"fgd-alterra-29/business/categories"
	"fgd-alterra-29/controllers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	CategoryUseCase categories.UseCase
}

func NewCategoryController(categoryUseCase categories.UseCase) *CategoryController {
	return &CategoryController{
		CategoryUseCase: categoryUseCase,
	}
}

func (handler CategoryController) GetUserActiveInCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	thread, err := handler.CategoryUseCase.GetUserActiveInCategory(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, thread)
}
