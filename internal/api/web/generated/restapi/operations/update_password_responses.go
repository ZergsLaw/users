// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/zergslaw/boilerplate/internal/api/web/generated/models"
)

// UpdatePasswordNoContentCode is the HTTP code returned for type UpdatePasswordNoContent
const UpdatePasswordNoContentCode int = 204

/*UpdatePasswordNoContent The server successfully processed the request and is not returning any content.

swagger:response updatePasswordNoContent
*/
type UpdatePasswordNoContent struct {
}

// NewUpdatePasswordNoContent creates UpdatePasswordNoContent with default headers values
func NewUpdatePasswordNoContent() *UpdatePasswordNoContent {

	return &UpdatePasswordNoContent{}
}

// WriteResponse to the client
func (o *UpdatePasswordNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*UpdatePasswordDefault Generic error response.

swagger:response updatePasswordDefault
*/
type UpdatePasswordDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdatePasswordDefault creates UpdatePasswordDefault with default headers values
func NewUpdatePasswordDefault(code int) *UpdatePasswordDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdatePasswordDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update password default response
func (o *UpdatePasswordDefault) WithStatusCode(code int) *UpdatePasswordDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update password default response
func (o *UpdatePasswordDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update password default response
func (o *UpdatePasswordDefault) WithPayload(payload *models.Error) *UpdatePasswordDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update password default response
func (o *UpdatePasswordDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePasswordDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
