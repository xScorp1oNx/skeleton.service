package service

import "skeleton.service/cars/database"

type PostRequest struct {
	Brand string `json:"brand" validate:"required,min=1,max=100"`
	Model string `json:"model" validate:"required,min=1,max=100"`
}

type HalResponse struct {
	Data     *database.Car `json:"data,omitempty"`
	Embedded Embedded      `json:"_embedded,omitempty"`
	Links    SelfURL       `json:"_links,omitempty"`
	Status   string        `json:"_status"`
}

type Embedded struct{}

type SelfURL struct{}
