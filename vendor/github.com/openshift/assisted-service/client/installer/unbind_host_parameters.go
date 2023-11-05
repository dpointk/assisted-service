// Code generated by go-swagger; DO NOT EDIT.

package installer

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
)

// NewUnbindHostParams creates a new UnbindHostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnbindHostParams() *UnbindHostParams {
	return &UnbindHostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnbindHostParamsWithTimeout creates a new UnbindHostParams object
// with the ability to set a timeout on a request.
func NewUnbindHostParamsWithTimeout(timeout time.Duration) *UnbindHostParams {
	return &UnbindHostParams{
		timeout: timeout,
	}
}

// NewUnbindHostParamsWithContext creates a new UnbindHostParams object
// with the ability to set a context for a request.
func NewUnbindHostParamsWithContext(ctx context.Context) *UnbindHostParams {
	return &UnbindHostParams{
		Context: ctx,
	}
}

// NewUnbindHostParamsWithHTTPClient creates a new UnbindHostParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnbindHostParamsWithHTTPClient(client *http.Client) *UnbindHostParams {
	return &UnbindHostParams{
		HTTPClient: client,
	}
}

/*
UnbindHostParams contains all the parameters to send to the API endpoint

	for the unbind host operation.

	Typically these are written to a http.Request.
*/
type UnbindHostParams struct {

	/* HostID.

	   The host that is being bound.

	   Format: uuid
	*/
	HostID strfmt.UUID

	/* InfraEnvID.

	   The infra-env of the host that is being bound.

	   Format: uuid
	*/
	InfraEnvID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unbind host params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnbindHostParams) WithDefaults() *UnbindHostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unbind host params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnbindHostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unbind host params
func (o *UnbindHostParams) WithTimeout(timeout time.Duration) *UnbindHostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unbind host params
func (o *UnbindHostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unbind host params
func (o *UnbindHostParams) WithContext(ctx context.Context) *UnbindHostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unbind host params
func (o *UnbindHostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unbind host params
func (o *UnbindHostParams) WithHTTPClient(client *http.Client) *UnbindHostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unbind host params
func (o *UnbindHostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHostID adds the hostID to the unbind host params
func (o *UnbindHostParams) WithHostID(hostID strfmt.UUID) *UnbindHostParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the unbind host params
func (o *UnbindHostParams) SetHostID(hostID strfmt.UUID) {
	o.HostID = hostID
}

// WithInfraEnvID adds the infraEnvID to the unbind host params
func (o *UnbindHostParams) WithInfraEnvID(infraEnvID strfmt.UUID) *UnbindHostParams {
	o.SetInfraEnvID(infraEnvID)
	return o
}

// SetInfraEnvID adds the infraEnvId to the unbind host params
func (o *UnbindHostParams) SetInfraEnvID(infraEnvID strfmt.UUID) {
	o.InfraEnvID = infraEnvID
}

// WriteToRequest writes these params to a swagger request
func (o *UnbindHostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param host_id
	if err := r.SetPathParam("host_id", o.HostID.String()); err != nil {
		return err
	}

	// path param infra_env_id
	if err := r.SetPathParam("infra_env_id", o.InfraEnvID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}