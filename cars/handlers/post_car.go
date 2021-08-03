package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"skeleton.service/cars/service"
)

// PostCar godoc
// @Summary Create car
// @Description Create car with the input payload
// @Tags cars
// @Accept json
// @Produce json
// @Param post_request body service.PostRequest true "Request for create car"
// @Success 201 {object} service.HalResponse
// @Failure 422 {object} Failure
// @Failure 400,500 {object} Fatal
// @Router /car [post]
func PostCar(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	request := service.PostRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		FatalResponse(ctx, err.Error(), http.StatusBadRequest, w)
		return
	}

	mess, err := service.ValidatePostRequest(ctx, request)
	if err != nil {
		FatalResponse(ctx, err.Error(), http.StatusInternalServerError, w)
		return
	}

	if len(mess) > 0 {
		FailureResponse(ctx, mess, http.StatusUnprocessableEntity, w)
		return
	}

	response, err := service.PostCar(ctx, request)
	if err != nil {
		FatalResponse(ctx, err.Error(), http.StatusInternalServerError, w)
		return
	}

	SuccessResponse(ctx, response, http.StatusCreated, w)
}
