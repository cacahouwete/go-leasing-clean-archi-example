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

type CarUsecases interface {
	GetCars(ctx context.Context) ([]entities.Car, error)
	CreateNewCar(ctx context.Context, dto dto.Car) (*entities.Car, error)
	GetCar(ctx context.Context, id string) (*entities.Car, error)
	UpdateCar(ctx context.Context, id string, dto dto.Car) (*entities.Car, error)
	DeleteCar(ctx context.Context, id string) (bool, error)
}

// carUsecases struct that store all dependencies.
type carUsecases struct {
	l   *zerolog.Logger
	cgw gateways.CarGateway
}

func (cuc carUsecases) GetCars(ctx context.Context) ([]entities.Car, error) {
	return cuc.cgw.FindAll(ctx)
}

func (cuc carUsecases) CreateNewCar(ctx context.Context, dto dto.Car) (*entities.Car, error) {
	entity := &entities.Car{
		ID:   uuid.NewString(),
		Name: dto.Name,
	}

	err := cuc.cgw.Save(ctx, entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (cuc carUsecases) UpdateCar(ctx context.Context, id string, dto dto.Car) (*entities.Car, error) {
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

func (cuc carUsecases) GetCar(ctx context.Context, id string) (*entities.Car, error) {
	return cuc.cgw.FindById(ctx, id)
}

func (cuc carUsecases) DeleteCar(ctx context.Context, id string) (bool, error) {
	car, err := cuc.cgw.FindById(ctx, id)
	if err != nil {
		return false, err
	}

	if car == nil {
		return false, nil
	}

	errD := cuc.cgw.Delete(ctx, car)
	if errD != nil {
		return false, errD
	}

	return true, nil
}

// newCarUsecases create carUsecases and return it.
func newCarUsecases(l *zerolog.Logger, cgw gateways.CarGateway) CarUsecases {
	return &carUsecases{
		l,
		cgw,
	}
}
