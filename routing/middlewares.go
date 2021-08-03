package routing

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"skeleton.service/constants"
)

func generateTraceID() string {
	random, err := uuid.NewRandom()
	if err != nil {
		return uuid.NewString()
	}

	return random.String()
}

func appendTracingContext(res http.ResponseWriter, req *http.Request) {
	c := context.WithValue(req.Context(), constants.TracingTraceId, generateTraceID())
	c = context.WithValue(c, constants.TracingClientIp, req.RemoteAddr)

	req.WithContext(c)
}
