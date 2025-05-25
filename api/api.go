package api

import (
	"github.com/wI2L/fizz"
	"web-crawler/api/middlewares"
	"web-crawler/api/routers"
)

func RegisterRoutes(f *fizz.Fizz) {
	f.Use(middlewares.AuthMw)

	routers.RegisterHealthRoutes(f)
}
