package design

import (
	. "goa.design/goa/dsl"
)

var _ = Service("Viron", func() {
	Method("authtype", func() {
		Description("Add viron_authtype")
		Result(CollectionOf(AuthType))
		HTTP(func() {
			GET("/viron_authtype")
			Response(StatusOK)
		})
	})

	Method("viron_menu", func() {
		Description("Add viron_menu")
		Result(MenuType)
		HTTP(func() {
			GET("/viron")
			Response(StatusOK)
		})
	})

	Files("/v1/swagger.json", "gen/http/openapi.json", func() {
		Description("This endpoint serves the API swagger definition for viron.")
	})
})
