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

// scheduleRoutes struct to store all dependencies.
type scheduleRoutes struct {
	l   *zerolog.Logger
	cuc usecases.ScheduleUsecases
}

// @Summary     Retrieves the collection of Schedules resources.
// @Description Retrieves the collection of Schedules resources.
// @Tags  	    Schedules
// @Accept      json
// @Produce     json
// @Success     200 {object} httputils.ResponseCollection[entities.Schedule]{member=[]entities.Schedule}
// @Failure     500 {object} httputils.ResponseError
// @Router      /schedules [get].
func (r *scheduleRoutes) getSchedules(c *gin.Context) {
	schedules, err := r.cuc.GetSchedules(c)
	if err != nil {
		httputils.ErrorResponse(c, err)

		return
	}

	c.JSON(http.StatusOK, httputils.ResponseCollection[entities.Schedule]{
		Member:     schedules,
		TotalItems: uint(len(schedules)),
	})
}

// @Summary     Creates a Schedules resource.
// @Description Creates a Schedules resource.
// @Tags  	    Schedules
// @Accept      json
// @Produce     json
// @Param		schedule     body dto.Schedule true "Schedule payload"
// @Success     201 {object} entities.Schedule
// @Failure     422 {object} httputils.ResponseViolations
// @Failure     500 {object} httputils.ResponseError
// @Router      /schedules [post].
func (r *scheduleRoutes) postSchedule(c *gin.Context) {
	d := dto.Schedule{}

	if err := c.ShouldBindJSON(&d); err != nil {
		httputils.ErrorResponse(c, err)

		return
	}

	schedule, errUc := r.cuc.CreateNewSchedule(c, d)
	if errUc != nil {
		httputils.ErrorResponse(c, errUc)

		return
	}

	c.JSON(http.StatusCreated, schedule)
}

// @Summary     Retrieves a Schedules resource.
// @Description Retrieves a Schedules resource.
// @Tags  	    Schedules
// @Accept      json
// @Produce     json
// @Param       id  path     string true  "Schedules ID"
// @Success     200 {object} entities.Schedule
// @Failure     404 {object} httputils.ResponseError
// @Failure     500 {object} httputils.ResponseError
// @Router      /schedules/{id} [get].
func (r *scheduleRoutes) getSchedule(c *gin.Context) {
	id, found := c.Params.Get("id")
	if !found {
		httputils.BadRequestResponse(c)

		return
	}

	schedule, err := r.cuc.GetSchedule(c, id)
	if err != nil {
		httputils.ErrorResponse(c, err)

		return
	}

	if schedule == nil {
		httputils.NotFoundResponse(c)

		return
	}

	c.JSON(http.StatusOK, schedule)
}

// @Summary     Replaces the Schedules resource.
// @Description Replaces the Schedules resource.
// @Tags  	    Schedules
// @Accept      json
// @Produce     json
// @Param       id  path     string true  "Schedules ID"
// @Param		schedule     body dto.ScheduleUpdate true "Schedule payload"
// @Success     200 {object} entities.Schedule
// @Failure     404 {object} httputils.ResponseError
// @Failure     422 {object} httputils.ResponseViolations
// @Failure     500 {object} httputils.ResponseError
// @Router      /schedules/{id} [put].
func (r *scheduleRoutes) putSchedule(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		httputils.BadRequestResponse(c)

		return
	}

	entity, errG := r.cuc.GetSchedule(c, id)
	if errG != nil {
		httputils.ErrorResponse(c, errG)

		return
	}

	if entity == nil {
		httputils.NotFoundResponse(c)

		return
	}

	d := dto.ScheduleUpdate{}

	if err := c.ShouldBindJSON(&d); err != nil {
		httputils.ErrorResponse(c, err)

		return
	}

	schedule, errUc := r.cuc.UpdateSchedule(c, id, d)
	if errUc != nil {
		httputils.ErrorResponse(c, errUc)

		return
	}

	if schedule == nil {
		httputils.NotFoundResponse(c)

		return
	}

	c.JSON(http.StatusOK, schedule)
}

// @Summary     Removes the Schedules resource.
// @Description Removes the Schedules resource.
// @Tags  	    Schedules
// @Accept      json
// @Produce     json
// @Param       id  path     string true  "Schedules ID"
// @Success     204
// @Failure     404 {object} httputils.ResponseError
// @Failure     500 {object} httputils.ResponseError
// @Router      /schedules/{id} [delete].
func (r *scheduleRoutes) deleteSchedule(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		httputils.BadRequestResponse(c)

		return
	}

	found, err := r.cuc.DeleteSchedule(c, id)
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

// newScheduleRoutes create scheduleRoutes and register all routes in gin handler.
func newScheduleRoutes(handler *gin.RouterGroup, cuc usecases.ScheduleUsecases, l *zerolog.Logger) {
	r := &scheduleRoutes{l, cuc}

	handler.GET("/schedules", r.getSchedules)
	handler.POST("/schedules", r.postSchedule)
	handler.GET("/schedules/:id", r.getSchedule)
	handler.PUT("/schedules/:id", r.putSchedule)
	handler.DELETE("/schedules/:id", r.deleteSchedule)
}
