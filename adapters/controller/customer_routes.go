//nolint:dupl // router are dupl because it's a simple example
package controller

import (
	"net/http"

	"github.com/rs/zerolog"
	"gitlab.com/alexandrevinet/leasing/business/entities"

	"gitlab.com/alexandrevinet/leasing/adapters/controller/httputils"

	"gitlab.com/alexandrevinet/leasing/business/dto"
	"gitlab.com/alexandrevinet/leasing/business/usecases"

	"github.com/gin-gonic/gin"
)

// customerRoutes struct to store all dependencies.
type customerRoutes struct {
	l   *zerolog.Logger
	cuc usecases.CustomerUsecases
}

// @Summary     Retrieves the collection of Customers resources.
// @Description Retrieves the collection of Customers resources.
// @Tags  	    Customers
// @Accept      json
// @Produce     json
// @Success     200 {object} httputils.ResponseCollection[entities.Customer]{member=[]entities.Customer}
// @Failure     500 {object} httputils.ResponseError
// @Router      /customers [get].
func (r *customerRoutes) getCustomers(c *gin.Context) {
	customers, err := r.cuc.GetCustomers(c)
	if err != nil {
		httputils.ErrorResponse(c, err)

		return
	}

	c.JSON(http.StatusOK, httputils.ResponseCollection[entities.Customer]{
		Member:     customers,
		TotalItems: uint(len(customers)),
	})
}

// @Summary     Creates a Customers resource.
// @Description Creates a Customers resource.
// @Tags  	    Customers
// @Accept      json
// @Produce     json
// @Param		customer     body dto.Customer true "Customer payload"
// @Success     201 {object} entities.Customer
// @Failure     422 {object} httputils.ResponseViolations
// @Failure     500 {object} httputils.ResponseError
// @Router      /customers [post].
func (r *customerRoutes) postCustomer(c *gin.Context) {
	d := dto.Customer{}

	if err := c.ShouldBindJSON(&d); err != nil {
		httputils.ErrorResponse(c, err)

		return
	}

	customer, errUc := r.cuc.CreateNewCustomer(c, d)
	if errUc != nil {
		httputils.ErrorResponse(c, errUc)

		return
	}

	c.JSON(http.StatusCreated, customer)
}

// @Summary     Retrieves a Customers resource.
// @Description Retrieves a Customers resource.
// @Tags  	    Customers
// @Accept      json
// @Produce     json
// @Param       id  path     string true  "Customers ID"
// @Success     200 {object} entities.Customer
// @Failure     404 {object} httputils.ResponseError
// @Failure     500 {object} httputils.ResponseError
// @Router      /customers/{id} [get].
func (r *customerRoutes) getCustomer(c *gin.Context) {
	id, found := c.Params.Get("id")
	if !found {
		httputils.BadRequestResponse(c)

		return
	}

	customer, err := r.cuc.GetCustomer(c, id)
	if err != nil {
		httputils.ErrorResponse(c, err)

		return
	}

	if customer == nil {
		httputils.NotFoundResponse(c)

		return
	}

	c.JSON(http.StatusOK, customer)
}

// @Summary     Replaces the Customers resource.
// @Description Replaces the Customers resource.
// @Tags  	    Customers
// @Accept      json
// @Produce     json
// @Param       id  path     string true  "Customers ID"
// @Param		customer     body dto.Customer true "Customer payload"
// @Success     200 {object} entities.Customer
// @Failure     404 {object} httputils.ResponseError
// @Failure     422 {object} httputils.ResponseViolations
// @Failure     500 {object} httputils.ResponseError
// @Router      /customers/{id} [put].
func (r *customerRoutes) putCustomer(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		httputils.BadRequestResponse(c)

		return
	}

	entity, errUc := r.cuc.GetCustomer(c, id)
	if errUc != nil {
		httputils.ErrorResponse(c, errUc)

		return
	}

	if entity == nil {
		httputils.NotFoundResponse(c)

		return
	}

	d := dto.Customer{
		Name: entity.Name,
	}

	if err := c.ShouldBindJSON(&d); err != nil {
		httputils.ErrorResponse(c, err)

		return
	}

	customer, errUc := r.cuc.UpdateCustomer(c, id, d)
	if errUc != nil {
		httputils.ErrorResponse(c, errUc)

		return
	}

	if customer == nil {
		httputils.NotFoundResponse(c)

		return
	}

	c.JSON(http.StatusOK, customer)
}

// @Summary     Removes the Customers resource.
// @Description Removes the Customers resource.
// @Tags  	    Customers
// @Accept      json
// @Produce     json
// @Param       id  path     string true  "Customers ID"
// @Success     204
// @Failure     404 {object} httputils.ResponseError
// @Failure     500 {object} httputils.ResponseError
// @Router      /customers/{id} [delete].
func (r *customerRoutes) deleteCustomer(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		httputils.BadRequestResponse(c)

		return
	}

	found, err := r.cuc.DeleteCustomer(c, id)
	if err != nil {
		httputils.ErrorResponse(c, err)

		return
	}

	if !found {
		httputils.NotFoundResponse(c)

		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// newCustomerRoutes create customerRoutes and register all routes in gin handler.
func newCustomerRoutes(handler *gin.RouterGroup, cuc usecases.CustomerUsecases, l *zerolog.Logger) {
	r := &customerRoutes{l, cuc}

	handler.GET("/customers", r.getCustomers)
	handler.POST("/customers", r.postCustomer)
	handler.GET("/customers/:id", r.getCustomer)
	handler.PUT("/customers/:id", r.putCustomer)
	handler.DELETE("/customers/:id", r.deleteCustomer)
}
