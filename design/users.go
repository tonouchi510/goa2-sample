package design

import (
	. "goa.design/goa/dsl"
)

// サービスの定義
var _ = Service("Users", func() {
	Description("users serves user account relative information.")

	HTTP(func() {
		Path("/api/v1/users")
	})

	Method("List", func() {
		Description("List all stored users")

		Result(CollectionOf(UserResponse))
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

	Method("Get", func() {
		Description("Show user by ID")

		Payload(GetUserPayload)
		Result(UserResponse)

		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("Create", func() {
		Description("Add new user and return its ID.")

		Payload(CreateUserPayload)
		Result(String)

		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("Update", func() {
		Description("Update user item.")
		Payload(UpdateUserPayload)
		Result(UserResponse)
		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
	})

	Method("Delete", func() {
		Description("Delete user by id.")
		Payload(DeleteUserPayload)
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})
})
