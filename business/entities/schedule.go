package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type Schedule struct {
	bun.BaseModel `bun:"table:schedule,alias:s" swaggerignore:"true"`

	ID         string    `bun:",pk," json:"id"`
	BeginAt    time.Time `bun:"begin_at" json:"beginAt"`
	EndAt      time.Time `bun:"end_at" json:"endAt"`
	CustomerID string    `json:"customerId"`
	Customer   Customer  `bun:"rel:belongs-to,join:customer_id=id" json:"-" swaggerignore:"true"`
	CarID      string    `json:"carId"`
	Car        Car       `bun:"rel:belongs-to,join:car_id=id" json:"-" swaggerignore:"true"`
}
