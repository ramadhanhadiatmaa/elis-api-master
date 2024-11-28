package routes

import (
	"master/controllers"
	"master/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {
	auth := r.Group("/api")

	auth.Get("/lang/", middlewares.Auth, controllers.Show)
	auth.Post("/lang/", middlewares.Auth, controllers.Create)
	auth.Put("/lang/:id", middlewares.Auth, controllers.Update)
	auth.Delete("/lang/:id", middlewares.Auth, controllers.Delete)
}
