package gateways

import (
	"github.com/uptrace/bun"
)

// Gateways struct to store all Gateway.
type Gateways struct {
	Car      CarGateway
	Customer CustomerGateway
	Schedule ScheduleGateway
}

// New will create and return Gateways.
func New(db *bun.DB) *Gateways {
	return &Gateways{
		Car:      newCarGateway(db),
		Customer: newCustomerGateway(db),
		Schedule: newScheduleGateway(db),
	}
}
