// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/atrariksa/fastrogos/rula/models"
)

// NewCreateUserParams creates a new CreateUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateUserParams() *CreateUserParams {
	return &CreateUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserParamsWithTimeout creates a new CreateUserParams object
// with the ability to set a timeout on a request.
func NewCreateUserParamsWithTimeout(timeout time.Duration) *CreateUserParams {
	return &CreateUserParams{
		timeout: timeout,
	}
}

// NewCreateUserParamsWithContext creates a new CreateUserParams object
// with the ability to set a context for a request.
func NewCreateUserParamsWithContext(ctx context.Context) *CreateUserParams {
	return &CreateUserParams{
		Context: ctx,
	}
}

// NewCreateUserParamsWithHTTPClient creates a new CreateUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateUserParamsWithHTTPClient(client *http.Client) *CreateUserParams {
	return &CreateUserParams{
		HTTPClient: client,
	}
}

/* CreateUserParams contains all the parameters to send to the API endpoint
   for the create user operation.

   Typically these are written to a http.Request.
*/
type CreateUserParams struct {

	/* ModelsCreateUserReq.

	   CreateUserReq
	*/
	ModelsCreateUserReq *models.ModelsCreateUserReq

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUserParams) WithDefaults() *CreateUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create user params
func (o *CreateUserParams) WithTimeout(timeout time.Duration) *CreateUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create user params
func (o *CreateUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create user params
func (o *CreateUserParams) WithContext(ctx context.Context) *CreateUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create user params
func (o *CreateUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create user params
func (o *CreateUserParams) WithHTTPClient(client *http.Client) *CreateUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create user params
func (o *CreateUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModelsCreateUserReq adds the modelsCreateUserReq to the create user params
func (o *CreateUserParams) WithModelsCreateUserReq(modelsCreateUserReq *models.ModelsCreateUserReq) *CreateUserParams {
	o.SetModelsCreateUserReq(modelsCreateUserReq)
	return o
}

// SetModelsCreateUserReq adds the modelsCreateUserReq to the create user params
func (o *CreateUserParams) SetModelsCreateUserReq(modelsCreateUserReq *models.ModelsCreateUserReq) {
	o.ModelsCreateUserReq = modelsCreateUserReq
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ModelsCreateUserReq != nil {
		if err := r.SetBodyParam(o.ModelsCreateUserReq); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
