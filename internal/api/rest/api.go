// Package rest contains all methods and middleware for working server.
package rest

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"path"
	"time"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sebest/xff"
	"github.com/sirupsen/logrus"
	"github.com/zergslaw/users/internal/api/rest/generated/restapi"
	"github.com/zergslaw/users/internal/api/rest/generated/restapi/operations"
	"github.com/zergslaw/users/internal/app"
)

// Log field names.
const (
	LogHost       = "host"
	LogPort       = "port"
	LogAddr       = "addr"
	LogRemote     = "remote" // aligned IPv4:port "   192.168.0.42:1234 "
	LogFunc       = "func"   // RPC method name, REST resource path
	LogHTTPMethod = "httpMethod"
	LogError      = "error"
	LogHTTPStatus = "httpStatus"
	LogUser       = "userID"
	LogAPI        = "api"
	LogVersion    = "version"
)

type (
	service struct {
		app app.App
	}

	config struct {
		host     string
		port     int
		basePath string
	}
	// Option for run server.
	Option func(*config)
)

// SetBasePath sets the base path to handlers.
// Default: /api/v1.
func SetBasePath(basePath string) Option {
	return func(c *config) {
		c.basePath = basePath
	}
}

// SetPort sets server port.
// Default: 8080.
func SetPort(port int) Option {
	return func(c *config) {
		c.port = port
	}
}

// SetHost sets server host.
// Default: localhost.
func SetHost(host string) Option {
	return func(c *config) {
		c.host = host
	}
}

func defaultConfig() *config {
	return &config{
		host:     "localhost",
		port:     8080,
		basePath: "",
	}
}

// New returns Swagger server configured to listen on the TCP network.
func New(application app.App, options ...Option) (*restapi.Server, error) {
	svc := &service{app: application}
	cfg := defaultConfig()

	for i := range options {
		options[i](cfg)
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, fmt.Errorf("load embedded swagger spec: %w", err)
	}
	if cfg.basePath == "" {
		cfg.basePath = swaggerSpec.BasePath()
	}
	swaggerSpec.Spec().BasePath = cfg.basePath
	api := operations.NewServiceUserAPI(swaggerSpec)
	api.Logger = logrus.New().WithField(LogAPI, "rest").Printf
	api.CookieKeyAuth = svc.cookieKeyAuth

	api.VerificationEmailHandler = operations.VerificationEmailHandlerFunc(svc.verificationEmail)
	api.VerificationUsernameHandler = operations.VerificationUsernameHandlerFunc(svc.verificationUsername)
	api.CreateUserHandler = operations.CreateUserHandlerFunc(svc.createUser)
	api.LoginHandler = operations.LoginHandlerFunc(svc.Login)
	api.LogoutHandler = operations.LogoutHandlerFunc(svc.logout)
	api.GetUserHandler = operations.GetUserHandlerFunc(svc.getUser)
	api.DeleteUserHandler = operations.DeleteUserHandlerFunc(svc.deleteUser)
	api.UpdatePasswordHandler = operations.UpdatePasswordHandlerFunc(svc.updatePassword)
	api.UpdateUsernameHandler = operations.UpdateUsernameHandlerFunc(svc.updateUsername)
	api.UpdateEmailHandler = operations.UpdateEmailHandlerFunc(svc.updateEmail)
	api.GetUsersHandler = operations.GetUsersHandlerFunc(svc.getUsers)

	server := restapi.NewServer(api)
	server.Host = cfg.host
	server.Port = cfg.port

	// The middlewareFunc executes before anything.
	globalMiddlewares := func(handler http.Handler) http.Handler {
		xffmw, _ := xff.Default()
		logger := logger(cfg.basePath)
		accesslog := accessLog(cfg.basePath)
		redocOpts := middleware.RedocOpts{
			BasePath: cfg.basePath,
			SpecURL:  path.Join(cfg.basePath, "/swagger.json"),
		}
		return xffmw.Handler(logger(recovery(accesslog(
			middleware.Spec(cfg.basePath, restapi.FlatSwaggerJSON,
				middleware.Redoc(redocOpts,
					handler))))))
	}

	server.SetHandler(globalMiddlewares(api.Serve(nil)))

	return server, nil
}

const authTimeout = 250 * time.Millisecond

func (svc *service) cookieKeyAuth(authToken string) (*app.AuthUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()
	profile, err := svc.app.UserByAuthToken(ctx, app.AuthToken(authToken))
	switch {
	case err != nil:
		return nil, fmt.Errorf("userByAuthToken: %w", err)
	default:
		return profile, nil
	}
}

func fromRequest(r *http.Request, authUser *app.AuthUser) (context.Context, logrus.FieldLogger, string) {
	ctx := r.Context()
	userID := app.UserID(0)
	if authUser != nil {
		userID = authUser.ID
	}
	log := logFromCtx(ctx).WithField(LogUser, userID)
	remoteIP, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ctx, log, remoteIP
}

// logFromCtx retrieves the current logger from the context. If no logger is
// available, the default logger is returned.
func logFromCtx(ctx context.Context) logrus.FieldLogger {
	val := ctx.Value(logKey)

	log, ok := val.(logrus.FieldLogger)
	if ok {
		return log
	}

	return logrus.New()
}