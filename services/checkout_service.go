package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/ecommerce-gateway-service/clients"
	"github.com/ecommerce-gateway-service/models"
	"github.com/labstack/echo/v4"
)

type CheckoutService interface {
	Checkout(echoContext echo.Context, checkout *models.CheckoutData) (string, error)
}

type CheckoutServiceImp struct {
	httpRequestClient clients.HttpRequestClient
}

func NewCheckoutService(httpRequestClient clients.HttpRequestClient) CheckoutService {
	return &CheckoutServiceImp{httpRequestClient: httpRequestClient}
}

func (c *CheckoutServiceImp) Checkout(echoContext echo.Context, checkout *models.CheckoutData) (string, error) {
	dataBytes, err := json.Marshal(checkout)
	if err != nil {
		return "", err
	}

	response, err := c.httpRequestClient.DoRequest("POST", "http://localhost:8080/checkout", bytes.NewReader(dataBytes))
	if err != nil {
		return "", err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	checkoutResponseData := models.CheckoutResponseData{}
	if err := json.Unmarshal(responseBytes, &checkoutResponseData); err != nil {
		return "", err
	}

	return checkoutResponseData.PurchaseCode, nil
}
