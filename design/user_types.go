package design

import (
	. "goa.design/goa/dsl"
)

// Users
var GetUserPayload = Type("GetUserPayload", func() {
	Attribute("id", String)
	Required("id")
})

var CreateUserPayload = Type("CreateUserPayload", func() {
	Attribute("id", String, "ID is the unique id of the user.", UserIDDefinition)
	Attribute("name", String, "Name of user")
	Attribute("email", String, "Email of user")
	Required("id", "name", "email")
})

var UpdateUserPayload = Type("UpdateUserPayload", func() {
	Attribute("id", String, "ID is the unique id of the user.", UserIDDefinition)
	Attribute("name", String, "Name of user")
	Attribute("email", String, "Email of user")
	Required("id")
})

var DeleteUserPayload = Type("DeleteUserPayload", func() {
	Attribute("id", String)
	Required("id")
})

var User = Type("User", func() {
	Description("User describes a user account information.")

	Attribute("id", String, "ID is the unique id of the user.", UserIDDefinition)
	Attribute("name", String, "Name of user", func() {
		Example("yamada taro")
	})
	Attribute("email", String, "Email of user", func() {
		Example("taro.yamada@xxx.jp")
	})

	Required("id", "name", "email")
})

var UserIDDefinition = func() {
	Description("User id")
	Example("XRQ85mtXnINISH25zfM0m5RlC6L2")
	MinLength(28)
	MaxLength(28)
}


// Admin
var Data = Type("data", func() {
	Attribute("key", String)
	Attribute("value", Any)
})

var GuideType = Type("StatsGuideType", func() {
	Attribute("x", LabelType)
	Attribute("y", LabelType)
})

var LabelType = Type("StatsLabelType", func() {
	Attribute("label", String)
	Required("label")
})
