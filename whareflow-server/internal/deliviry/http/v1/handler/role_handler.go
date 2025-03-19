package handler

import (
	"encoding/json"
	"fmt"
	delivery "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/model"
	"github.com/Miroslovelife/whareflow/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type RoleHandler interface {
	GetEmployerRole(echo.Context) error
	GetAllEmployerRoles(echo.Context) error
	GiveRoleForEmployer(echo.Context) error
	RemoveRoleForEmployer(echo.Context) error
	ChangeRoleForEmployer(echo.Context) error
}

type IRoleHandler struct {
	permUsecase usecase.PermissionUsecase
}

func NewIRolehandler(permUsecase usecase.PermissionUsecase) *IRoleHandler {
	return &IRoleHandler{
		permUsecase: permUsecase,
	}
}

//func (rh *IRoleHandler) GetEmployerRole(c echo.Context) error {
//	userId := c.Get("x-user-id")
//	warehouseId, err := strconv.Atoi(c.Param("warehouse_id"))
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, "Can't parse warehouse id")
//	}
//	roleId, err := strconv.Atoi(c.Param("role"))
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, "Can't parse role id")
//	}
//
//	role, err := rh.permUsecase.GetUserRole()
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, "Error when getting the role")
//	}
//
//	return c.JSON(http.StatusOK, role)
//}

// GiveRoleForEmployer godoc
// @Summary Выдача роли работнику
// @Description Выдает роль работнику
// @Tags roles
// @Accept			json
// @Produce		json
// @Param			warehouse_id	path		int	true	"warehouse id"
// @Param request body delivery.RoleReq true "Данные для роли"
// @Success 200 {object} map[string]string "message: role success created for user"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security		ApiKeyAuth
// @Router /role/{warehouse_id} [post]
func (rh *IRoleHandler) GiveRoleForEmployer(c echo.Context) error {
	ownerId := c.Get("x-user-id").(string)
	warehouseId, err := strconv.Atoi(c.Param("warehouse_id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't parse warehouse id")
	}

	reqBody := new(delivery.RoleReq)

	if err = c.Bind(reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	err = rh.permUsecase.CreateRole(reqBody, uint(warehouseId), ownerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Can't create role for user: %s", reqBody.UserName))
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("role success created for user: %s", reqBody.UserName))
}

// GetAllPermissionTypes godoc
// @Summary Получение всех типов прав
// @Description Возвращает все типы прав
// @Tags roles
// @Accept			json
// @Produce		json
// @Success 200 {object} map[string]string "[]delivery.PermissionResponse"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security		ApiKeyAuth
// @Router /role [get]
func (rh *IRoleHandler) GetAllPermissionTypes(c echo.Context) error {
	permissionsType, err := rh.permUsecase.GetAllPermissions()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't get permission types")
	}

	return c.JSON(http.StatusOK, permissionsType)
}

// GetAllUserPermissionOnWh godoc
// @Summary Получение всех прав пользователя
// @Description Возвращает все права пользователя
// @Tags roles
// @Accept			json
// @Produce		json
// @Param request body map[string]string true "Данные для роли"
// @Success 200 {object} map[string]string "[]delivery.PermissionResponse"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security		ApiKeyAuth
// @Router /role/{warehouse_id} [post]
func (rh *IRoleHandler) GetAllUserPermissionOnWh(c echo.Context) error {
	warehouseId, err := strconv.Atoi(c.Param("warehouse_id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't parse warehouse id")
	}

	ownerId := c.Get("x-user-id").(string)

	var reqBody map[string]string

	if err := json.NewDecoder(c.Request().Body).Decode(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
	}

	username, ok := reqBody["username"]
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "username is required"})
	}

	permissions, err := rh.permUsecase.GetAllUserPermission(uint(warehouseId), ownerId, username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Can't get permission for user: %s", username))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
	    "permissions": permissions,
	})
}
