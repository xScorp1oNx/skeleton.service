package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

type Failure struct {
	Message []string `json:"message"`
}

type Fatal struct {
	Message string `json:"message"`
}

func SuccessResponse(ctx context.Context, resp interface{}, httpStatus int, w http.ResponseWriter) {
	message, err := json.Marshal(resp)
	if err != nil {
		FatalResponse(ctx, err.Error(), http.StatusInternalServerError, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_, _ = w.Write(message)
}

func FailureResponse(ctx context.Context, mess []string, httpStatus int, w http.ResponseWriter) {
	resp := Failure{Message: mess}

	message, err := json.Marshal(resp)
	if err != nil {
		FatalResponse(ctx, err.Error(), http.StatusInternalServerError, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_, _ = w.Write(message)
}

func FatalResponse(ctx context.Context, mess string, httpStatus int, w http.ResponseWriter) {
	resp := Fatal{Message: mess}

	message, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_, _ = w.Write(message)
}

func SuccessResponseNoContent(ctx context.Context, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
