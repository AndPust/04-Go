package controllers

import (
	"echoApi/mydb"
	"github.com/labstack/echo/v4"
	"net/http"
	"echoApi/models"
	"strconv"
)

// Get all products
func GetProducts(c echo.Context) error {
	var products []models.Product
	db := mydb.DB()
	db.Find(&products)
	return c.JSON(http.StatusOK, products)
}

// Get product by ID
func GetProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	db := mydb.DB()
	if result := db.First(&product, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
	}
	return c.JSON(http.StatusOK, product)
}

// Create new product
func CreateProduct(c echo.Context) error {
	Name := c.QueryParam("Name")
	Price, perr := strconv.Atoi(c.QueryParam("Price"))
	Category_Id, cerr := strconv.Atoi(c.QueryParam("Category_Id"))

	if perr != nil || cerr != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid product fields"})
	}

	product := new(models.Product)
	product.Name = Name
	product.Price = Price
	product.Category_Id = Category_Id
	db := mydb.DB()
	if err := c.Bind(product); err != nil {
		return err
	}
	db.Create(&product)
	return c.JSON(http.StatusCreated, product)
}

// Update existing product
func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	Name := c.QueryParam("Name")
	Price, perr := strconv.Atoi(c.QueryParam("Price"))
	Category_Id, cerr := strconv.Atoi(c.QueryParam("Category_Id"))

	if perr != nil || cerr != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid product fields"})
	}

	var product models.Product
	db := mydb.DB()
	if result := db.First(&product, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
	}
	updatedProduct := new(models.Product)
	if err := c.Bind(updatedProduct); err != nil {
		return err
	}
	product.Name = Name
	product.Price = Price
	product.Category_Id = Category_Id
	db.Save(&product)
	return c.JSON(http.StatusOK, product)
}

// Delete a product by ID
func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	db := mydb.DB()
	if result := db.First(&product, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
	}
	db.Delete(&product)
	return c.NoContent(http.StatusNoContent)
}