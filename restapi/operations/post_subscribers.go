// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostSubscribersHandlerFunc turns a function with the right signature into a post subscribers handler
type PostSubscribersHandlerFunc func(PostSubscribersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostSubscribersHandlerFunc) Handle(params PostSubscribersParams) middleware.Responder {
	return fn(params)
}

// PostSubscribersHandler interface for that can handle valid post subscribers params
type PostSubscribersHandler interface {
	Handle(PostSubscribersParams) middleware.Responder
}

// NewPostSubscribers creates a new http.Handler for the post subscribers operation
func NewPostSubscribers(ctx *middleware.Context, handler PostSubscribersHandler) *PostSubscribers {
	return &PostSubscribers{Context: ctx, Handler: handler}
}

/*
	PostSubscribers swagger:route POST /subscribers postSubscribers

# Register a new user

Endpoint to register a new user with name and email
*/
type PostSubscribers struct {
	Context *middleware.Context
	Handler PostSubscribersHandler
}

func (o *PostSubscribers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostSubscribersParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}