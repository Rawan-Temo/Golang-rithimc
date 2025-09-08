package main

import (
	"fmt"
	"log"

	database "github.com/Rawan-Temo/Golang-rithimc.git/dataBase"
	"github.com/Rawan-Temo/Golang-rithimc.git/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	fmt.Println("hello world!")
    database.ConnectDb()
	app.Get("/api" , welcome )

	app.Route("api/v1/users", routes.UserRouter)
	log.Fatal(app.Listen(":8000"))
}


func welcome(c *fiber.Ctx) error{

	fmt.Println("hello world")

	return c.SendString("welcome")
}