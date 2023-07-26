package routes

import (
	"github.com/banknapat/go-test-two/src/controllers"
	"github.com/banknapat/go-test-two/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	// show all products
	app.Get("/products", controllers.GetProducts)
	// show product by id
	app.Get("/products/:id", controllers.GetProduct)
	// create new product
	app.Post("/products", controllers.CreateProduct)
	// update product by id
	app.Put("/products/:id", controllers.UpdateProduct)
	// delete product by id
	app.Delete("/products/:id", controllers.DeleteProduct)

	// 404
	app.Use(middlewares.NotFoundErr())
}
