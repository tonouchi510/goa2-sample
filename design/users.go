package design

import (
	. "goa.design/goa/dsl"
)

// サービスの定義
var _ = Service("users", func() {
	Description("users serves user relative information.")

	HTTP(func() {
		Path("/api/v1/users")
	})

	Method("list", func() {
		Description("List all stored users")
		Result(CollectionOf(StoredUser))
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

	Method("show", func() {
		Description("Show user by ID")
		Payload(func() {
			Attribute("id", Int64, "ID of user to show")
			Required("id")
		})
		Result(StoredUser)
		Error("not_found", NotFound, "user not found")
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("add", func() {
		Description("Add new user and return its ID.")
		Payload(User)
		Result(Int64)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("update", func() {
		Description("Update user item.")
		Payload(func() {
			Attribute("id", Int64, "ID of user to show")
			Attribute("name", String, "Name of user")
			Attribute("email", String, "Email of user")
			Required("id")
		})
		Result(StoredUser)
		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
	})

	Method("remove", func() {
		Description("Remove user from storage")
		Payload(func() {
			Attribute("id", Int64, "ID of user to remove")
			Required("id")
		})
		Error("not_found", NotFound, "user not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})
})

// レスポンスデータの定義
var StoredUser = ResultType("application/vnd.goa2-sample.stored-user", func() {
	Description("A StoredUser describes a user access by the storage service.")
	Reference(User)
	TypeName("StoredUser")

	Attributes(func() {
		Attribute("id", Int64, "ID is the unique id of the user.", func() {
			Example(123)
		})
		Attribute("name")
		Attribute("email")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("email")
	})

	Required("id", "name", "email")
})

var User = Type("User", func() {
	Description("User describes a user account information to be stored.")
	Attribute("name", String, "Name of user", func() {
		MaxLength(30)
		Example("hoge fuga")
	})
	Attribute("email", String, "Email of user")

	Required("name", "email")
})

var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when attempting to show or delete a user that does not exist.")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("user 1 not found")
	})
	Attribute("id", Int64, "ID of missing user")
	Required("message", "id")
})