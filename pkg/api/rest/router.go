// Package rest is port handler.
// # This manifest was generated by ymir. DO NOT EDIT.
package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	pkgRest "github.com/kubuskotak/asgard/rest"
	pkgTracer "github.com/kubuskotak/asgard/tracer"

	"github.com/abialemuel/ymirblog/pkg/infrastructure"
	"github.com/abialemuel/ymirblog/pkg/version"
)

// RegisterFunc is type func to register handler.
type RegisterFunc func(h chi.Router) http.Handler

// Router is the data struct.
type Router struct {
	h *chi.Mux
}

// Register will assign rest handler.
func (r *Router) Register(fn RegisterFunc) http.Handler {
	// set router middleware
	r.h.Use(pkgTracer.Middleware(
		infrastructure.Envs.App.ServiceName))
	r.h.Use(pkgRest.SemanticVersion(
		infrastructure.Envs.App.ServiceName,
		version.GetVersion().VersionNumber(),
	))
	r.h.NotFound(pkgRest.NotFoundDefault()) // Not Found Handler
	return fn(r.h)
}

// Routes create Router instance.
func Routes() *Router {
	r := &Router{
		h: chi.NewRouter(), // port http
	}
	return r
}
