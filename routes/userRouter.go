package routes

import (
	"fmt"

	"github.com/Rawan-Temo/Golang-rithimc.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(api fiber.Router) {
	api.Get("/" ,AuthFunc ,controllers.AllUsers)
	api.Post("/" ,AuthFunc ,controllers.CreateUser)
	api.Get("/:id" ,AuthFunc ,controllers.SingleUser)
	api.Patch("/:id" ,AuthFunc ,controllers.UpdateUser)
	api.Delete("/:id" ,AuthFunc ,controllers.DeleteUser)
	api.Post("/login"  ,controllers.Login)

}
func AuthFunc(c *fiber.Ctx) error {
    fmt.Print("Authentication logic here\n")
    return c.Next()
}