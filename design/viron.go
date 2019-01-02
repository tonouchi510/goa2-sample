package design

import (
	. "goa.design/goa/dsl"
)

var _ = Service("viron", func() {
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

	Files("/js/*filepath", "../../submodule/viron/public/js/")
})

// NumberType media type of status
var NumberType = ResultType("vnd.application/number+json", func() {
	Attributes(func() {
		Required(
			"value",
		)
		Attribute("value", Int32)
	})
	View("default", func() {
		Attribute("value")
	})
})

// AuthType AuthType for viron
var AuthType = ResultType("vnd.application/viron_authtype+json", func() {
	Attributes(func() {
		Required(
			"type",
			"provider",
			"url",
			"method",
		)
		Attribute("type", String, "type name")
		Attribute("provider", String, "provider name")
		Attribute("url", String, "url")
		Attribute("method", String, "request method to submit this auth")
	})
	View("default", func() {
		Attribute("type")
		Attribute("provider")
		Attribute("url")
		Attribute("method")
	})
})

// MenuType menu information on /viron
var MenuType = ResultType("vnd.application/viron_menu+json", func() {
	Attributes(func() {
		Required(
			"name",
			"pages",
		)
		Attribute("name", String)
		Attribute("thumbnail", String)
		Attribute("tags", ArrayOf(String))
		Attribute("color", String, func() {
			Enum("purple", "blue", "green", "yellow", "red", "gray", "black", "white")
		})
		Attribute("theme", String, func() {
			Enum("standard", "midnight", "terminal")
		})
		Attribute("pages", ArrayOf(PageType))
		Attribute("sections", ArrayOf(SectionType))
	})
	View("default", func() {
		Attribute("name")
		Attribute("thumbnail")
		Attribute("tags")
		Attribute("color")
		Attribute("theme")
		Attribute("pages")
	})
})

// SectionType component for section in viron
var SectionType = ResultType("vnd.application/viron_section+json", func() {
	Attributes(func() {
		Required(
			"id",
			"label",
		)
		Attribute("id", String)
		Attribute("label", String)
	})
	View("default", func() {
		Attribute("id")
		Attribute("label")
	})
})

// PageType pagetype media
var PageType = ResultType("vnd.application/viron_page+json", func() {
	Attributes(func() {
		Required(
			"id",
			"name",
			"section",
			"components",
		)
		Attribute("id", String)
		Attribute("name", String)
		Attribute("section", String)
		Attribute("group", String)
		Attribute("components", ArrayOf(ComponentType))
	})
	View("default", func() {
		Attribute("section")
		Attribute("group")
		Attribute("id")
		Attribute("name")
		Attribute("components")
	})
})

// ComponentType media type for component in viron
var ComponentType = ResultType("vnd.application/viron_component+json", func() {
	Attributes(func() {
		Required(
			"name",
			"style",
			"api",
		)
		Attribute("name", String)
		Attribute("style", String, func() {
			Enum("number",
				"table",
				"graph-bar",
				"graph-scatterplot",
				"graph-line",
				"graph-horizontal-bar",
				"graph-stacked-bar",
				"graph-horizontal-stacked-bar",
				"graph-stacked-area",
			)
		})
		Attribute("unit", String)
		Attribute("actions", ArrayOf(String))
		Attribute("api", APIType)
		Attribute("pagination", Boolean)
		Attribute("primary", String)
		Attribute("table_labels", ArrayOf(String))
		Attribute("query", ArrayOf(QueryType))
		Attribute("auto_refresh_sec", Int32)
	})
	View("default", func() {
		Attribute("name")
		Attribute("style")
		Attribute("unit")
		Attribute("actions")
		Attribute("api")
		Attribute("pagination")
		Attribute("primary")
		Attribute("table_labels")
		Attribute("query")
		Attribute("auto_refresh_sec")
	})
})

// APIType media type for api in viron
var APIType = ResultType("vnd.application/viron_api+json", func() {
	Attributes(func() {
		Required(
			"method",
			"path",
		)
		Attribute("method", String)
		Attribute("path", String)
	})
	View("default", func() {
		Attribute("method")
		Attribute("path")
	})
})

// QueryType mediatype for query in viron
var QueryType = ResultType("vnd.application/viron_query+json", func() {
	Attributes(func() {
		Required(
			"key",
			"type",
		)
		Attribute("key", String)
		Attribute("type", String)
	})
	View("default", func() {
		Attribute("key")
		Attribute("type")
	})
})