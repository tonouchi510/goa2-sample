package controller

import (
	"context"
	"log"

	viron "github.com/tonouchi510/goa2-sample/gen/viron"
)

// Viron service example implementation.
// The example methods log the requests and return zero values.
type vironsrvc struct {
	logger *log.Logger
}

// NewViron returns the Viron service implementation.
func NewViron(logger *log.Logger) viron.Service {
	return &vironsrvc{logger}
}

// Add viron_authtype
func (s *vironsrvc) Authtype(ctx context.Context) (res viron.VironAuthtypeCollection, err error) {
	s.logger.Print("viron.authtype")
	res = viron.VironAuthtypeCollection{
		&viron.VironAuthtype{
			Method:   "POST",
			Provider: "",
			Type:     "jwt",
			URL:      "/signin",
		},
		&viron.VironAuthtype{
			Method:   "POST",
			Provider: "",
			Type:     "signout",
			URL:      "/signout",
		},
	}
	return
}

// Add viron_menu
func (s *vironsrvc) VironMenu(ctx context.Context) (res *viron.VironMenu, err error) {
	res = &viron.VironMenu{}
	s.logger.Print("viron.viron_menu")
	cl := "green"
	th := "standard"
	pk := "id"
	pagenation := true

	res = &viron.VironMenu{
		Name: "goa2-sample Admin Screen",
		Tags: []string{
			"local",
		},
		Color: &cl,
		Theme: &th,
		Pages: []*viron.VironPage{
			&viron.VironPage{
				Section:    "dashboard",
				Name:       "ダッシュボード",
				ID:         "quickview",
				Components: []*viron.VironComponent{
					&viron.VironComponent{
						Name:    "Users(bar)",
						API: &viron.VironAPI{
							Method: "get",
							Path:   "/api/v1/admin/user_number",
						},
						Style: "graph-bar",
					},
				},
			},
			&viron.VironPage{
				Section: "manage",
				ID:      "user-admin",
				Name:    "ユーザ管理",
				Components: []*viron.VironComponent{
					&viron.VironComponent{
						API: &viron.VironAPI{
							Method: "get",
							Path:   "/api/v1/admin/users",
						},
						Name:       "ユーザ一覧",
						Style:      "table",
						Primary:    &pk,
						Pagination: &pagenation,
						Query: []*viron.VironQuery{
							&viron.VironQuery{
								Key:  "id",
								Type: "integer",
							},
							&viron.VironQuery{
								Key:  "name",
								Type: "string",
							},
							&viron.VironQuery{
								Key:  "email",
								Type: "string",
							},
							&viron.VironQuery{
								Key:  "created_at",
								Type: "string",
							},
							&viron.VironQuery{
								Key:  "updated_at",
								Type: "string",
							},
						},
						TableLabels: []string{
							"id",
							"name",
							"email",
							"created_at",
							"updated_at",
						},
					},
				},
			},
		},
	}
	return
}
