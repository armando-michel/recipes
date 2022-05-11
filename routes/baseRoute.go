package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(recipe *fiber.App) {
	// rutos en root
	recipe.Get("/", hello)

	// rutas de Usarios
	UserRoute(recipe) //<----- crearemos la ruta usuarios
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hola Armando Michel 👋!")
}
