package routes

import (
	"github.com/faisal-990/age/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func RoutesInit(app fiber.Router, h *handler.UserHandler) {
	//wire all the possible routes here
	users := app.Group("/users")
	users.Post("/", h.CreateUser)
	users.Get("/:id", h.GetUser)
	users.Put("/:id", h.UpdateUser)
	users.Delete("/:id", h.DeleteUser)
	users.Get("/", h.ListUsers)
}
