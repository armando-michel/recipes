package recipes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"recipes.com/database"
	"recipes.com/routes"
)

var recipe = fiber.New()

func ConnectDB() {
	//  Conectar a la base de datos
	database.ConnectDb()

}

func Init(port string) {
	// Inicializar la aplicacion de Recetas
	recipe.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Rutas de mi aplicacion
	routes.Routes(recipe)

	// Levantar el servidor web en el puerto ...
	recipe.Listen(port)
}
