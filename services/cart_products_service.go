package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

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
	response, err := c.httpRequestClient.DoRequest("GET", "http://localhost:8081/cart-products", bytes.NewReader([]byte{}))
	if err != nil {
		return nil, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	products := new([]models.CartProductData)
	if err := json.Unmarshal(responseBytes, products); err != nil {
		return nil, err
	}

	return *products, nil
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
