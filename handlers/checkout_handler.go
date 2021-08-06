package handlers

import (
	"fmt"
	"net/http"

	"github.com/ecommerce-gateway-service/models"
	"github.com/ecommerce-gateway-service/services"
	"github.com/labstack/echo/v4"
)

type CheckoutHandler interface {
	Checkout(echoContext echo.Context) error
}

type CheckoutHandlerImp struct {
	service services.CheckoutService
}

func NewCheckoutHandler(service services.CheckoutService) CheckoutHandler {
	return &CheckoutHandlerImp{
		service: service,
	}
}

func (c *CheckoutHandlerImp) Checkout(echoContext echo.Context) error {
	checkoutData := models.CheckoutData{}

	if err := echoContext.Bind(&checkoutData); err != nil {
		return echoContext.JSON(http.StatusBadRequest, "Bad Request")
	}

	purchaseCode, err := c.service.Checkout(echoContext, &checkoutData)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return echoContext.JSON(http.StatusCreated, fmt.Sprintf("Purchase requested sucessfully. ID: %s", purchaseCode))
}
