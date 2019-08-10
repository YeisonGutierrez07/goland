package persons

import "github.com/goland-project/models"

type Person struct {
	models.BaseModel
	IDLocal   string  `json:"id_local" db:"id_local"`
	FirstName string  `json:"first_name" db:"first_name"`
	LastName  string  `json:"last_name" db:"last_name"`
	Address   Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}
