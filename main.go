package main

import (


	"github.com/labstack/echo/v4"
	"echoApi/mydb"
	"echoApi/controllers"
)



func main() {
	// Initialize Echo
	e := echo.New()
	mydb.DatabaseInit()

	// Product routes
	e.GET("/products", controllers.GetProducts)
	e.GET("/products/:id", controllers.GetProduct)
	e.POST("/products", controllers.CreateProduct)
	e.PUT("/products/:id", controllers.UpdateProduct)
	e.DELETE("/products/:id", controllers.DeleteProduct)

	// Category routes
	e.GET("/categories", controllers.GetCategories)
	e.GET("/categories/:id", controllers.GetCategory)
	e.POST("/categories/:id", controllers.GetCategories)
	e.PUT("/categories/:id", controllers.GetCategories)

	// Cart routes
	e.GET("/cart", controllers.GetCart)
	e.POST("/cart/products/:id", controllers.AddProductToCart)
	e.DELETE("/cart/products/:id", controllers.RemoveProductFromCart)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

