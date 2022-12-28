package usecases

import (
	"github.com/rs/zerolog"
	"gitlab.com/alexandrevinet/leasing/adapters/gateways"
)

// UseCases store all usecases.
type UseCases struct {
	Car      CarUsecases
	Customer CustomerUsecases
	Schedule ScheduleUsecases
}

// New create UseCases and return it.
func New(gw *gateways.Gateways, l *zerolog.Logger) *UseCases {
	return &UseCases{
		Car:      newCarUsecases(l, gw.Car),
		Customer: newCustomerUsecases(l, gw.Customer),
		Schedule: newScheduleUsecases(l, gw.Schedule),
	}
}
