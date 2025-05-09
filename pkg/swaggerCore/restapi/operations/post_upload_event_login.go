// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostUploadEventLoginHandlerFunc turns a function with the right signature into a post upload event login handler
type PostUploadEventLoginHandlerFunc func(PostUploadEventLoginParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostUploadEventLoginHandlerFunc) Handle(params PostUploadEventLoginParams) middleware.Responder {
	return fn(params)
}

// PostUploadEventLoginHandler interface for that can handle valid post upload event login params
type PostUploadEventLoginHandler interface {
	Handle(PostUploadEventLoginParams) middleware.Responder
}

// NewPostUploadEventLogin creates a new http.Handler for the post upload event login operation
func NewPostUploadEventLogin(ctx *middleware.Context, handler PostUploadEventLoginHandler) *PostUploadEventLogin {
	return &PostUploadEventLogin{Context: ctx, Handler: handler}
}

/*
	PostUploadEventLogin swagger:route POST /upload/eventLogin postUploadEventLogin

PostUploadEventLogin post upload event login API
*/
type PostUploadEventLogin struct {
	Context *middleware.Context
	Handler PostUploadEventLoginHandler
}

func (o *PostUploadEventLogin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostUploadEventLoginParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
