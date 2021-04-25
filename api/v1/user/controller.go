package user

import (
	"api-go-hexa/api/common"
	"api-go-hexa/business/user"
	"api-go-hexa/business/user/model"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service user.Service
}

func NewController(service user.Service) *Controller {
	return &Controller{
		service,
	}
}

func (u *Controller) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.BadRequestResponse())
	}
	id := int(idP)
	res, err := u.service.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, common.NewSuccessResponse(res))
}

func (u *Controller) UserRegister(c echo.Context) error {
	payload := new(model.UserModel)
	err := c.Bind(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.BadRequestResponse())
	}

	if err := validator.New().Struct(payload); err != nil {
		return c.JSON(http.StatusBadRequest, common.BadRequestResponse())
	}

	res, err := u.service.UserRegister(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, common.NewSuccessResponse(res))
}

func (u *Controller) UserLogin(c echo.Context) error {
	payload := new(model.UserLoginModel)
	err := c.Bind(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.BadRequestResponse())
	}

	if err := validator.New().Struct(payload); err != nil {
		return c.JSON(http.StatusBadRequest, common.BadRequestResponse())
	}

	res, err := u.service.UserLogin(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, common.NewSuccessResponse(res))
}

func (u *Controller) Update(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.BadRequestResponse())
	}
	id := int(idP)
	payload := new(model.UserModel)
	err = c.Bind(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.BadRequestResponse())
	}

	if err := validator.New().Struct(payload); err != nil {
		return c.JSON(http.StatusBadRequest, common.BadRequestResponse())
	}

	err = u.service.Update(id, payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, common.NewSuccessResponse("Update Success"))
}

func (u *Controller) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.BadRequestResponse())
	}
	id := int(idP)
	err = u.service.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, common.NewSuccessResponse("Delete Success"))
}
