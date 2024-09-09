package controllers

import (
	"echoApi/mydb"
	"github.com/labstack/echo/v4"
	"net/http"
	"echoApi/models"
	"strconv"
	// "fmt"
)

type result struct {
	Amount  int
	Name    string
	Price   int
	Item_Id int
  }

// GetCart retrieves the cart and its products
func GetCart(c echo.Context) error {
	var products []models.Product
	db := mydb.DB()
	db.Find(&products)

	var ids []int
	var amounts []int
	db.Model(&models.Cart_Item{}).Pluck("Item_Id", &ids)
	db.Model(&models.Cart_Item{}).Pluck("Amount", &amounts)

	if len(ids) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "No products in the cart"})
	}
	
	db.Where(ids).Find(&products)

	var results []result

	for index, element := range products {
		r := result{
			Amount:   amounts[index],
			Name:     element.Name,
			Price:    element.Price,
			Item_Id:  ids[index],
		}
		results = append(results, r)
	}

	return c.JSON(http.StatusOK, results)
}

// AddProductToCart adds a product to the cart
func AddProductToCart(c echo.Context) error {
	id := c.Param("id")
	// var cart []models.Cart_Item
	var cart_item models.Cart_Item
	var product models.Product
	db := mydb.DB()
	// db.Preload("Products").FirstOrCreate(&cart)
	// var product models.Product

	if result := db.Where("Item_Id = ?", id).First(&cart_item); result.Error != nil {
		id_checkd, iderr := strconv.Atoi(id)
		if iderr != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid product id"})
		}

		if result := db.First(&product, id); result.Error != nil {
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
		}

		new_cart_item := new(models.Cart_Item)
		new_cart_item.Item_Id = id_checkd
		new_cart_item.Amount = 1
		db.Create(&new_cart_item)
		return c.JSON(http.StatusOK, new_cart_item)
	}
	updatedCart_Item := new(models.Cart_Item)
	if err := c.Bind(updatedCart_Item); err != nil {
		return err
	}
	cart_item.Amount++
	db.Save(&cart_item)
	return c.JSON(http.StatusOK, cart_item)
}

// RemoveProductFromCart removes a product from the cart
func RemoveProductFromCart(c echo.Context) error {
	id := c.Param("id")
	var cart_item models.Cart_Item
	db := mydb.DB()
	// db.Preload("Products").FirstOrCreate(&cart)
	// var product models.Product
	if result := db.Where("Item_Id = ?", id).First(&cart_item); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
	}

	if cart_item.Amount == 1 {
		db.Delete(&cart_item)
		return c.JSON(http.StatusOK, nil)
	} else {
		updatedCart_Item := new(models.Cart_Item)
		if err := c.Bind(updatedCart_Item); err != nil {
			return err
		}
		cart_item.Amount -= 1
		db.Save(&cart_item)
		return c.JSON(http.StatusOK, cart_item)
	}
}
