// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generated command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/zergslaw/boilerplate/internal/api/web/generated/models"
)

// UpdateEmailReader is a Reader for the UpdateEmail structure.
type UpdateEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateEmailNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateEmailDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateEmailNoContent creates a UpdateEmailNoContent with default headers values
func NewUpdateEmailNoContent() *UpdateEmailNoContent {
	return &UpdateEmailNoContent{}
}

/*UpdateEmailNoContent handles this case with default header values.

The server successfully processed the request and is not returning any content.
*/
type UpdateEmailNoContent struct {
}

func (o *UpdateEmailNoContent) Error() string {
	return fmt.Sprintf("[PATCH /user/email][%d] updateEmailNoContent ", 204)
}

func (o *UpdateEmailNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEmailDefault creates a UpdateEmailDefault with default headers values
func NewUpdateEmailDefault(code int) *UpdateEmailDefault {
	return &UpdateEmailDefault{
		_statusCode: code,
	}
}

/*UpdateEmailDefault handles this case with default header values.

Generic error response.
*/
type UpdateEmailDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update email default response
func (o *UpdateEmailDefault) Code() int {
	return o._statusCode
}

func (o *UpdateEmailDefault) Error() string {
	return fmt.Sprintf("[PATCH /user/email][%d] updateEmail default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateEmailDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateEmailDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateEmailBody update email body
swagger:model UpdateEmailBody
*/
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