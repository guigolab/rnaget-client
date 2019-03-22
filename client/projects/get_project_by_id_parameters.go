// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetProjectByIDParams creates a new GetProjectByIDParams object
// with the default values initialized.
func NewGetProjectByIDParams() *GetProjectByIDParams {
	var ()
	return &GetProjectByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectByIDParamsWithTimeout creates a new GetProjectByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProjectByIDParamsWithTimeout(timeout time.Duration) *GetProjectByIDParams {
	var ()
	return &GetProjectByIDParams{

		timeout: timeout,
	}
}

// NewGetProjectByIDParamsWithContext creates a new GetProjectByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProjectByIDParamsWithContext(ctx context.Context) *GetProjectByIDParams {
	var ()
	return &GetProjectByIDParams{

		Context: ctx,
	}
}

// NewGetProjectByIDParamsWithHTTPClient creates a new GetProjectByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProjectByIDParamsWithHTTPClient(client *http.Client) *GetProjectByIDParams {
	var ()
	return &GetProjectByIDParams{
		HTTPClient: client,
	}
}

/*GetProjectByIDParams contains all the parameters to send to the API endpoint
for the get project by Id operation typically these are written to a http.Request
*/
type GetProjectByIDParams struct {

	/*ProjectID
	  ID of project to return

	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get project by Id params
func (o *GetProjectByIDParams) WithTimeout(timeout time.Duration) *GetProjectByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project by Id params
func (o *GetProjectByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project by Id params
func (o *GetProjectByIDParams) WithContext(ctx context.Context) *GetProjectByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project by Id params
func (o *GetProjectByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project by Id params
func (o *GetProjectByIDParams) WithHTTPClient(client *http.Client) *GetProjectByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project by Id params
func (o *GetProjectByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the get project by Id params
func (o *GetProjectByIDParams) WithProjectID(projectID string) *GetProjectByIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get project by Id params
func (o *GetProjectByIDParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
