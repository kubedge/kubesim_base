// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ProbeReadinessHandlerFunc turns a function with the right signature into a probe readiness handler
type ProbeReadinessHandlerFunc func(ProbeReadinessParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ProbeReadinessHandlerFunc) Handle(params ProbeReadinessParams) middleware.Responder {
	return fn(params)
}

// ProbeReadinessHandler interface for that can handle valid probe readiness params
type ProbeReadinessHandler interface {
	Handle(ProbeReadinessParams) middleware.Responder
}

// NewProbeReadiness creates a new http.Handler for the probe readiness operation
func NewProbeReadiness(ctx *middleware.Context, handler ProbeReadinessHandler) *ProbeReadiness {
	return &ProbeReadiness{Context: ctx, Handler: handler}
}

/*
ProbeReadiness swagger:route GET /readiness probeReadiness

ProbeReadiness probe readiness API
*/
type ProbeReadiness struct {
	Context *middleware.Context
	Handler ProbeReadinessHandler
}

func (o *ProbeReadiness) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewProbeReadinessParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
