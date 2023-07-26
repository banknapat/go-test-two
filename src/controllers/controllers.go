package controllers

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Product struct {
	ID       int    `json:"id"`
	Name	 string `json:"name"`
	Price	 string `json:"price"`
}

var products = []Product{
	{ ID: 1, Name: "สินค้าที่ 1", Price: "100" },
	{ ID: 2, Name: "สินค้าที่ 2", Price: "200" },
}

// show all products
func GetProducts(c *fiber.Ctx) error {
	return c.JSON(products)
}

// show product by id
func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, product := range products {
		if strconv.Itoa(product.ID) == id {
			return c.JSON(product)
		}
	}
	return c.Status(404).JSON(fiber.Map{
		"message": "Not found",
	})
}

// create product
func CreateProduct(c *fiber.Ctx) error {
	type Request struct {
		Name string `json:"name"`
		Price string `json:"price"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	product := Product{
		ID: len(products) + 1,
		Name: req.Name,
		Price: req.Price,
	}
	products = append(products, product)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "เพิ่มข้อมูลสำเร็จ",
	})
}

// update product
func UpdateProduct(c *fiber.Ctx) error {
	type Request struct {
		Name string `json:"name"`
		Price string `json:"price"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	id := c.Params("id")
	for i, product := range products {
		if strconv.Itoa(product.ID) == id {
			products[i].Name = req.Name
			products[i].Price = req.Price
			return c.JSON(products[i])
		}
	}
	return c.Status(404).JSON(fiber.Map{
		"message": "Not found",
	})
}

// delete product
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, product := range products {
		if strconv.Itoa(product.ID) == id {
			products = append(products[:i], products[i+1:]...)
			return c.SendStatus(204)
		}
	}
	return c.Status(404).JSON(fiber.Map{
		"message": "Not found",
	})
}







