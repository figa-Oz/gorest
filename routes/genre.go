package routes

import (
	"example/gorest/controllers"

	"github.com/gofiber/fiber/v2"
)

func GenreRoute(route fiber.Router) {
	routes := route.Group("genre")
	routes.Get("/", controllers.Index)
	routes.Get("/:id", controllers.Detail)
	routes.Post("/", controllers.Add)
}
