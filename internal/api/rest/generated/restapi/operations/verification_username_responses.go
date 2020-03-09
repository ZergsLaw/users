// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/zergslaw/boilerplate/internal/api/rest/generated/models"
)

// VerificationUsernameNoContentCode is the HTTP code returned for type VerificationUsernameNoContent
const VerificationUsernameNoContentCode int = 204

/*VerificationUsernameNoContent The server successfully processed the request and is not returning any content.

swagger:response verificationUsernameNoContent
*/
type VerificationUsernameNoContent struct {
}

// NewVerificationUsernameNoContent creates VerificationUsernameNoContent with default headers values
func NewVerificationUsernameNoContent() *VerificationUsernameNoContent {

	return &VerificationUsernameNoContent{}
}

// WriteResponse to the client
func (o *VerificationUsernameNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*VerificationUsernameDefault Generic error response.

swagger:response verificationUsernameDefault
*/
type VerificationUsernameDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewVerificationUsernameDefault creates VerificationUsernameDefault with default headers values
func NewVerificationUsernameDefault(code int) *VerificationUsernameDefault {
	if code <= 0 {
		code = 500
	}

	return &VerificationUsernameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the verification username default response
func (o *VerificationUsernameDefault) WithStatusCode(code int) *VerificationUsernameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the verification username default response
func (o *VerificationUsernameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the verification username default response
func (o *VerificationUsernameDefault) WithPayload(payload *models.Error) *VerificationUsernameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the verification username default response
func (o *VerificationUsernameDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VerificationUsernameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
