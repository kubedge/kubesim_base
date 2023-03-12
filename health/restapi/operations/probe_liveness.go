// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ProbeLivenessHandlerFunc turns a function with the right signature into a probe liveness handler
type ProbeLivenessHandlerFunc func(ProbeLivenessParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ProbeLivenessHandlerFunc) Handle(params ProbeLivenessParams) middleware.Responder {
	return fn(params)
}

// ProbeLivenessHandler interface for that can handle valid probe liveness params
type ProbeLivenessHandler interface {
	Handle(ProbeLivenessParams) middleware.Responder
}

// NewProbeLiveness creates a new http.Handler for the probe liveness operation
func NewProbeLiveness(ctx *middleware.Context, handler ProbeLivenessHandler) *ProbeLiveness {
	return &ProbeLiveness{Context: ctx, Handler: handler}
}

/*
ProbeLiveness swagger:route GET /liveness probeLiveness

ProbeLiveness probe liveness API
*/
type ProbeLiveness struct {
	Context *middleware.Context
	Handler ProbeLivenessHandler
}

func (o *ProbeLiveness) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewProbeLivenessParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
