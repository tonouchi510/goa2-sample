package controller

import (
	"context"
	"database/sql"
	"log"

	users "github.com/tonouchi510/goa2-sample/gen/users"
)

// Users service example implementation.
// The example methods log the requests and return zero values.
type userssrvc struct {
	logger *log.Logger
	DB     *sql.DB
}

// NewUsers returns the Users service implementation.
func NewUsers(logger *log.Logger, db *sql.DB) users.Service {
	return &userssrvc{logger, db}
}

// List all stored users
func (s *userssrvc) ListUser(ctx context.Context) (res users.Goa2SampleUserCollection, err error) {
	res = users.Goa2SampleUserCollection{}
	s.logger.Print("users.List")

	rows, err := s.DB.Query("SELECT id, name, email FROM users")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id, name, email string
		user := &users.Goa2SampleUser{}
		err = rows.Scan(&id, &name, &email)
		if err != nil {
			return
		}
		user.ID = id
		user.Name = name
		user.Email = email
		res = append(res, user)
	}
	return
}

// Show user by ID
func (s *userssrvc) GetUser(ctx context.Context, p *users.GetUserPayload) (res *users.Goa2SampleUser, err error) {
	s.logger.Print("users.Get")

	var id, name, email string
	err = s.DB.QueryRow("SELECT id, name, email FROM users WHERE id=?", p.ID).Scan(&id, &name, &email)
	if err != nil {
		return
	}
	res = &users.Goa2SampleUser{ID: id, Name: name, Email: email}
	return
}

// Add new user and return its ID.
func (s *userssrvc) CreateUser(ctx context.Context, p *users.CreateUserPayload) (res string, err error) {
	s.logger.Print("users.Create")
	stmt, err := s.DB.Prepare("INSERT INTO users(id, name, email) VALUES(?,?,?)")
	defer stmt.Close()
	if err != nil {
		return
	}
	_, err = stmt.Exec(p.ID, p.Name, p.Email)
	if err != nil {
		return
	}
	s.logger.Printf("INSERT: ID: %s | Name: %s | Email: %s", p.ID, p.Name, p.Email)
	res = p.ID
	return
}

// Update user item.
func (s *userssrvc) UpdateUser(ctx context.Context, p *users.UpdateUserPayload) (res *users.Goa2SampleUser, err error) {
	res = &users.Goa2SampleUser{}
	s.logger.Print("users.Update")
	var stmt *sql.Stmt
	if p.Name != nil && p.Email != nil {
		stmt, err = s.DB.Prepare("UPDATE users SET name=?, email=? WHERE id=?")
		_, err = stmt.Exec(p.Name, p.Email, p.ID)
	} else if p.Name != nil {
		stmt, err = s.DB.Prepare("UPDATE users SET name=? WHERE id=?")
		_, err = stmt.Exec(p.Name, p.ID)
	} else if p.Email != nil {
		stmt, err = s.DB.Prepare("UPDATE users SET email=? WHERE id=?")
		_, err = stmt.Exec(p.Email, p.ID)
	}
	if err != nil {
		return
	}
	defer stmt.Close()

	var id, name, email string
	err = s.DB.QueryRow("SELECT id, name, email FROM users WHERE id=?", p.ID).Scan(&id, &name, &email)
	if err != nil {
		return
	}
	res = &users.Goa2SampleUser{ID: id, Name: name, Email: email}
	s.logger.Printf("UPDATE: ID: %s | Name: %s | Email: %s", id, name, email)
	return
}

// Delete user by id.
func (s *userssrvc) DeleteUser(ctx context.Context, p *users.DeleteUserPayload) (err error) {
	s.logger.Print("users.Delete")
	stmt, err := s.DB.Prepare("DELETE FROM users WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	_, err = stmt.Exec(p.ID)
	if err != nil {
		return
	}
	s.logger.Printf("DELETE: Id: %s", p.ID)
	return
}
