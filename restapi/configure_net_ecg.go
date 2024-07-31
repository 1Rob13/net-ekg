// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/1Rob13/net-ekg/db"
	subscribers "github.com/1Rob13/net-ekg/handlers"
	"github.com/1Rob13/net-ekg/restapi/operations"
)

//go:generate swagger generate server --target ../../net-ekg --name NetEcg --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.NetEcgAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.NetEcgAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	//new db

	db := db.NewSQClient()

	subscriberHandler := subscribers.New(db)

	//new subscriber handler

	api.GetSubscribersHandler = operations.GetSubscribersHandlerFunc(subscriberHandler.HandleGet)
	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.GetSubscribersHandler == nil {
		api.GetSubscribersHandler = operations.GetSubscribersHandlerFunc(func(params operations.GetSubscribersParams) middleware.Responder {
			return middleware.NotImplemented("hello")
		})
	}
	if api.PostSubscribersHandler == nil {
		api.PostSubscribersHandler = operations.PostSubscribersHandlerFunc(func(params operations.PostSubscribersParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostSubscribers has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {

}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}