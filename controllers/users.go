package goa2sample

import (
	"context"
	"log"

	users "github.com/tonouchi510/goa2-sample/gen/users"
)

// users service example implementation.
// The example methods log the requests and return zero values.
type usersSvc struct {
	logger *log.Logger
}

// NewUsers returns the users service implementation.
func NewUsers(logger *log.Logger) users.Service {
	return &usersSvc{logger}
}

// List all stored users
func (s *usersSvc) List(ctx context.Context) (res users.StoredUserCollection, err error) {
	s.logger.Print("users.list")
	return
}

// Show user by ID
func (s *usersSvc) Show(ctx context.Context, p *users.ShowPayload) (res *users.StoredUser, err error) {
	res = &users.StoredUser{}
	s.logger.Print("users.show")
	return
}

// Add new user and return its ID.
func (s *usersSvc) Add(ctx context.Context, p *users.User) (res int64, err error) {
	s.logger.Print("users.add")
	return
}

// Update user item.
func (s *usersSvc) Update(ctx context.Context, p *users.UpdatePayload) (res *users.StoredUser, err error) {
	res = &users.StoredUser{}
	s.logger.Print("users.update")
	return
}

// Remove user from storage
func (s *usersSvc) Remove(ctx context.Context, p *users.RemovePayload) (err error) {
	s.logger.Print("users.remove")
	return
}
