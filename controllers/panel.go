package controllers

import (
	// "fmt"

	"fmt"
	"reflect"
	"strconv"
	"strings"

	// "strings"

	D "docker/database"
	M "docker/models"
	F "docker/database/filters"
	"github.com/gofiber/fiber/v2"
	// M "docker/models"
	U "docker/utils"
)

type User struct {
	ID       int
	Username string
	Email    string
}

// var ModelLink = map[string]interface{}{
// 	"users": M.User{},
// 	"students": Student{
// 		FavoriteColor: "dsad",
// 	},
// 	"laws": &M.Law{},
// }
// var ModelLink = map[string]interface{}{
// 	// "users": M.,
// 	"user":  &M.User{},
// 	"users": &[]M.User{},
// }

// func Home(c *fiber.Ctx) error {
// 	// fmt.Println(ModelLink)
// 	// user := &M.User{}
// 	fmt.Printf("ta inja1\n")
// 	user := ModelLink["user"]
// 	// var user []ModelLink["user"]
// 	fmt.Printf("ta inja2\n")
// 	// result := D.DB().Where("id = ?", c.Query("id")).Find(&user)
// 	fmt.Printf("user: %v\n\n\n", user)
// 	result := D.DB().Where("id = ?", c.Query("id")).Find(&user)
// 	// result := D.DB().Find(user)
// 	fmt.Printf("ta inja3\n")
// 	if result.RowsAffected == 0 { // can check the same condition with user.Name == ""
// 		return U.ResErr(c, "کاربر وجود ندارد")
// 	}
// 	return c.JSON(fiber.Map{"users": user})
// 	// return c.JSON(fiber.Map{"users" : D.DB().First(ModelLink["laws"])})
// }
func Update(c *fiber.Ctx) error {
	// import (
	// 	"reflect"
	// )

	// // Assume you have retrieved the user instance you want to update and stored it in a variable called "user"

	// // Get a reflect.Value object for the user instance
	userValue := reflect.ValueOf(ModelLink["user"]).Elem()

	// Get a reflect.Value object for the "Name" field
	usernameField := userValue.FieldByName("Name")

	// // Update the value of the "username" field
	newUsername := c.Query("name")
	if usernameField.IsValid() && usernameField.CanSet() {
		usernameField.SetString(newUsername)
	}
	return c.JSON(fiber.Map{"newUsername": ModelLink["user"]})
}

var ModelLink = map[string]M.Model{
	"posts": {
		SingleReference: &M.Post{},
		ArrayReference:  &[]M.Post{},
	},
	"tags" : {
		SingleReference: &M.Tag{},
		ArrayReference:  &[]M.Tag{},
	},
}
// ! adminPanel routes
func SingleModel(c *fiber.Ctx) error {
	ref := ModelLink[c.Params("model")].SingleReference
	result := D.DB().Scopes(F.RelationPreload(ModelLink[c.Params("model")].Fields)).Find(ref, c.Params("id"))
	if result.Error != nil {
		return U.DBError(c, result.Error)
	}
	if result.RowsAffected == 0 {
		return U.ResErr(c, "داده یافت نشد")
	}
	return c.JSON(fiber.Map{"data": ref})
}
func IndexModel(c *fiber.Ctx) error {
	// return c.JSON(fiber.Map{"model": InitAdminModels(ModelLink["users"])})
	ref := ModelLink[c.Params("model")].ArrayReference
	pagination := U.ParsedPagination(c)
	D.DB().Scopes(U.Paginate(ref, pagination)).Find(ref)
	return c.JSON(fiber.Map{"data": ref, "meta": pagination})
	// return c.JSON(fiber.Map{"model": ModelLink[c.Params("model")], "fields": InitAdminModel(ModelLink[c.Params("model")])})
}
func UpdateSingleFieldModel(c *fiber.Ctx) error {
	fmt.Printf("model: %v\n", ModelLink[c.Params("model")])
	userValue := reflect.ValueOf(ModelLink[c.Params("model")].SingleReference).Elem()
	// Get a reflect.Value object for the "Name" field
	usernameField := userValue.FieldByName(c.Params("field"))

	// // Update the value of the "username" field
	newUsername := c.FormValue(c.Params("field"))
	fmt.Printf("query value: %v\n", newUsername)
	if usernameField.IsValid() && usernameField.CanSet() {
		usernameField.SetString(newUsername)
	} else {
		return c.JSON(fiber.Map{"error": "cannot set"})
	}
	return c.JSON(fiber.Map{"newUsername": ModelLink[c.Params("model")].SingleReference})
}
func DeleteModel(c *fiber.Ctx) error {
	ref := ModelLink[c.Params("model")].SingleReference
	result := D.DB().Delete(ref, c.Params("id"))
	if result.Error != nil {
		return U.DBError(c, result.Error)
	}
	return U.ResErr(c, "عملیات با موفقیت انجام شد")
}
func CreateModel(c *fiber.Ctx) error {
	ref := ModelLink[c.Params("model")].SingleReference
	elem := reflect.ValueOf(ref).Elem()
	fmt.Printf("elem: %v\n", elem)
	fields := ModelLink[c.Params("model")].Fields
	// go through the fields and get the value from FormValue function and then set it
	for _, value := range fields {
		elemField := elem.FieldByName(value.Name)
		if c.FormValue(value.Name) == "" {
			continue
		}
		if elemField.IsValid() && elemField.CanSet() {
			fmt.Printf("trying to set %v |  Element kind: %v\n", value.Name, elemField.Kind())
			switch elemField.Kind() {
			case reflect.String:
				elemField.SetString(c.FormValue(value.Name))
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				intValue, err := strconv.ParseInt(c.FormValue(value.Name), 10, 64)
				if err != nil {
					return U.ResErr(c, err.Error())
				}
				elemField.SetInt(intValue)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				uIntValue, err := strconv.ParseUint(c.FormValue(value.Name), 10, 64)
				if err != nil {
					return U.ResErr(c, err.Error())
				}
				elemField.SetUint(uIntValue)
			case reflect.Float32, reflect.Float64:
				floatValue, err := strconv.ParseFloat(c.FormValue(value.Name), 64)
				if err != nil {
					return U.ResErr(c, err.Error())
				}
				elemField.SetFloat(floatValue)
			case reflect.Bool:
				boolValue, err := strconv.ParseBool(c.FormValue(value.Name))
				if err != nil {
					return U.ResErr(c, err.Error())
				}
				elemField.SetBool(boolValue)
			}
		}
	}
	fmt.Printf("till here\n")
	U.PrintStruct(ref)
	result := D.DB().Create(ref)
	if result.Error != nil {
		return U.DBError(c, result.Error)
	}
	return U.ResErr(c, "عملیات با موفقیت انجام شد")
}
func UpdateModel(c *fiber.Ctx) error {
	ref := ModelLink[c.Params("model")].SingleReference
	result := D.DB().Find(ref, c.Params("id"))
	if result.Error != nil {
		return U.DBError(c, result.Error)
	}
	if result.RowsAffected == 0 {
		return U.ResErr(c, "داده یافت نشد")
	}
	elem := reflect.ValueOf(ref).Elem()
	fmt.Printf("elem: %v\n", elem)
	fields := ModelLink[c.Params("model")].Fields
	// go through the fields and get the value from FormValue function and then set it
	for _, value := range fields {
		elemField := elem.FieldByName(value.Name)
		// fmt.Printf("trying to set %v\n", value.Name)
		if c.FormValue(value.Name) == "" {
			continue
		}
		// fmt.Printf("valueName: %v, isValid: %v, canSet: %v, Kind: %v \n", value.Name, elemField.IsValid(), elemField.CanSet(), elemField.Kind())
		if elemField.IsValid() && elemField.CanSet() {
			switch elemField.Kind() {
			case reflect.String:
				elemField.SetString(c.FormValue(value.Name))
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				intValue, err := strconv.ParseInt(c.FormValue(value.Name), 10, 64)
				if err != nil {
					return U.ResErr(c, err.Error())
				}
				elemField.SetInt(intValue)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				uIntValue, err := strconv.ParseUint(c.FormValue(value.Name), 10, 64)
				if err != nil {
					return U.ResErr(c, err.Error())
				}
				elemField.SetUint(uIntValue)
			case reflect.Float32, reflect.Float64:
				floatValue, err := strconv.ParseFloat(c.FormValue(value.Name), 64)
				if err != nil {
					return U.ResErr(c, err.Error())
				}
				elemField.SetFloat(floatValue)
			case reflect.Bool:
				boolValue, err := strconv.ParseBool(c.FormValue(value.Name))
				if err != nil {
					return U.ResErr(c, err.Error())
				}
				elemField.SetBool(boolValue)
			}
		}
	}
	fmt.Printf("Printing update model status:\n")
	U.PrintStruct(ref)
	result = D.DB().Save(ref)
	if result.Error != nil {
		return U.DBError(c, result.Error)
	}
	if result.RowsAffected == 0 {
		return U.ResErr(c, "داده یافت نشد")
	}
	return U.ResMessage(c, "عملیات با موفقیت انجام شد")
	// // ? first, you need to make sure InitAdminModels syncs all models > ! done
	// // ? then go through each field and see if the form value exists, just update it
	// fmt.Printf("model: %v\n", ModelLink[c.Params("model")])
	// userValue := reflect.ValueOf(ModelLink[c.Params("model")].SingleReference).Elem()
	// // Get a reflect.Value object for the "Name" field
	// usernameField := userValue.FieldByName(c.Params("field"))

	// // // Update the value of the "username" field
	// newUsername := c.FormValue(c.Params("field"))
	// fmt.Printf("query value: %v\n", newUsername)
	// if usernameField.IsValid() && usernameField.CanSet() {
	// 	usernameField.SetString(newUsername)
	// } else {
	// 	return c.JSON(fiber.Map{"error": "cannot set"})
	// }
	// return c.JSON(fiber.Map{"newUsername": ModelLink[c.Params("model")].SingleReference})
}

// func InitAdminModels(models ...Model) {
// func InitAdminModels(student *Student) []Field {
func InitAdminModels() {
	// ! cannot change a field directly!
	for key, value := range ModelLink {
		value.Fields = InitAdminModel(value)
		ModelLink[key] = value
	}
}

func InitAdminModel(model M.Model) []M.Field {
	fields := []M.Field{}
	v := reflect.ValueOf(model.SingleReference).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := M.Field{}
		field.ShowOnTable = true // default is true, we check if it's false in admin, table, then we set it to false
		property := t.Field(i)
		if adminProps, ok := t.Field(i).Tag.Lookup("admin"); ok {
			adminParts := strings.Split(adminProps, ",")
			for i := 0; i < len(adminParts); i++ {
				reflectValues := strings.Split(adminParts[i], ":")
				if len(reflectValues) < 2 {
					panic(fmt.Sprintf("adminPanel: property %s is not correct", adminParts[i]))
				}
				switch reflectValues[0] {
				case "type":
					// ! todo add validation
					// ! type should be in [ text, string, numeric ]
					field.Type = reflectValues[1]
				case "name":
					field.ShowName = reflectValues[1]
				case "table": // default value is true
					if reflectValues[1] == "false" {
						field.ShowOnTable = false
					}
				case "rel":
					field.RelName = reflectValues[1]
					if property.Type.Kind() == reflect.Slice {
						field.RelType = U.HasMany
					} else if property.Type.Kind() == reflect.Int || property.Type.Kind() == reflect.Int8 || property.Type.Kind() == reflect.Int16 || property.Type.Kind() == reflect.Int32 || property.Type.Kind() == reflect.Int64 {
						field.RelType = U.BelongsTo
					}

					fmt.Printf("reflect relationship type: %v\n", reflect.TypeOf(property).Kind())
				default:
				}
			}
		}
		if field.Type == "" {
			field.Type = property.Type.String()
			// field.Type = property.Type.Name()
		}
		if field.Name == "" {
			field.Name = property.Name
		}
		if field.ShowName == "" {
			field.ShowName = property.Name
		}
		fields = append(fields, field)
	}
	fmt.Printf("field type: %v\n", reflect.TypeOf(v).Kind())
	return fields
}

type AdminPanel struct {
	Name   string  `json:"name"`
	Models []M.Model `json:"models"`
}


// ! how to handle relationships in the admin panel:
// ! developer should mention "rel" in the admin attr.s and "rel" value should be a model already submitted
// ! if property is an int, so the relationship is belongsTo
// ! but if property is an array, the rel. is hasMany
// var FileTypes = map[string]uint16{
// 	"plan":        1,
// 	"certificate": 2,
// 	"attachment":  3,
// }
