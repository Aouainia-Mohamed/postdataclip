package Models

import (
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Dataclip struct {
	//gorm.Model

	Dataclipkey string         `gorm:"primary_key" json:"dataclipkey"`
	Groupename  postgres.Jsonb `json:"groupename"`
	Name        postgres.Jsonb `json:"name"`
	Description postgres.Jsonb `json:"description"`
	//Descriptionarg postgres.Jsonb `json:"argdesc"`
	Sqltext  string `json:"sqltext"`
	Nargs    int    `json:"nargs"`
	Saastype string `json:"saastype"`
}
