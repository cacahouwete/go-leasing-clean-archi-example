package gateways

import (
	"context"

	"github.com/uptrace/bun"
)

type GenericGateway[T interface{}] interface {
	FindAll(ctx context.Context) ([]T, error)
	Save(ctx context.Context, car *T) error
	Update(ctx context.Context, car *T) error
	Delete(ctx context.Context, car *T) error
	FindById(ctx context.Context, id string) (*T, error)
}

// genericGateway struct to store all dependencies.
type genericGateway[T interface{}] struct {
	db *bun.DB
}

// FindAll will return all entities in db.
func (ggw genericGateway[T]) FindAll(ctx context.Context) ([]T, error) {
	var items []T

	err := ggw.db.NewSelect().Model(&items).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// Save will insert a new entity in db.
func (ggw genericGateway[T]) Save(ctx context.Context, entity *T) error {
	_, err := ggw.db.NewInsert().Model(entity).Exec(ctx)

	return err
}

// Update will update an entity in db.
func (ggw genericGateway[T]) Update(ctx context.Context, entity *T) error {
	_, err := ggw.db.NewUpdate().Model(entity).WherePK().Exec(ctx)

	return err
}

// Delete will delete an entity in db.
func (ggw genericGateway[T]) Delete(ctx context.Context, entity *T) error {
	_, err := ggw.db.NewDelete().Model(entity).WherePK().Exec(ctx)

	return err
}

// FindById will find an entity in db. It will return nil if nothing found.
func (ggw genericGateway[T]) FindById(ctx context.Context, id string) (*T, error) {
	var entity *T

	nb, err := ggw.db.NewSelect().Model(entity).Where("id = ?", id).Count(ctx)
	if err != nil {
		return nil, err
	}

	if nb == 0 {
		return entity, nil
	}

	entity = new(T)

	err = ggw.db.NewSelect().Model(entity).Where("id = ?", id).Limit(1).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
