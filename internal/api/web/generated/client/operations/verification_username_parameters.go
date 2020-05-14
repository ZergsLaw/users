// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generated command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewVerificationUsernameParams creates a new VerificationUsernameParams object
// with the default values initialized.
func NewVerificationUsernameParams() *VerificationUsernameParams {
	var ()
	return &VerificationUsernameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVerificationUsernameParamsWithTimeout creates a new VerificationUsernameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVerificationUsernameParamsWithTimeout(timeout time.Duration) *VerificationUsernameParams {
	var ()
	return &VerificationUsernameParams{

		timeout: timeout,
	}
}

// NewVerificationUsernameParamsWithContext creates a new VerificationUsernameParams object
// with the default values initialized, and the ability to set a context for a request
func NewVerificationUsernameParamsWithContext(ctx context.Context) *VerificationUsernameParams {
	var ()
	return &VerificationUsernameParams{

		Context: ctx,
	}
}

// NewVerificationUsernameParamsWithHTTPClient creates a new VerificationUsernameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVerificationUsernameParamsWithHTTPClient(client *http.Client) *VerificationUsernameParams {
	var ()
	return &VerificationUsernameParams{
		HTTPClient: client,
	}
}

/*VerificationUsernameParams contains all the parameters to send to the API endpoint
for the verification username operation typically these are written to a http.Request
*/
type VerificationUsernameParams struct {

	/*Args*/
	Args VerificationUsernameBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the verification username params
func (o *VerificationUsernameParams) WithTimeout(timeout time.Duration) *VerificationUsernameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the verification username params
func (o *VerificationUsernameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the verification username params
func (o *VerificationUsernameParams) WithContext(ctx context.Context) *VerificationUsernameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the verification username params
func (o *VerificationUsernameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the verification username params
func (o *VerificationUsernameParams) WithHTTPClient(client *http.Client) *VerificationUsernameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the verification username params
func (o *VerificationUsernameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgs adds the args to the verification username params
func (o *VerificationUsernameParams) WithArgs(args VerificationUsernameBody) *VerificationUsernameParams {
	o.SetArgs(args)
	return o
}

// SetArgs adds the args to the verification username params
func (o *VerificationUsernameParams) SetArgs(args VerificationUsernameBody) {
	o.Args = args
}

// WriteToRequest writes these params to a swagger request
func (o *VerificationUsernameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Args); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}