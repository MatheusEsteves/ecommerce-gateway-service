package services

import (
	"bytes"
	"encoding/json"

	"github.com/ecommerce-gateway-service/clients"
	"github.com/ecommerce-gateway-service/models"
	"github.com/labstack/echo/v4"
)

type CartProductsService interface {
	Get(echoContext echo.Context) ([]models.CartProductData, error)
	Save(echoContext echo.Context, cartProduct *models.CartProductData) error
}

type CartProductsServiceImp struct {
	httpRequestClient clients.HttpRequestClient
}

func NewCartProductsService(httpRequestClient clients.HttpRequestClient) CartProductsService {
	return &CartProductsServiceImp{httpRequestClient: httpRequestClient}
}

func (c *CartProductsServiceImp) Get(echoContext echo.Context) ([]models.CartProductData, error) {
	return nil, nil
}

func (c *CartProductsServiceImp) Save(echoContext echo.Context, cartProduct *models.CartProductData) error {
	dataBytes, err := json.Marshal(cartProduct)
	if err != nil {
		return err
	}

	if _, err := c.httpRequestClient.DoRequest("POST", "http://localhost:8081/cart-products", bytes.NewReader(dataBytes)); err != nil {
		return err
	}

	return nil
}
