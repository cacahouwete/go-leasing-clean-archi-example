package entities

import "github.com/uptrace/bun"

type Customer struct {
	bun.BaseModel `bun:"table:customer,alias:c" swaggerignore:"true"`

	ID        string      `bun:",pk," json:"id"`
	Name      string      `json:"name"`
	Schedules []*Schedule `bun:"rel:has-many,join:id=customer_id" json:"schedules"`
}
