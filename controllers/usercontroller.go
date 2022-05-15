package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"recipes.com/database"
	"recipes.com/models"
)

/////////////////////////////////////////////////////////////////////////
func CreateUser(c *fiber.Ctx) error {
	//create user code goes here
	user := new(models.User)

	if err := c.BodyParser(&user); err != nil {
		fmt.Println("error = ", err)
		return c.Status(503).SendString(err.Error())
	}

	user.Prepare()

	err := user.Validate("login")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	encontrar := database.Db.Find(&user, "email = ?", user.Email)
	if encontrar.RowsAffected > 0 {
		return c.Status(400).SendString("Email dupicated")
	}

	userCreated, err := user.SaveUser(database.Db)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(userCreated)
}

/////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////////
func GetUser(c *fiber.Ctx) error {

	uid, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	user := models.User{}
	userGotten, err := user.FindUserByID(database.Db, uint64(uid))
	if err != nil {
		return err
	}
	return c.JSON(userGotten)
}

/////////////////////////////////////////////////////////////////////////
func GetAllUser(c *fiber.Ctx) error {
	user := models.User{}
	users, err := user.FindAllUsers(database.Db)

	if err != nil {
		return c.Status(400).SendString("Error no hay usuarios registrados")
	}
	return c.Status(200).JSON(users)
	// return c.Status(200).SendString("Todos los usuarios")
}

/////////////////////////////////////////////////////////////////////////
func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("Usuario Modificado....")
}

/////////////////////////////////////////////////////////////////////////
func DeleteUser(c *fiber.Ctx) error {
	return c.SendString("Eliminando el usuario... " + c.Params("id"))
}
