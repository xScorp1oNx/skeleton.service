package handlers

import (
	"context"
	"net/http"
	"skeleton.service/cars/service"
)

// GetCar godoc
// @Summary Get car
// @Description Get car by ID
// @Tags cars
// @Accept json
// @Produce json
// @Param id query string true "Car ID"
// @Success 200 {object} service.HalResponse
// @Failure 400,404 {object} Failure
// @Failure 500 {object} Fatal
// @Router /car [get]
func GetCar(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var mess []string
	if r.URL.Query().Get("id") == "" {
		mess = append(mess, "Query param 'id' is required")
	}
	if len(mess) != 0 {
		FailureResponse(ctx, mess, http.StatusBadRequest, w)
		return
	}

	response, err := service.GetCarByID(ctx, r.URL.Query().Get("id"))
	if err != nil {
		FatalResponse(ctx, err.Error(), http.StatusNotFound, w)
		return
	}

	SuccessResponse(ctx, response, http.StatusOK, w)
}
