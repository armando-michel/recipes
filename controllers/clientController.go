package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// var User models.User

func CreateUser(c *fiber.Ctx) error {

	return c.Status(200).JSON("{user:'Usuario Creado satisfactoriamente ...👋!'}")

}

func GetUser(c *fiber.Ctx) error {
	return c.SendString("Mostrando el usurio.... 👋: " + c.Params("id"))
}

func GetAllUser(c *fiber.Ctx) error {
	return c.SendString("Mostrando todos los usurios.... 👋!")
}

func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("Usuario Modificado....")
}

func DeleteUser(c *fiber.Ctx) error {
	return c.SendString("Eliminando el usuario... " + c.Params("id"))
}
