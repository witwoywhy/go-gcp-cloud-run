package httpserv

import (
	"go-gcp-cloud-run/httpserv/router"

	"github.com/witwoywhy/go-cores/apps"
	"github.com/witwoywhy/go-cores/gins"
)

func Run() {
	app := gins.New()
	app.UseMiddleware(gins.Log())

	router.BindGreetingRoute(app)
	router.BindLoginRoute(app)

	app.ListenAndServe(":" + apps.AppConfig.Port)
}
