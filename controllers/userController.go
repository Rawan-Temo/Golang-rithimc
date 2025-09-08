package controllers

import (
	"fmt"
	"time"

	database "github.com/Rawan-Temo/Golang-rithimc.git/dataBase"
	"github.com/Rawan-Temo/Golang-rithimc.git/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)
func hashPassword (password string) (string ,error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes) , err 
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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
     hashedPassword, err := hashPassword(user.Password)
       if err != nil {
	return c.Status(500).JSON(fiber.Map{
		"status": "error",
		"message": "could not hash password",
	})
        }

	user.Password = hashedPassword
	fmt.Println(user)
	database.Database.Db.Create(&user)


	return c.Status(201).JSON(fiber.Map{
		"status": "success",
		"user":   user,
	})

}
func SingleUser(c *fiber.Ctx) error {
	fmt.Println("getting single user")
	return nil
}
func UpdateUser(c *fiber.Ctx) error {
	fmt.Println("getting single user")
	return nil
}

func DeleteUser(c *fiber.Ctx) error {
	fmt.Println("getting single user")
	return nil
}
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
    Identity string `json:"identity"` // username or email
    Password string `json:"password"`
}

var input LoginInput
if err := c.BodyParser(&input); err != nil {
    return c.Status(400).JSON(fiber.Map{"status":"error","message": err.Error()})
}
var user models.User
result := database.Database.Db.Where("username = ?", input.Identity).First(&user)
if result.Error != nil {
    return c.Status(404).JSON(fiber.Map{
        "status":  "error",
        "message": "user not found",
    })
}
if !CheckPasswordHash(input.Password, user.Password) {
    return c.Status(401).JSON(fiber.Map{
        "status":  "error",
        "message": "invalid credentials",
    })
}
    token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("test"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}


	newUser := user

	newUser.Password = ""
	return c.Status(200).JSON(fiber.Map{"status": "success",
	"user":newUser,
	"message": "Success login", "data": t})
}