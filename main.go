package main

import (
	"net/http"

	"github.com/ecommerce-gateway-service/clients"
	"github.com/ecommerce-gateway-service/handlers"
	"github.com/ecommerce-gateway-service/services"
	"github.com/labstack/echo/v4"
)

func main() {
	echoInstance := echo.New()

	client := &http.Client{}
	httpRequestClient := clients.NewHttpRequestClient(client)
	checkoutService := services.NewCheckoutService(httpRequestClient)
	checkoutHandler := handlers.NewCheckoutHandler(checkoutService)

	cartProductsService := services.NewCartProductsService(httpRequestClient)
	cartProductsHandler := handlers.NewCartProductsHandler(cartProductsService)

	gatewayRoute := echoInstance.Group("/ecommerce")
	gatewayRoute.POST("/checkout", checkoutHandler.Checkout)

	cartProductsRoute := gatewayRoute.Group("/cart-products")
	cartProductsRoute.POST("", cartProductsHandler.Save)
	cartProductsRoute.GET("", cartProductsHandler.Get)

	echoInstance.Start(":8082")
}
