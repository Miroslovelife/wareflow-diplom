package handler

import (
	"fmt"
	delivery "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/model"
	"github.com/Miroslovelife/whareflow/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProductHandler interface {
	CreateProduct(echo.Context) error
	GetProduct(echo.Context) error
	GetAllProductsFromZone(echo.Context) error
	GetAllProductsFromWarehouse(echo.Context) error
	UpdateProduct(echo.Context) error
	//DeleteProduct(echo.Context) error
}

type IProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewIProductHandler(productUsecase usecase.ProductUsecase) *IProductHandler {
	return &IProductHandler{
		productUsecase: productUsecase,
	}
}

// CreateProduct godoc
// @Summary Создание продукта
// @Description Создает новый продукт
// @Tags product
// @Accept			json
// @Produce		json
// @Param warehouse_id	path		string	true	"warehouse id"
// @Param zone_id	path		string	true	"zone id"
// @Param request body delivery.ProductModelRequest true "Данные для создания склада"
// @Success 200 {object} map[string]string "message: product success created"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security		ApiKeyAuth
// @Router /warehouse/{warehouse_id}/zone/{zone_id}/product [post]
func (ph *IProductHandler) CreateProduct(c echo.Context) error {
	reqBody := &delivery.ProductModelRequest{}

	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	userId := c.Get("x-user-id").(string)

	warehouseId, err := strconv.Atoi(c.Param("warehouse_id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}



	if err := ph.productUsecase.CreateProduct(reqBody, userId, warehouseId, reqBody.ZoneId); err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "product success created")
}

// GetProduct godoc
// @Summary Получение товара
// @Description Возвращает товар
// @Tags product
// @Accept			json
// @Produce		json
// @Param warehouse_id	path		string	true	"warehouse id"
// @Param zone_id	path		string	true	"zone id"
// @Param product_id	path		string	true	"product id"
// @Success 200 {object} map[string]string "delivery.ProductModelResponse"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security		ApiKeyAuth
// @Router /warehouse/{warehouse_id}/zone/{zone_id}/product/{product_id} [get]
func (ph *IProductHandler) GetProduct(c echo.Context) error {
	userId := c.Get("x-user-id").(string)
	productId := c.Param("product_id")

	fmt.Println(productId, userId)
	product, err := ph.productUsecase.FindProduct(userId, productId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, product)
}

// GetAllProductsFromZone godoc
// @Summary Получение всех товаров с зоны склада
// @Description Возвращает список всех товаров с зоны склада
// @Tags product
// @Accept			json
// @Produce		json
// @Param warehouse_id	path		string	true	"warehouse id"
// @Param zone_id	path		string	true	"zone id"
// @Success 200 {object} map[string]string "[]delivery.ProductModelResponse"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security		ApiKeyAuth
// @Router /warehouse/{warehouse_id}/zone/{zone_id}/product [get]
func (ph *IProductHandler) GetAllProductsFromZone(c echo.Context) error {
	fmt.Println(c.Get("x-user-id"))
	userId := c.Get("x-user-id").(string)

	zoneId, err := strconv.Atoi(c.Param("zone_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	products, err := ph.productUsecase.FindAllProductFromZone(userId, zoneId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, products)
}

// GetAllProductsFromWarehouse godoc
// @Summary Получение всех товаров со склада
// @Description Возвращает список всех товаров со склада
// @Tags product
// @Accept			json
// @Produce		json
// @Param warehouse_id	path		string	true	"warehouse id"
// @Success 200 {object} map[string]string "[]delivery.ProductModelResponse"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security		ApiKeyAuth
// @Router /warehouse/{warehouse_id}/product/{product_id}  [get]
func (ph *IProductHandler) GetAllProductsFromWarehouse(c echo.Context) error {
	fmt.Println(c.Get("x-user-id"))
	userId := c.Get("x-user-id").(string)

	warehouseId, err := strconv.Atoi(c.Param("warehouse_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	products, err := ph.productUsecase.FindAllProductFromWarehouse(userId, warehouseId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, products)
}

// UpdateProduct godoc
// @Summary Обновление продукта
// @Description Обнвляет данные о продукте
// @Tags product
// @Accept			json
// @Produce		json
// @Param warehouse_id	path		string	true	"warehouse id"
// @Param product_id	path		string	true	"warehouse id"
// @Param request body delivery.ProductModelRequest true "Данные для создания склада"
// @Success 200 {object} map[string]string "message: product success updated"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security		ApiKeyAuth
// @Router /warehouse/{warehouse_id}/product/{product_id}  [put]
func (ph *IProductHandler) UpdateProduct(c echo.Context) error {
	userId := c.Get("x-user-id").(string)
	productId := c.Param("product_id")
	reqBody := &delivery.ProductModelRequest{}

	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	warehouseId, err := strconv.Atoi(c.Param("warehouse_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	err = ph.productUsecase.UpdateProduct(reqBody, warehouseId, productId, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "product success updated")
}
