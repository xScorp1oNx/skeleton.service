package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	validators2 "skeleton.service/validators"
)

func ValidatePostRequest(ctx context.Context, request PostRequest) ([]string, error) {
	var messages []string

	if err := validators2.Validator.GetValidator().Struct(request); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			messages = append(
				messages,
				e.Translate(validators2.Validator.GetTranslator()),
			)
		}
	}

	return messages, nil
}

func GetCarByID(ctx context.Context, id string) (*HalResponse, error) {
	car, err := getCarByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &HalResponse{
		Data:     car,
		Embedded: Embedded{},
		Links:    SelfURL{},
		Status:   "completed",
	}, nil
}

func PostCar(ctx context.Context, request PostRequest) (*HalResponse, error) {
	car, err := postCar(ctx, request)
	if err != nil {
		return nil, err
	}

	return &HalResponse{
		Data:     car,
		Embedded: Embedded{},
		Links:    SelfURL{},
		Status:   "completed",
	}, nil
}
