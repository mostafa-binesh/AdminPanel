package controllers

import (
	D "docker/database"
	// F "docker/database/filters"
	M "docker/models"
	U "docker/utils"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func TranslationTest(c *fiber.Ctx) error {
	type User struct {
		Username string `validate:"required" json:"username"`
		Password string `validate:"required" json:"wirdpass"`
	}

	user := User{Username: "kurosh79"}
	if errs := U.Validate(user); errs != nil {
		return c.JSON(fiber.Map{"errors": U.Validate(user)})
	}
	return c.JSON(fiber.Map{"msg": "everything is fine"})
}
func PaginationTest(c *fiber.Ctx) error {
	enactments := []M.Law{}
	pagination := new(U.Pagination)
	if err := c.QueryParser(pagination); err != nil {
		U.ResErr(c, err.Error())
	}
	D.DB().Where("type = ?", 3).Scopes(U.Paginate(enactments, pagination)).Find(&enactments)
	return c.JSON(fiber.Map{
		"meta": pagination,
		"data": enactments,
	})
}

func DevAllUsers(c *fiber.Ctx) error {
	users := []M.User{}
	pagination := new(U.Pagination)
	if err := c.QueryParser(pagination); err != nil {
		U.ResErr(c, err.Error())
	}
	D.DB().Scopes(U.Paginate(users, pagination)).Find(&users)
	return c.JSON(fiber.Map{
		"meta": pagination,
		"data": users,
	})
}
func UploadFile(c *fiber.Ctx) error {
	type Upload struct {
		FirstName string `json:"firstName" validate:"required"`
		LastName  string `json:"lastName" validate:"required"`
		File      string `json:"file" validate:"required"`
	}
	payload := new(Upload)
	if err := c.BodyParser(payload); err != nil {
		return c.JSON(fiber.Map{
			"error": err,
		})
	}
	file, err := c.FormFile("file")
	// if err != nil {
	// ! if file not exists, we get error: there is no uploaded file associated with the given key
	// 	return c.JSON(fiber.Map{"error": err.Error()})
	// }
	if file != nil {
		payload.File = file.Filename
	}
	if errs := U.Validate(payload); errs != nil {
		return c.Status(400).JSON(fiber.Map{"errors": errs})
	}
	// check if file with this name already exists
	if U.FileExistenceCheck(file.Filename, U.UploadLocation) {
		return U.ResErr(c, "file already exists")
	}
	// ! file extension check
	if !(U.HasImageSuffixCheck(file.Filename) || U.HasSuffixCheck(file.Filename, []string{"pdf"})) {
		return c.SendString("file should be image or pdf! please fix it")

	}
	// Save file to disk
	err = c.SaveFile(file, fmt.Sprintf(U.UploadLocation+"/%s", file.Filename))
	if err != nil {
		return U.ResErr(c, "cannot save | "+err.Error())
	}
	return c.JSON(fiber.Map{"msg": "فایل آپلود شد"})
}

func ExistenceCheck(c *fiber.Ctx) error {
	filename := c.FormValue("fileName")
	directory := c.FormValue("dir")
	if _, err := os.Stat(directory + "/" + filename); os.IsNotExist(err) {
		return c.SendString("File does not exist")
	} else {
		return c.SendString("File exists")
	}
}
func GormG(c *fiber.Ctx) error {
	type pashm struct {
		Name         string `json:"name" validate:"required,dunique=users.name"` // users table, name column
		PersonalCode string `json:"personalCode" validate:"required,dexists=users"`
	}
	payload := new(pashm)
	// parse payload
	if err := c.BodyParser(payload); err != nil {
		U.ResErr(c, err.Error())
	}
	// ! if you're in edit and wanna ignore the user's information rows
	// ! - you need to pass the id to validation function as well
	// ! -- eg. the user's phoneNumber is 1234 and you've used dunique in phoneNumber field
	// ! --- but if you check the user's row, you'll get the user's phoneNumber and unique validation will fail
	// ! ---- but you don't want this. so you need to ignore that specific id
	if errs := U.Validate(payload); errs != nil {
		return c.Status(400).JSON(fiber.Map{"errors": errs})
	}
	return c.SendString("no error")
}
