package controllers

import (
	"time"

	database "github.com/Rawan-Temo/Golang-rithimc.git/dataBase"
	"github.com/Rawan-Temo/Golang-rithimc.git/models"
	"github.com/gofiber/fiber/v2"
)
type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func createResponseUser(userModel models.User) User{
	return User{
		ID:userModel.ID,
		CreatedAt:userModel.CreatedAt,
		FirstName:userModel.FirstName,
		LastName:userModel.LastName,
	}
}

func AllUsers(c *fiber.Ctx) error {
	var users []models.User
    result := database.Database.Db.Find(&users) // fills users slice

     if result.Error != nil {
    return c.Status(500).JSON(fiber.Map{
        "status": "error",
        "message": result.Error.Error(),
      })
    }

return c.Status(200).JSON(fiber.Map{
    "status": "success",
    "users":  users,
})
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&user)
	responseUser := createResponseUser(user)

	return c.Status(201).JSON(fiber.Map{
		"status": "success",
		"user":   responseUser,
	})
}