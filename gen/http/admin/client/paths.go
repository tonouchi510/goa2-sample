// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// HTTP request path constructors for the Admin service.
//
// Command:
// $ goa gen github.com/tonouchi510/goa2-sample/design

package client

import (
	"fmt"
)

// UserNumberAdminPath returns the URL path to the Admin service user_number HTTP endpoint.
func UserNumberAdminPath() string {
	return "/api/v1/admin/user_number"
}

// AdminListUserAdminPath returns the URL path to the Admin service admin list user HTTP endpoint.
func AdminListUserAdminPath() string {
	return "/api/v1/admin/users"
}

// AdminGetUserAdminPath returns the URL path to the Admin service admin get user HTTP endpoint.
func AdminGetUserAdminPath(id string) string {
	return fmt.Sprintf("/api/v1/admin/users/%v", id)
}

// AdminCreateUserAdminPath returns the URL path to the Admin service admin create user HTTP endpoint.
func AdminCreateUserAdminPath() string {
	return "/api/v1/admin/users"
}

// AdminUpdateUserAdminPath returns the URL path to the Admin service admin update user HTTP endpoint.
func AdminUpdateUserAdminPath(id string) string {
	return fmt.Sprintf("/api/v1/admin/users/%v", id)
}

// AdminDeleteUserAdminPath returns the URL path to the Admin service admin delete user HTTP endpoint.
func AdminDeleteUserAdminPath(id string) string {
	return fmt.Sprintf("/api/v1/admin/users/%v", id)
}
