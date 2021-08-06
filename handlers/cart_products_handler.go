package handlers

import (
	"net/http"

	"github.com/ecommerce-gateway-service/models"
	"github.com/ecommerce-gateway-service/services"
	"github.com/labstack/echo/v4"
)

type CartProductsHandler interface {
	Get(echoContext echo.Context) error
	Save(echoContext echo.Context) error
}

type CartProductsHandlerImp struct {
	service services.CartProductsService
}

func NewCartProductsHandler(service services.CartProductsService) CartProductsHandler {
	return &CartProductsHandlerImp{
		service: service,
	}
}

func (c *CartProductsHandlerImp) Get(echoContext echo.Context) error {
	products, err := c.service.Get(echoContext)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return echoContext.JSON(http.StatusOK, products)
}

func (c *CartProductsHandlerImp) Save(echoContext echo.Context) error {
	cartProduct := models.CartProductData{}

	if err := echoContext.Bind(&cartProduct); err != nil {
		return echoContext.JSON(http.StatusBadRequest, "Bad request")
	}
	if err := c.service.Save(echoContext, &cartProduct); err != nil {
		return echoContext.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return echoContext.JSON(http.StatusCreated, "Product saved sucessfully")
}
