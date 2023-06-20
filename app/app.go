package app

/*
App package Loading and managing layer of systems and modules
this package use Gorrila package for handling endpoints
important loading item are
 - Database
 - Services
 - Models
 - Handlers

*/
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	hStripe "github.com/hamed-amini-dev/stripe-go-scripts/internal/handler/stripe"
	"github.com/hamed-amini-dev/stripe-go-scripts/internal/routes"
	"github.com/hamed-amini-dev/stripe-go-scripts/pkg/constants"
	sStripe "github.com/hamed-amini-dev/stripe-go-scripts/services/stripe"
	"github.com/spf13/viper"
)

type app struct {
	router *mux.Router
}

/* var (
	onceInitConfig = sync.Once{}
) */

// NewApp - Loads the configuration file for setting up the server
// and returns the app server
func NewApp() (server *http.Server, err error) {

	// ─────────────────────────────────────────────────────── APP INITIALIZATION ─────
	newApp := new(app)

	//Initialize Services for business logic of system layer
	serviceIStripe, err := sStripe.New(sStripe.InitOptionStripeKey(viper.GetString(constants.StripeKey)))
	if err != nil {
		return nil, err
	}

	//Initialize Handler layer for prepare endpoints
	handlerStripe, err := hStripe.New(hStripe.InitOptionService(serviceIStripe))
	if err != nil {
		return nil, err
	}
	//Initialize all router handle for listen on route address
	var allRoutes []routes.Route

	StripeRoutes := routes.StripeRoutes(handlerStripe)
	allRoutes = append(allRoutes, StripeRoutes...)

	// create resource for running app
	err = newApp.createResources(allRoutes...)
	if err != nil {
		return nil, err
	}

	return newApp.server(), nil
}

// server- Returns pointer to an instance of type http.Server
func (a *app) server() *http.Server {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString(constants.Port)),
		Handler: a.router,
	}

	log.Println("System is ready to transfer")
	log.Println("listen on port:", server.Addr)

	return server
}

// createResources - Creates router instance and registers handler to the router
func (a *app) createResources(rs ...routes.Route) error {
	a.router = mux.NewRouter().StrictSlash(true)

	for _, r := range rs {
		err := a.router.
			Name(r.Name).
			Path(r.Path).
			Methods(r.Method, http.MethodOptions).
			HandlerFunc(r.Handler).GetError()
		if err != nil {
			return err
		}
	}

	return nil
}
