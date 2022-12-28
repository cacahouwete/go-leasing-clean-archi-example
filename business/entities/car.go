package entities

import "github.com/uptrace/bun"

type Car struct {
	bun.BaseModel `bun:"table:car,alias:a" swaggerignore:"true"`

	ID        string      `bun:",pk," json:"id"`
	Name      string      `json:"name"`
	Schedules []*Schedule `bun:"rel:has-many,join:id=car_id" json:"schedules"`
}
