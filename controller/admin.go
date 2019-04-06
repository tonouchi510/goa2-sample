package controller

import (
	"context"
	"database/sql"
	"log"

	"github.com/tonouchi510/goa2-sample/gen/admin"
)

// Admin service example implementation.
// The example methods log the requests and return zero values.
type adminsrvc struct {
	logger *log.Logger
	DB     *sql.DB
}

// NewAdmin returns the Admin service implementation.
func NewAdmin(logger *log.Logger, db *sql.DB) admin.Service {
	return &adminsrvc{logger, db}
}

// Statistic of users
func (s *adminsrvc) UserNumber(ctx context.Context) (res *admin.Goa2SampleAdminUserNumber, err error) {
	res = &admin.Goa2SampleAdminUserNumber{}
	s.logger.Print("admin.user_stats")

	st1 := "2018/11"
	st2 := "2018/12"
	st3 := "2019/01"
	c := "key"
	size := "value"

	// StatsPlanetController_Bar: end_implement
	res = &admin.Goa2SampleAdminUserNumber{
		Data: []*admin.Data{
			&admin.Data{Key:&st1, Value:1},
			&admin.Data{Key:&st2, Value:2},
			&admin.Data{Key:&st3, Value:5},
		},
		X: "key",
		Y: "value",
		Color: &c,
		Size: &size,
		Guide: &admin.StatsGuideType{
			X: &admin.StatsLabelType{Label: "年月"},
			Y: &admin.StatsLabelType{Label: "人数"},
		},
	}
	return
}
// List all stored users
func (s *adminsrvc) AdminListUser(ctx context.Context) (res admin.Goa2SampleAdminUserCollection, err error) {
	res = admin.Goa2SampleAdminUserCollection{}
	s.logger.Print("admin.Admin list user")

	rows, err := s.DB.Query("SELECT * FROM users")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id, name, email string
		var createdAt, updatedAt string
		var deletedAt *string
		err = rows.Scan(&id, &name, &email, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			return
		}
		user := &admin.Goa2SampleAdminUser{
			ID: id,
			Name: name,
			Email: email,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		}
		res = append(res, user)
	}
	return
}

// Show user by ID
func (s *adminsrvc) AdminGetUser(ctx context.Context, p *admin.GetUserPayload) (res *admin.Goa2SampleAdminUser, err error) {
	res = &admin.Goa2SampleAdminUser{}
	s.logger.Print("admin.Admin get user")

	var id, name, email string
	var createdAt, updatedAt string
	var deletedAt *string
	err = s.DB.QueryRow("SELECT * FROM users WHERE id=?", p.ID).Scan(&id, &name, &email, &createdAt, &updatedAt, &deletedAt)
	if err != nil {
		return
	}
	res = &admin.Goa2SampleAdminUser{
		ID: id,
		Name: name,
		Email: email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
	return
}

// Add new user and return its ID.
func (s *adminsrvc) AdminCreateUser(ctx context.Context, p *admin.CreateUserPayload) (res string, err error) {
	s.logger.Print("admin.Admin create user")

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
func (s *adminsrvc) AdminUpdateUser(ctx context.Context, p *admin.UpdateUserPayload) (res *admin.Goa2SampleAdminUser, err error) {
	res = &admin.Goa2SampleAdminUser{}
	s.logger.Print("admin.Admin update user")
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
	var createdAt, updatedAt string
	var deletedAt *string
	err = s.DB.QueryRow("SELECT * FROM users WHERE id=?", p.ID).Scan(&id, &name, &email, &createdAt, &updatedAt, &deletedAt)
	if err != nil {
		return
	}
	res = &admin.Goa2SampleAdminUser{
		ID: id,
		Name: name,
		Email: email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
	s.logger.Printf("UPDATE: ID: %s | Name: %s | Email: %s", id, name, email)
	return
}

// Delete user by id.
func (s *adminsrvc) AdminDeleteUser(ctx context.Context, p *admin.DeleteUserPayload) (err error) {
	s.logger.Print("admin.Admin delete user")
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
