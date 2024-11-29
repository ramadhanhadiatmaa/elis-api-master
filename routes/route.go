package routes

import (
	"master/controllers"
	"master/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {
	auth := r.Group("/api")

	auth.Get("/lang/", middlewares.Auth, controllers.Show)
	auth.Get("/lang/:id", middlewares.Auth, controllers.Index)
	auth.Post("/lang/", middlewares.Auth, controllers.Create)
	auth.Put("/lang/:id", middlewares.Auth, controllers.Update)
	auth.Delete("/lang/:id", middlewares.Auth, controllers.Delete)

	auth.Get("/des/", middlewares.Auth, controllers.ShowKel)
	auth.Post("/des/", middlewares.Auth, controllers.CreateKel)
	auth.Put("/des/:kd_kel", middlewares.Auth, controllers.UpdateKel)
	auth.Delete("/des/:kd_kel", middlewares.Auth, controllers.DeleteKel)
}
