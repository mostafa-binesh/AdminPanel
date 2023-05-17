package adminPanel

// import (
// 	"github.com/gofiber/fiber/v2"
// 	// M "docker/models"
// 	M "docker/models"
// )
// var sdasas fiber.Map{
// 	"sdasdas" : M.Law{},
// }
// func Home(c *fiber.Ctx) error {
// 	ModelLink :=  map[string]interface{
// 		"users": M.Law,
// 		"laws": "asdsa",
// 	}
// 	return c.JSON(fiber.Map{ "sdad" : M.User{}})
// }
// type AdminPanel struct {
// 	Name string `json:"name"`
// 	Models []Model `json:"models"`
// }
// type Model struct {
// 	Struct interface{} `json:"struct"`
// 	Fields []Field `json:"fields"` // should be generated automatically 
// 	ShowOnMenu bool `json:"showOnMenu"`
// 	Create bool `json:"create"`
// 	Read bool `json:"read"`
// 	Update bool `json:"update"`
// 	Delete bool `json:"delete"`
// }
// type Field struct {
// 	Name string // get from admin, name
// 	Type string  // get from admin, type
// }
// type Student struct {
// 	Id uint64 `json:"id" admin:"type:text,name:name Karbari"`
// 	FavoriteColor string `json:"favoriteColor" admin:"type:text,name:name Karbari"`
// 	UserID uint64 `json:"userID" admin:"rel:users"` // if type is int, so relationship is belongsTo
// 	Laws []M.Law `json:"books" admin:"rel:laws"`
// }

// // var FileTypes = map[string]uint16{
// // 	"plan":        1,
// // 	"certificate": 2,
// // 	"attachment":  3,
// // }