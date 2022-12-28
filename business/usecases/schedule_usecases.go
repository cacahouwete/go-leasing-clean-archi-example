package usecases

import (
	"context"
	"time"

	"github.com/rs/zerolog"

	"gitlab.com/alexandrevinet/leasing/business/errors"

	"github.com/google/uuid"
	"gitlab.com/alexandrevinet/leasing/adapters/gateways"
	"gitlab.com/alexandrevinet/leasing/business/dto"
	"gitlab.com/alexandrevinet/leasing/business/entities"
)

type ScheduleUsecases interface {
	GetSchedules(ctx context.Context) ([]entities.Schedule, error)
	CreateNewSchedule(ctx context.Context, dto dto.Schedule) (*entities.Schedule, error)
	GetSchedule(ctx context.Context, id string) (*entities.Schedule, error)
	UpdateSchedule(ctx context.Context, id string, dto dto.ScheduleUpdate) (*entities.Schedule, error)
	DeleteSchedule(ctx context.Context, id string) (bool, error)
}

// scheduleUsecases struct that store all dependencies.
type scheduleUsecases struct {
	l   *zerolog.Logger
	cgw gateways.ScheduleGateway
}

func (cuc scheduleUsecases) GetSchedules(ctx context.Context) ([]entities.Schedule, error) {
	return cuc.cgw.FindAll(ctx)
}

func (cuc scheduleUsecases) CreateNewSchedule(ctx context.Context, dto dto.Schedule) (*entities.Schedule, error) {
	beginAt, errB := time.Parse("2006-01-02T15:04:05Z07:00", dto.BeginAt)
	if errB != nil {
		return nil, errB
	}

	endAt, errE := time.Parse("2006-01-02T15:04:05Z07:00", dto.EndAt)
	if errE != nil {
		return nil, errE
	}

	id := uuid.NewString()

	overlap, err := cuc.cgw.CheckOverlap(ctx, id, dto.CarID, beginAt, endAt)
	if err != nil {
		return nil, err
	}

	if overlap {
		cuc.l.Info().Str("id", id).Str("carId", dto.CarID).Time("beginAt", beginAt).Time("endAt", endAt).Msg("timewindow overlap an other schedule for the same car")

		return nil, errors.ViolationError{
			PropertyPath: "",
			Message:      "timewindow overlap an other schedule for the same car",
			Code:         "timewindow_overlap",
		}
	}

	entity := &entities.Schedule{
		ID:         id,
		BeginAt:    beginAt,
		EndAt:      endAt,
		CustomerID: dto.CustomerID,
		CarID:      dto.CarID,
	}

	errS := cuc.cgw.Save(ctx, entity)
	if errS != nil {
		return nil, errS
	}

	return entity, nil
}

// UpdateSchedule will update the entity with the given payload
// If an overlap is detected it will return a business violation.
func (cuc scheduleUsecases) UpdateSchedule(ctx context.Context, id string, dto dto.ScheduleUpdate) (*entities.Schedule, error) {
	entity, err := cuc.cgw.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	if entity == nil {
		return entity, nil
	}

	beginAt, errB := time.Parse("2006-01-02T15:04:05Z07:00", dto.BeginAt)
	if errB != nil {
		return nil, errB
	}

	endAt, errE := time.Parse("2006-01-02T15:04:05Z07:00", dto.EndAt)
	if errE != nil {
		return nil, errE
	}

	overlap, err := cuc.cgw.CheckOverlap(ctx, id, entity.CarID, beginAt, endAt)
	if err != nil {
		return nil, err
	}

	if overlap {
		cuc.l.Info().Str("id", id).Str("carId", entity.CarID).Time("beginAt", beginAt).Time("endAt", endAt).Msg("timewindow overlap an other schedule for the same car")

		return nil, errors.ViolationError{
			PropertyPath: "",
			Message:      "timewindow overlap an other schedule for the same car",
			Code:         "timewindow_overlap",
		}
	}

	entity.BeginAt = beginAt
	entity.EndAt = endAt

	errS := cuc.cgw.Update(ctx, entity)
	if errS != nil {
		return nil, errS
	}

	return entity, nil
}

func (cuc scheduleUsecases) GetSchedule(ctx context.Context, id string) (*entities.Schedule, error) {
	return cuc.cgw.FindById(ctx, id)
}

func (cuc scheduleUsecases) DeleteSchedule(ctx context.Context, id string) (bool, error) {
	schedule, err := cuc.cgw.FindById(ctx, id)
	if err != nil {
		return false, err
	}

	if schedule == nil {
		return false, nil
	}

	errD := cuc.cgw.Delete(ctx, schedule)
	if errD != nil {
		return false, errD
	}

	return true, nil
}

// newScheduleUsecases create carUsecases and return it.
func newScheduleUsecases(l *zerolog.Logger, cgw gateways.ScheduleGateway) ScheduleUsecases {
	return &scheduleUsecases{
		l,
		cgw,
	}
}
