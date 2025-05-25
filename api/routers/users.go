package routers

import (
	"net/http"
	"web-crawler/api/routers/users"

	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
)

func RegisterUserRoutes(f *fizz.Fizz) {
	f.Generator().SetSecuritySchemes(map[string]*openapi.SecuritySchemeOrRef{
		"JsonWebToken": {
			SecurityScheme: &openapi.SecurityScheme{
				Type: "apiKey",
				In:   "header",
				Name: "Authorization",
			},
		},
	})

	g := f.Group("/api/users", "Users", "Collection of API endpoints to manage users")

	g.GET("/:id",
		[]fizz.OperationOption{
			fizz.Summary("get the user by id"),
			fizz.Security(&openapi.SecurityRequirement{"JsonWebToken": []string{}}),
		},
		tonic.Handler(users.GetUserById, http.StatusOK))
}
