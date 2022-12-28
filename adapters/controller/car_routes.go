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

// carRoutes struct to store all dependencies.
type carRoutes struct {
	l   *zerolog.Logger
	cuc usecases.CarUsecases
}

// @Summary     Retrieves the collection of Cars resources.
// @Description Retrieves the collection of Cars resources.
// @Tags  	    Cars
// @Accept      json
// @Produce     json
// @Success     200 {object} httputils.ResponseCollection[entities.Car]{member=[]entities.Car}
// @Failure     500 {object} httputils.ResponseError
// @Router      /cars [get].
func (r *carRoutes) getCars(c *gin.Context) {
	cars, err := r.cuc.GetCars(c)
	if err != nil {
		httputils.ErrorResponse(c, err)

		return
	}

	c.JSON(http.StatusOK, httputils.ResponseCollection[entities.Car]{
		Member:     cars,
		TotalItems: uint(len(cars)),
	})
}

// @Summary     Creates a Cars resource.
// @Description Creates a Cars resource.
// @Tags  	    Cars
// @Accept      json
// @Produce     json
// @Param		car     body dto.Car true "Car payload"
// @Success     201 {object} entities.Car
// @Failure     422 {object} httputils.ResponseViolations
// @Failure     500 {object} httputils.ResponseError
// @Router      /cars [post].
func (r *carRoutes) postCar(c *gin.Context) {
	d := dto.Car{}

	if err := c.ShouldBindJSON(&d); err != nil {
		httputils.ErrorResponse(c, err)

		return
	}

	car, errUc := r.cuc.CreateNewCar(c, d)
	if errUc != nil {
		httputils.ErrorResponse(c, errUc)

		return
	}

	c.JSON(http.StatusCreated, car)
}

// @Summary     Retrieves a Cars resource.
// @Description Retrieves a Cars resource.
// @Tags  	    Cars
// @Accept      json
// @Produce     json
// @Param       id  path     string true  "Cars ID"
// @Success     200 {object} entities.Car
// @Failure     404 {object} httputils.ResponseError
// @Failure     500 {object} httputils.ResponseError
// @Router      /cars/{id} [get].
func (r *carRoutes) getCar(c *gin.Context) {
	id, found := c.Params.Get("id")
	if !found {
		httputils.BadRequestResponse(c)

		return
	}

	car, err := r.cuc.GetCar(c, id)
	if err != nil {
		httputils.ErrorResponse(c, err)

		return
	}

	if car == nil {
		httputils.NotFoundResponse(c)

		return
	}

	c.JSON(http.StatusOK, car)
}

// @Summary     Replaces the Cars resource.
// @Description Replaces the Cars resource.
// @Tags  	    Cars
// @Accept      json
// @Produce     json
// @Param       id  path     string true  "Cars ID"
// @Param		car     body dto.Car true "Car payload"
// @Success     200 {object} entities.Car
// @Failure     404 {object} httputils.ResponseError
// @Failure     422 {object} httputils.ResponseViolations
// @Failure     500 {object} httputils.ResponseError
// @Router      /cars/{id} [put].
func (r *carRoutes) putCar(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		httputils.BadRequestResponse(c)

		return
	}

	entity, errUc := r.cuc.GetCar(c, id)
	if errUc != nil {
		httputils.ErrorResponse(c, errUc)

		return
	}

	if entity == nil {
		httputils.NotFoundResponse(c)

		return
	}

	d := dto.Car{
		Name: entity.Name,
	}

	if err := c.ShouldBindJSON(&d); err != nil {
		httputils.ErrorResponse(c, err)

		return
	}

	car, errUc := r.cuc.UpdateCar(c, id, d)
	if errUc != nil {
		httputils.ErrorResponse(c, errUc)

		return
	}

	if car == nil {
		httputils.NotFoundResponse(c)

		return
	}

	c.JSON(http.StatusOK, car)
}

// @Summary     Removes the Cars resource.
// @Description Removes the Cars resource.
// @Tags  	    Cars
// @Accept      json
// @Produce     json
// @Param       id  path     string true  "Cars ID"
// @Success     204
// @Failure     404 {object} httputils.ResponseError
// @Failure     500 {object} httputils.ResponseError
// @Router      /cars/{id} [delete].
func (r *carRoutes) deleteCar(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		httputils.BadRequestResponse(c)

		return
	}

	found, err := r.cuc.DeleteCar(c, id)
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

// newCarRoutes create carRoutes and register all routes in gin handler.
func newCarRoutes(handler *gin.RouterGroup, cuc usecases.CarUsecases, l *zerolog.Logger) {
	r := &carRoutes{l, cuc}

	handler.GET("/cars", r.getCars)
	handler.POST("/cars", r.postCar)
	handler.GET("/cars/:id", r.getCar)
	handler.PUT("/cars/:id", r.putCar)
	handler.DELETE("/cars/:id", r.deleteCar)
}
