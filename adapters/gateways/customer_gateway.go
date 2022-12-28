package gateways

import (
	"github.com/uptrace/bun"
	"gitlab.com/alexandrevinet/leasing/business/entities"
)

type CustomerGateway interface {
	GenericGateway[entities.Customer]
}

// customerGateway struct to store all dependencies.
type customerGateway struct {
	genericGateway[entities.Customer]
}

// newCustomerGateway create a new customerGateway and return it.
func newCustomerGateway(db *bun.DB) CustomerGateway {
	return &customerGateway{
		genericGateway: genericGateway[entities.Customer]{db},
	}
}
