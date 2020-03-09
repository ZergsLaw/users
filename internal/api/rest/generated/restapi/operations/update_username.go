// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/zergslaw/boilerplate/internal/api/rest/generated/models"
	"github.com/zergslaw/boilerplate/internal/app"
)

// UpdateUsernameHandlerFunc turns a function with the right signature into a update username handler
type UpdateUsernameHandlerFunc func(UpdateUsernameParams, *app.AuthUser) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateUsernameHandlerFunc) Handle(params UpdateUsernameParams, principal *app.AuthUser) middleware.Responder {
	return fn(params, principal)
}

// UpdateUsernameHandler interface for that can handle valid update username params
type UpdateUsernameHandler interface {
	Handle(UpdateUsernameParams, *app.AuthUser) middleware.Responder
}

// NewUpdateUsername creates a new http.Handler for the update username operation
func NewUpdateUsername(ctx *middleware.Context, handler UpdateUsernameHandler) *UpdateUsername {
	return &UpdateUsername{Context: ctx, Handler: handler}
}

/*UpdateUsername swagger:route PATCH /user/username updateUsername

Change username.

*/
type UpdateUsername struct {
	Context *middleware.Context
	Handler UpdateUsernameHandler
}

func (o *UpdateUsername) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateUsernameParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *app.AuthUser
	if uprinc != nil {
		principal = uprinc.(*app.AuthUser) // this is really a app.AuthUser, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UpdateUsernameBody update username body
//
// swagger:model UpdateUsernameBody
type UpdateUsernameBody struct {

	// username
	// Required: true
	Username models.Username `json:"username"`
}

// Validate validates this update username body
func (o *UpdateUsernameBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateUsernameBody) validateUsername(formats strfmt.Registry) error {

	if err := o.Username.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("args" + "." + "username")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateUsernameBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateUsernameBody) UnmarshalBinary(b []byte) error {
	var res UpdateUsernameBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
