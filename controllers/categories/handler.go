package categories

import (
	"fgd-alterra-29/business/categories"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/categories/request"
	"fgd-alterra-29/controllers/categories/responses"
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

func (handler CategoryController) GetCategoriesController(c echo.Context) error {
	ctx := c.Request().Context()

	categories, err := handler.CategoryUseCase.GetCategoriesController(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToListCategories(categories))
}

func (handler CategoryController) GetUserActiveInCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	categories, err := handler.CategoryUseCase.GetUserActiveInCategory(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, categories)
}

func (handler CategoryController) CreateCategoryController(c echo.Context) error {
	var NewCategory = request.AddCategory{}
	c.Bind(&NewCategory)

	domain := NewCategory.ToDomain()

	ctx := c.Request().Context()

	categories, err := handler.CategoryUseCase.CreateCategoriesController(ctx, domain)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NoDataSuccessResponse(c, categories)
}
