package controller

import (
	"context"
	"database/sql"
	"log"

	"github.com/tonouchi510/goa2-sample/gen/users"
)

// users service example implementation.
// The example methods log the requests and return zero values.
type usersSvc struct {
	logger *log.Logger
	DB     *sql.DB
}

// NewUsers returns the users service implementation.
func NewUsers(logger *log.Logger, db *sql.DB) users.Service {
	return &usersSvc{logger, db}
}

// List all stored users
func (s *usersSvc) List(ctx context.Context) (res users.StoredUserCollection, err error) {
	s.logger.Print("users.list")

	rows, err := s.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	for rows.Next() {
		var id int64
		var name, email string
		user := &users.StoredUser{}
		err = rows.Scan(&id, &name, &email)
		if err != nil {
			s.logger.Fatal(err.Error())
		}
		user.ID = id
		user.Name = name
		user.Email = email
		res = append(res, user)
	}
	return
}

// Show user by ID
func (s *usersSvc) Show(ctx context.Context, p *users.ShowPayload) (res *users.StoredUser, err error) {
	res = &users.StoredUser{}
	s.logger.Print("users.show")
	rows, err := s.DB.Query("SELECT id, name, email FROM users WHERE id=?", p.ID)
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	for rows.Next() {
		var id int64
		var name, email string
		err = rows.Scan(&id, &name, &email)
		if err != nil {
			s.logger.Fatal(err.Error())
		}
		res.ID = id
		res.Name = name
		res.Email = email
	}
	return
}

// Add new user and return its ID.
func (s *usersSvc) Add(ctx context.Context, p *users.User) (res int64, err error) {
	s.logger.Print("users.add")
	stmt, err := s.DB.Prepare("INSERT INTO users(name, email) VALUES(?,?)")
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	result, err := stmt.Exec(p.Name, p.Email)
	res, _ = result.LastInsertId()
	s.logger.Printf("INSERT: ID: %d | Name: %s | Email: %s", res, p.Name, p.Email)
	return
}

// Update user item.
func (s *usersSvc) Update(ctx context.Context, p *users.UpdatePayload) (res *users.StoredUser, err error) {
	res = &users.StoredUser{}
	s.logger.Print("users.update")
	stmt, err := s.DB.Prepare("UPDATE users SET name=?, email=? WHERE id=?")
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	_, err = stmt.Exec(p.Name, p.Email, p.ID)
	res.ID = p.ID
	res.Name = *p.Name
	res.Email = *p.Email
	s.logger.Printf("UPDATE: ID: %d | Name: %s | Email: %s", p.ID, *p.Name, *p.Email)

	return
}

// Remove user from storage
func (s *usersSvc) Remove(ctx context.Context, p *users.RemovePayload) (err error) {
	s.logger.Print("users.remove")
	stmt, err := s.DB.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	_, err = stmt.Exec(p.ID)
	s.logger.Printf("DELETE: Id: %d", p.ID)
	return
}
