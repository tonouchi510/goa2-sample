package design

import (
	. "goa.design/goa/dsl"
	cors "goa.design/plugins/cors/dsl"
)

// API describes the global properties of the API server.
var _ = API("goa2-sample", func() {
	Title("SampleAPI")
	Description("goa2 sample code.")
	Version("1.0")
	Contact(func() {
		Name("tonouchi510")
		Email("tonouchi27@gmail.com")
		URL("https://github.com/tonouchi510/goa2-sample/issues")
	})
	Docs(func() {
		Description("wiki")
		URL("https://github.com/tonouchi510/goa2-sample/wiki")
	})

	cors.Origin("/.*localhost.*/", func() {
		cors.Headers("authorization, content-type")
		cors.Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
		cors.MaxAge(600)
	})

	Server("goa2-sample", func() {
		Services( "users", "swagger")
		Host("localhost", func() {
			Description("development host")
			URI("http://localhost:8000")
		})
	})
})
