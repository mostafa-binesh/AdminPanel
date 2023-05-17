package models

import (
	U "docker/utils"
)

type Model struct {
	// Struct interface{} `json:"struct"`
	SingleReference interface{} // ! required
	ArrayReference  interface{} // ! required
	Fields          []Field     `json:"fields"`     // ! this field get completed automatically
	ShowOnMenu      bool        `json:"showOnMenu"` // ! not required, default: yes
	Create          bool        `json:"create"`     // ! not required, default: yes
	Read            bool        `json:"read"`       // ! not required, default: yes
	Update          bool        `json:"update"`     // ! not required, default: yes
	Delete          bool        `json:"delete"`     // ! not required, default: yes
}
type Field struct {
	Name        string         // get the variable name
	ShowName    string         // get from admin, name, if not exist, get the variable name
	Type        string         // get from admin, type, if not exist, get the variable type
	ShowOnTable bool           // get from admin, table, if not exist, set to true
	RelName     string         // relation name, get from admin, rel
	RelType     U.RelationType // if variable is int: belongsTo, if v. is array: hasMany
	Value       any
}
