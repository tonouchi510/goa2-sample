package design

import (
	. "goa.design/goa/dsl"
)

var _ = Service("stats", func() {
	Description("Stats describes stats information of this services")

	HTTP(func() {
		Path("/api/v1/stats")
	})

	Method("user_number", func() {
		Description("Users Information")

		Result(StatsUserBar)
		HTTP(func() {
			GET("/user_number")
			Response(StatusOK)
		})
	})
})

// StatsPlanetMediaType of media type.
var StatsUserBar = ResultType("application/vnd.statsuser+json", func() {
	Description("Users data")
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

	View("default", func() {
		Attribute("data")
		Attribute("x")
		Attribute("y")
		Attribute("size")
		Attribute("color")
		Attribute("guide")
	})
})

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