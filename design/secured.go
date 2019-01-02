package design

import (
	. "goa.design/goa/dsl"
)


var _ = Service("secured", func() {
	Description("The secured service exposes endpoints that require valid authorization credentials.")

	Method("signin", func() {
		Description("Creates a valid JWT")

		// The signin endpoint is secured via basic auth
		Security(BasicAuth)

		Payload(func() {
			Description("Credentials used to authenticate to retrieve JWT token")
			Username("username", String, "Username used to perform signin", func() {
				Example("user")
			})
			Password("password", String, "Password used to perform signin", func() {
				Example("password")
			})
			Required("username", "password")
		})
		Result(JWT)
		Error("unauthorized", String, "Credentials are invalid")

		HTTP(func() {
			POST("/login")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
		})

	})
})

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint. Supports scopes "api:read" and "api:write".`)
	Scope("api:read", "Read-only access")
	Scope("api:write", "Read and write access")
})

// BasicAuth defines a security scheme using basic authentication. The scheme
// protects the "signin" action used to create JWTs.
var BasicAuth = BasicAuthSecurity("basic", func() {
	Description("Basic authentication used to authenticate security principal during signin")
})

var JWT = ResultType("application/vnd.goa.jwt", func() {
	Description("JWT describes access token of api.")
	Attribute("Authorization", String, "New Jwt token")
	Required("Authorization")
})
