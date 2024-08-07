package handlers

import (
	"dev/register.git/business/register"
	"dev/register.git/business/i"
	"dev/register.git/business/mid"
	"dev/register.git/foundation/web"
	"net/http"
	"os"
)

type Register struct {
	Service *register.Service
}

// API constructs a http.Handler with all application routes defined
func API(log i.Logger, r Register, shutdown chan os.Signal) *web.App {

	// Create web app with middleware
	app := web.NewApp(
		shutdown,
		mid.Logger(log),
		mid.Errors(log),
		mid.Panics(log),
	)

	// Check Service
	ch := check{}
	app.Handle(http.MethodGet, "/readiness", ch.readiness)
	app.Handle(http.MethodGet, "/liveliness", ch.liveliness)

	// Register Handlers
	app.Handle(http.MethodPost, "/create", r.create)
	return app

}

// Init will initialise the Service
func Init(db register.Store, log i.Logger) Register {

	// Initialise services
	r := Register{
		Service: &register.Service{
			Log:   log,
			Store: db,
		},
	}
	return r

}
