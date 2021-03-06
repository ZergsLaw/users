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

	"github.com/zergslaw/boilerplate/internal/api/web/generated/models"
	"github.com/zergslaw/boilerplate/internal/app"
)

// UpdateEmailHandlerFunc turns a function with the right signature into a update email handler
type UpdateEmailHandlerFunc func(UpdateEmailParams, *app.AuthUser) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateEmailHandlerFunc) Handle(params UpdateEmailParams, principal *app.AuthUser) middleware.Responder {
	return fn(params, principal)
}

// UpdateEmailHandler interface for that can handle valid update email params
type UpdateEmailHandler interface {
	Handle(UpdateEmailParams, *app.AuthUser) middleware.Responder
}

// NewUpdateEmail creates a new http.Handler for the update email operation
func NewUpdateEmail(ctx *middleware.Context, handler UpdateEmailHandler) *UpdateEmail {
	return &UpdateEmail{Context: ctx, Handler: handler}
}

/*UpdateEmail swagger:route PATCH /user/email updateEmail

Change email.

*/
type UpdateEmail struct {
	Context *middleware.Context
	Handler UpdateEmailHandler
}

func (o *UpdateEmail) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateEmailParams()

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

// UpdateEmailBody update email body
//
// swagger:model UpdateEmailBody
type UpdateEmailBody struct {

	// email
	// Required: true
	// Format: email
	Email models.Email `json:"email"`
}

// Validate validates this update email body
func (o *UpdateEmailBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateEmailBody) validateEmail(formats strfmt.Registry) error {

	if err := o.Email.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("args" + "." + "email")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateEmailBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateEmailBody) UnmarshalBinary(b []byte) error {
	var res UpdateEmailBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
