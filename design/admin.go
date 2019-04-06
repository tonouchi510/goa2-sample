package design

import (
	. "goa.design/goa/dsl"
)

var _ = Service("Admin", func() {
	Description("Admin provide functions for the management screen.")

	HTTP(func() {
		Path("/api/v1/admin")
	})

	Method("user_number", func() {
		Description("Number of users")

		Result(AdminUserStatsResponse)
		HTTP(func() {
			GET("/user_number")
			Response(StatusOK)
		})
	})

	Method("admin list user", func() {
		Description("List all stored users")

		Result(CollectionOf(AdminUserResponse))
		HTTP(func() {
			GET("/users")
			Response(StatusOK)
		})
	})

	Method("admin get user", func() {
		Description("Show user by ID")

		Payload(GetUserPayload)
		Result(AdminUserResponse)

		HTTP(func() {
			GET("/users/{id}")
			Response(StatusOK)
		})
	})

	Method("admin create user", func() {
		Description("Add new user and return its ID.")

		Payload(CreateUserPayload)
		Result(String)

		HTTP(func() {
			POST("/users")
			Response(StatusCreated)
		})
	})

	Method("admin update user", func() {
		Description("Update user item.")
		Payload(UpdateUserPayload)
		Result(AdminUserResponse)
		HTTP(func() {
			PUT("/users/{id}")
			Response(StatusOK)
		})
	})

	Method("admin delete user", func() {
		Description("Delete user by id.")
		Payload(DeleteUserPayload)
		HTTP(func() {
			DELETE("/users/{id}")
			Response(StatusNoContent)
		})
	})
})
