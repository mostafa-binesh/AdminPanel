package filters

import (
	// "docker/utils"
	// "github.com/iancoleman/strcase"

	// "fmt"
	// "reflect"

	// U "docker/utils"
	M "docker/models"
	// "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RelationPreload(fields []M.Field) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, value := range fields {
			// fmt.Printf("checking field %v\n", value)
			if value.RelName == "" {
				continue
			}
			// fmt.Printf("preloading %s\n", value.RelName)	
			db = db.Preload(value.RelName)
		}
		return db
	}
}