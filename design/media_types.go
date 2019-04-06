package design

import (
	. "goa.design/goa/dsl"
)

// レスポンスデータの定義
// MediaType of Users API.
var UserResponse = ResultType("application/vnd.goa2-sample.user+json", func() {
	Description("User Response")
	ContentType("application/json")

	Reference(User)
	Attributes(func() {
		Attribute("id")
		Attribute("name")
		Attribute("email")

		Required("id", "name", "email")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("email")
	})
})


// MediaType of Admin API.
var AdminUserStatsResponse = ResultType("application/vnd.goa2-sample.admin.user_number+json", func() {
	Description("statistic of users")
	ContentType("application/json")

	Attributes(func() {
		Attribute("data", ArrayOf(Data), "グラフデータ")
		Attribute("x", String, "X軸に使用するkey")
		Attribute("y", String, "Y軸に使用するkey")
		Attribute("size", String, "ドットの大きさに使用するkey")
		Attribute("color", String, "ドットの色分けに使用するkey")
		Attribute("guide", GuideType)

		Required("data", "x", "y", "guide")
	})
})

var AdminUserResponse = ResultType("application/vnd.goa2-sample.admin.user+json", func() {
	Description("User Response")
	ContentType("application/json")

	Reference(User)
	Attributes(func() {
		Attribute("id")
		Attribute("name")
		Attribute("email")
		Attribute("created_at", String)
		Attribute("updated_at", String)
		Attribute("deleted_at", String)

		Required("id", "name", "email", "created_at", "updated_at")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("email")
		Attribute("created_at")
		Attribute("updated_at")
		Attribute("deleted_at")
	})
})


// MediaType of Viron API
// for /viron_authtype
var AuthType = ResultType("application/vnd.viron.authtype+json", func() {
	Attributes(func() {

		Attribute("type", String, "type name")
		Attribute("provider", String, "provider name")
		Attribute("url", String, "url")
		Attribute("method", String, "request method to submit this auth")

		Required("type", "provider", "url", "method")
	})

	View("default", func() {
		Attribute("type")
		Attribute("provider")
		Attribute("url")
		Attribute("method")
	})
})

// MenuType menu information on /viron
var MenuType = ResultType("application/vnd.viron.menu+json", func() {
	Attributes(func() {
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

		Required("name", "pages")
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
var SectionType = ResultType("application/vnd.viron.section+json", func() {
	Attributes(func() {
		Attribute("id", String)
		Attribute("label", String)

		Required("id", "label")
	})
	View("default", func() {
		Attribute("id")
		Attribute("label")
	})
})

// PageType pagetype media
var PageType = ResultType("application/vnd.viron.page+json", func() {
	Attributes(func() {
		Attribute("id", String)
		Attribute("name", String)
		Attribute("section", String)
		Attribute("group", String)
		Attribute("components", ArrayOf(ComponentType))

		Required("id", "name", "section", "components")
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
var ComponentType = ResultType("application/vnd.viron.component+json", func() {
	Attributes(func() {
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

		Required("name", "style", "api")
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
var APIType = ResultType("application/vnd.viron.api+json", func() {
	Attributes(func() {
		Attribute("method", String)
		Attribute("path", String)

		Required("method", "path")
	})

	View("default", func() {
		Attribute("method")
		Attribute("path")
	})
})

// QueryType mediatype for query in viron
var QueryType = ResultType("application/vnd.viron.query+json", func() {
	Attributes(func() {
		Attribute("key", String)
		Attribute("type", String)

		Required("key", "type")
	})
	View("default", func() {
		Attribute("key")
		Attribute("type")
	})
})

