package controllers

import (
	"echoApi/mydb"
	"github.com/labstack/echo/v4"
	"net/http"
	"echoApi/models"
	// "strconv"
)


// Get all categories
func GetCategories(c echo.Context) error {
	var categories []models.Category
	db := mydb.DB()
	db.Find(&categories)
	return c.JSON(http.StatusOK, categories)
}

// Get products within category by ID
func GetCategory(c echo.Context) error {
	id := c.Param("id")
	var category models.Category
	var products []models.Product
	db := mydb.DB()
	if result := db.First(&category, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Category not found"})
	}

	if result := db.Where("Category_Id = ?", id).Find(&products); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "No items in this category"})
	}

	return c.JSON(http.StatusOK, products)
}

// Create new category
func CreateCategory(c echo.Context) error {
	category := new(models.Category)
	name := c.Param("name")
	category.Name = name
	if err := c.Bind(category); err != nil {
		return err
	}
	db := mydb.DB()
	db.Create(&category)
	return c.JSON(http.StatusCreated, category)
}

// Update existing category
func UpdateCategory(c echo.Context) error {
	id := c.Param("id")
	Name := c.QueryParam("Name")
	var category models.Category
	db := mydb.DB()
	if result := db.First(&category, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Category not found"})
	}
	updatedCategory := new(models.Category)
	if err := c.Bind(updatedCategory); err != nil {
		return err
	}
	category.Name = Name
	db.Save(&category)
	return c.JSON(http.StatusOK, category)
}

// Delete a category by ID
func DeleteCategory(c echo.Context) error {
	id := c.Param("id")
	var category models.Category
	db := mydb.DB()
	if result := db.First(&category, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Category not found"})
	}
	db.Delete(&category)
	return c.NoContent(http.StatusNoContent)
}