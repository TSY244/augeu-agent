// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"augeu-agent/pkg/swaggerCore/restapi/operations"
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

//go:generate swagger generate server --target ../../swaggerCore --name Augeu --spec ../../../../backEnd/swagger.yaml --principal models.Principle

func configureFlags(api *operations.AugeuAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.AugeuAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

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

	if api.GetGetClientsHandler == nil {
		api.GetGetClientsHandler = operations.GetGetClientsHandlerFunc(func(params operations.GetGetClientsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetGetClients has not yet been implemented")
		})
	}
	if api.GetVersionHandler == nil {
		api.GetVersionHandler = operations.GetVersionHandlerFunc(func(params operations.GetVersionParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetVersion has not yet been implemented")
		})
	}
	if api.PostGetApplicationEventHandler == nil {
		api.PostGetApplicationEventHandler = operations.PostGetApplicationEventHandlerFunc(func(params operations.PostGetApplicationEventParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostGetApplicationEvent has not yet been implemented")
		})
	}
	if api.PostGetClientIDHandler == nil {
		api.PostGetClientIDHandler = operations.PostGetClientIDHandlerFunc(func(params operations.PostGetClientIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostGetClientID has not yet been implemented")
		})
	}
	if api.PostGetLoginEventHandler == nil {
		api.PostGetLoginEventHandler = operations.PostGetLoginEventHandlerFunc(func(params operations.PostGetLoginEventParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostGetLoginEvent has not yet been implemented")
		})
	}
	if api.PostGetPowershellEventHandler == nil {
		api.PostGetPowershellEventHandler = operations.PostGetPowershellEventHandlerFunc(func(params operations.PostGetPowershellEventParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostGetPowershellEvent has not yet been implemented")
		})
	}
	if api.PostGetProcessEventHandler == nil {
		api.PostGetProcessEventHandler = operations.PostGetProcessEventHandlerFunc(func(params operations.PostGetProcessEventParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostGetProcessEvent has not yet been implemented")
		})
	}
	if api.PostGetRdpEventHandler == nil {
		api.PostGetRdpEventHandler = operations.PostGetRdpEventHandlerFunc(func(params operations.PostGetRdpEventParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostGetRdpEvent has not yet been implemented")
		})
	}
	if api.PostGetSecurityEventHandler == nil {
		api.PostGetSecurityEventHandler = operations.PostGetSecurityEventHandlerFunc(func(params operations.PostGetSecurityEventParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostGetSecurityEvent has not yet been implemented")
		})
	}
	if api.PostGetServiceEventHandler == nil {
		api.PostGetServiceEventHandler = operations.PostGetServiceEventHandlerFunc(func(params operations.PostGetServiceEventParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostGetServiceEvent has not yet been implemented")
		})
	}
	if api.PostGetSystemEventHandler == nil {
		api.PostGetSystemEventHandler = operations.PostGetSystemEventHandlerFunc(func(params operations.PostGetSystemEventParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostGetSystemEvent has not yet been implemented")
		})
	}
	if api.PostLoginHandler == nil {
		api.PostLoginHandler = operations.PostLoginHandlerFunc(func(params operations.PostLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostLogin has not yet been implemented")
		})
	}
	if api.PostRegisterHandler == nil {
		api.PostRegisterHandler = operations.PostRegisterHandlerFunc(func(params operations.PostRegisterParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostRegister has not yet been implemented")
		})
	}
	if api.PostUpdataApplicationEventHandler == nil {
		api.PostUpdataApplicationEventHandler = operations.PostUpdataApplicationEventHandlerFunc(func(params operations.PostUpdataApplicationEventParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostUpdataApplicationEvent has not yet been implemented")
		})
	}
	if api.PostUpdataPowershellEventHandler == nil {
		api.PostUpdataPowershellEventHandler = operations.PostUpdataPowershellEventHandlerFunc(func(params operations.PostUpdataPowershellEventParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostUpdataPowershellEvent has not yet been implemented")
		})
	}
	if api.PostUpdataProcessEventHandler == nil {
		api.PostUpdataProcessEventHandler = operations.PostUpdataProcessEventHandlerFunc(func(params operations.PostUpdataProcessEventParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostUpdataProcessEvent has not yet been implemented")
		})
	}
	if api.PostUpdataSecurityEventHandler == nil {
		api.PostUpdataSecurityEventHandler = operations.PostUpdataSecurityEventHandlerFunc(func(params operations.PostUpdataSecurityEventParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostUpdataSecurityEvent has not yet been implemented")
		})
	}
	if api.PostUpdataServiceEventHandler == nil {
		api.PostUpdataServiceEventHandler = operations.PostUpdataServiceEventHandlerFunc(func(params operations.PostUpdataServiceEventParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostUpdataServiceEvent has not yet been implemented")
		})
	}
	if api.PostUpdataSystemEventHandler == nil {
		api.PostUpdataSystemEventHandler = operations.PostUpdataSystemEventHandlerFunc(func(params operations.PostUpdataSystemEventParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostUpdataSystemEvent has not yet been implemented")
		})
	}
	if api.PostUploadLoginEventHandler == nil {
		api.PostUploadLoginEventHandler = operations.PostUploadLoginEventHandlerFunc(func(params operations.PostUploadLoginEventParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostUploadLoginEvent has not yet been implemented")
		})
	}
	if api.PostUploadRdpEventHandler == nil {
		api.PostUploadRdpEventHandler = operations.PostUploadRdpEventHandlerFunc(func(params operations.PostUploadRdpEventParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostUploadRdpEvent has not yet been implemented")
		})
	}
	if api.PostUploadUserInfoHandler == nil {
		api.PostUploadUserInfoHandler = operations.PostUploadUserInfoHandlerFunc(func(params operations.PostUploadUserInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostUploadUserInfo has not yet been implemented")
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
// If you need to modify a rule, store server instance to stop it individually later, this is the place.
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
