package routes

import (
	"fmt"

	"github.com/Rawan-Temo/Golang-rithimc.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(api fiber.Router) {
	api.Get("/" ,AuthFunc ,controllers.AllUsers)
	api.Post("/" ,AuthFunc ,controllers.CreateUser)
}
func AuthFunc(c *fiber.Ctx) error {
    fmt.Print("Authentication logic here\n")
    return c.Next()
}