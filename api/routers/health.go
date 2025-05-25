package routers

import (
	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
	"net/http"
	"web-crawler/api/routers/health"
)

func RegisterHealthRoutes(f *fizz.Fizz) {
	g := f.Group("/health", "Health", "Collection of endpoint to check health of the server")
	g.GET("", []fizz.OperationOption{fizz.Summary("Get the server status")}, tonic.Handler(health.CheckHealth, http.StatusOK))

}
