package routes

import (
	C "docker/controllers"
	// AC "docker/controllers/admin"
	// AP "docker/controllers/adminPanel"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

func AdminInit(router *fiber.App) {
	admin := router.Group("adminPanel")
	// admin.Get("/",C.Home)
	// admin.Get("/update",C.Update)
	// admin.Get(":model", C.SingleModel)
	admin.Get(":model", C.IndexModel)
	admin.Get(":model/:id", C.SingleModel)
	admin.Post(":model", C.CreateModel)
	// admin.Get(":model/:id", AdvancedLawSearch)
	admin.Put(":model/single/:field", C.UpdateSingleFieldModel)
	admin.Put(":model/:id", C.UpdateModel)
	admin.Delete(":model/:id", C.DeleteModel)
}