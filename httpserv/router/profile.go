package router

import (
	"go-gcp-cloud-run/httpserv/handler"
	"go-gcp-cloud-run/services/greeting"
	"go-gcp-cloud-run/services/login"
	"net/http"

	"github.com/witwoywhy/go-cores/gins"
)

func BindGreetingRoute(app gins.GinApps) {
	svc := greeting.New()

	hdl := handler.NewGreetingHandler(svc)
	app.Register(
		http.MethodGet,
		"greeting",
		hdl.Handle,
	)
}

func BindLoginRoute(app gins.GinApps) {
	svc := login.New()

	hdl := handler.NewLoginHandler(svc)
	app.Register(
		http.MethodPost,
		"login",
		hdl.Handle,
	)
}
