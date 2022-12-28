//nolint:dupl // businesses are dupl because it's a simple example
package usecases

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/google/uuid"
	"gitlab.com/alexandrevinet/leasing/adapters/gateways"
	"gitlab.com/alexandrevinet/leasing/business/dto"
	"gitlab.com/alexandrevinet/leasing/business/entities"
)

type CustomerUsecases interface {
	GetCustomers(ctx context.Context) ([]entities.Customer, error)
	CreateNewCustomer(ctx context.Context, dto dto.Customer) (*entities.Customer, error)
	GetCustomer(ctx context.Context, id string) (*entities.Customer, error)
	UpdateCustomer(ctx context.Context, id string, dto dto.Customer) (*entities.Customer, error)
	DeleteCustomer(ctx context.Context, id string) (bool, error)
}

// customerUsecases struct that store all dependencies.
type customerUsecases struct {
	l   *zerolog.Logger
	cgw gateways.CustomerGateway
}

func (cuc customerUsecases) GetCustomers(ctx context.Context) ([]entities.Customer, error) {
	return cuc.cgw.FindAll(ctx)
}

func (cuc customerUsecases) CreateNewCustomer(ctx context.Context, dto dto.Customer) (*entities.Customer, error) {
	entity := &entities.Customer{
		ID:   uuid.NewString(),
		Name: dto.Name,
	}

	err := cuc.cgw.Save(ctx, entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (cuc customerUsecases) UpdateCustomer(ctx context.Context, id string, dto dto.Customer) (*entities.Customer, error) {
	entity, err := cuc.cgw.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	if entity == nil {
		return entity, nil
	}

	entity.Name = dto.Name

	errS := cuc.cgw.Update(ctx, entity)
	if errS != nil {
		return nil, errS
	}

	return entity, nil
}

func (cuc customerUsecases) GetCustomer(ctx context.Context, id string) (*entities.Customer, error) {
	return cuc.cgw.FindById(ctx, id)
}

func (cuc customerUsecases) DeleteCustomer(ctx context.Context, id string) (bool, error) {
	customer, err := cuc.cgw.FindById(ctx, id)
	if err != nil {
		return false, err
	}

	if customer == nil {
		return false, nil
	}

	errD := cuc.cgw.Delete(ctx, customer)
	if errD != nil {
		return false, errD
	}

	return true, nil
}

// newCustomerUsecases create carUsecases and return it.
func newCustomerUsecases(l *zerolog.Logger, cgw gateways.CustomerGateway) CustomerUsecases {
	return &customerUsecases{
		l,
		cgw,
	}
}
