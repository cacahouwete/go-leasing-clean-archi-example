package gateways

import (
	"github.com/uptrace/bun"
	"gitlab.com/alexandrevinet/leasing/business/entities"
)

type CarGateway interface {
	GenericGateway[entities.Car]
}

// carGateway struct to store all dependencies.
type carGateway struct {
	genericGateway[entities.Car]
}

// newCarGateway create a new carGateway and return it.
func newCarGateway(db *bun.DB) CarGateway {
	return &carGateway{
		genericGateway: genericGateway[entities.Car]{db},
	}
}
