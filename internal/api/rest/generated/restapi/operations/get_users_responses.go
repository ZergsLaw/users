// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/zergslaw/boilerplate/internal/api/rest/generated/models"
)

// GetUsersOKCode is the HTTP code returned for type GetUsersOK
const GetUsersOKCode int = 200

/*GetUsersOK OK

swagger:response getUsersOK
*/
type GetUsersOK struct {

	/*
	  In: Body
	*/
	Payload *GetUsersOKBody `json:"body,omitempty"`
}

// NewGetUsersOK creates GetUsersOK with default headers values
func NewGetUsersOK() *GetUsersOK {

	return &GetUsersOK{}
}

// WithPayload adds the payload to the get users o k response
func (o *GetUsersOK) WithPayload(payload *GetUsersOKBody) *GetUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users o k response
func (o *GetUsersOK) SetPayload(payload *GetUsersOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetUsersDefault Generic error response.

swagger:response getUsersDefault
*/
type GetUsersDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUsersDefault creates GetUsersDefault with default headers values
func NewGetUsersDefault(code int) *GetUsersDefault {
	if code <= 0 {
		code = 500
	}

	return &GetUsersDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get users default response
func (o *GetUsersDefault) WithStatusCode(code int) *GetUsersDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get users default response
func (o *GetUsersDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get users default response
func (o *GetUsersDefault) WithPayload(payload *models.Error) *GetUsersDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users default response
func (o *GetUsersDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
