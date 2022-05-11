package routes

import (
	"github.com/gofiber/fiber/v2"
	"recipes.com/controllers" //add this
)

func UserRoute(recipes *fiber.App) {
	user := recipes.Group("/user")
	user.Post("/", controllers.CreateUser) //add this
	user.Get("/", controllers.GetAllUser)
	user.Get("/:id", controllers.GetUser)

	user.Put("/:id", controllers.UpdateUser)    //add this
	user.Delete("/:id", controllers.DeleteUser) //delete this
}
