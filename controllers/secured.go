package controllers

import (
	"context"
	"log"

	secured "github.com/tonouchi510/goa2-sample/gen/secured"
)

// secured service example implementation.
// The example methods log the requests and return zero values.
type securedSvc struct {
	logger *log.Logger
}

// NewSecured returns the secured service implementation.
func NewSecured(logger *log.Logger) secured.Service {
	return &securedSvc{logger}
}

// Creates a valid JWT token for auth to api.
func (s *securedSvc) Signin(ctx context.Context, p *secured.SigninPayload) (res *secured.GoaJWT, err error) {
	res = &secured.GoaJWT{}
	s.logger.Print("secured.signin")
	return
}
