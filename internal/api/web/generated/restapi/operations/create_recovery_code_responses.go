// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/zergslaw/boilerplate/internal/api/web/generated/models"
)

// CreateRecoveryCodeNoContentCode is the HTTP code returned for type CreateRecoveryCodeNoContent
const CreateRecoveryCodeNoContentCode int = 204

/*CreateRecoveryCodeNoContent The server successfully processed the request and is not returning any content.

swagger:response createRecoveryCodeNoContent
*/
type CreateRecoveryCodeNoContent struct {
}

// NewCreateRecoveryCodeNoContent creates CreateRecoveryCodeNoContent with default headers values
func NewCreateRecoveryCodeNoContent() *CreateRecoveryCodeNoContent {

	return &CreateRecoveryCodeNoContent{}
}

// WriteResponse to the client
func (o *CreateRecoveryCodeNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*CreateRecoveryCodeDefault Generic error response.

swagger:response createRecoveryCodeDefault
*/
type CreateRecoveryCodeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateRecoveryCodeDefault creates CreateRecoveryCodeDefault with default headers values
func NewCreateRecoveryCodeDefault(code int) *CreateRecoveryCodeDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateRecoveryCodeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create recovery code default response
func (o *CreateRecoveryCodeDefault) WithStatusCode(code int) *CreateRecoveryCodeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create recovery code default response
func (o *CreateRecoveryCodeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create recovery code default response
func (o *CreateRecoveryCodeDefault) WithPayload(payload *models.Error) *CreateRecoveryCodeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create recovery code default response
func (o *CreateRecoveryCodeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRecoveryCodeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
