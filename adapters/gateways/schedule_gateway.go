package gateways

import (
	"context"
	"time"

	"github.com/uptrace/bun"
	"gitlab.com/alexandrevinet/leasing/business/entities"
)

type ScheduleGateway interface {
	GenericGateway[entities.Schedule]
	CheckOverlap(ctx context.Context, scheduleId, carId string, beginAt, endAt time.Time) (bool, error)
}

// scheduleGateway struct to store all dependencies.
type scheduleGateway struct {
	genericGateway[entities.Schedule]
}

// CheckOverlap will check if a car have been schedule in part of timewindow given
// Return true if there are one or more overlap.
func (cgw scheduleGateway) CheckOverlap(ctx context.Context, scheduleId, carId string, beginAt, endAt time.Time) (bool, error) {
	entity := new(entities.Schedule)

	nb, err := cgw.db.NewSelect().Model(entity).Where("id != ? AND car_id = ? AND end_at > ? AND begin_at < ?", scheduleId, carId, beginAt, endAt).Count(ctx)
	if err != nil {
		return false, err
	}

	return nb > 0, err
}

// newScheduleGateway create a new ScheduleGateway and return it.
func newScheduleGateway(db *bun.DB) ScheduleGateway {
	return &scheduleGateway{
		genericGateway: genericGateway[entities.Schedule]{db},
	}
}
