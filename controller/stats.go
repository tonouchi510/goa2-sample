package controller

import (
	"context"
	"log"

	stats "github.com/tonouchi510/goa2-sample/gen/stats"
)

// stats service example implementation.
// The example methods log the requests and return zero values.
type statsSvc struct {
	logger *log.Logger
}

// NewStats returns the stats service implementation.
func NewStats(logger *log.Logger) stats.Service {
	return &statsSvc{logger}
}

// Users Information
func (s *statsSvc) UserNumber(ctx context.Context) (res *stats.Statsuser, err error) {
	res = &stats.Statsuser{}
	s.logger.Print("stats.user_number")

	st1 := "2018/11"
	st2 := "2018/12"
	st3 := "2019/01"
	c := "key"
	size := "value"

	// StatsPlanetController_Bar: end_implement
	res = &stats.Statsuser{
		Data: []*stats.Data{
			&stats.Data{Key:&st1, Value:1},
			&stats.Data{Key:&st2, Value:2},
			&stats.Data{Key:&st3, Value:5},
		},
		X: "key",
		Y: "value",
		Color: &c,
		Size: &size,
		Guide: &stats.StatsGuideType{
			X: &stats.StatsLabelType{Label: "年月"},
			Y: &stats.StatsLabelType{Label: "人数"},
		},
	}
	return
}
