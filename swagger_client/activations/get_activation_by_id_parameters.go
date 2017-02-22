package activations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetActivationByIDParams creates a new GetActivationByIDParams object
// with the default values initialized.
func NewGetActivationByIDParams() *GetActivationByIDParams {
	var ()
	return &GetActivationByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetActivationByIDParamsWithTimeout creates a new GetActivationByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetActivationByIDParamsWithTimeout(timeout time.Duration) *GetActivationByIDParams {
	var ()
	return &GetActivationByIDParams{

		timeout: timeout,
	}
}

// NewGetActivationByIDParamsWithContext creates a new GetActivationByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetActivationByIDParamsWithContext(ctx context.Context) *GetActivationByIDParams {
	var ()
	return &GetActivationByIDParams{

		Context: ctx,
	}
}

/*GetActivationByIDParams contains all the parameters to send to the API endpoint
for the get activation by Id operation typically these are written to a http.Request
*/
type GetActivationByIDParams struct {

	/*Activationid
	  Name of activation to fetch

	*/
	Activationid string
	/*Namespace
	  The entity namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get activation by Id params
func (o *GetActivationByIDParams) WithTimeout(timeout time.Duration) *GetActivationByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get activation by Id params
func (o *GetActivationByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get activation by Id params
func (o *GetActivationByIDParams) WithContext(ctx context.Context) *GetActivationByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get activation by Id params
func (o *GetActivationByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithActivationid adds the activationid to the get activation by Id params
func (o *GetActivationByIDParams) WithActivationid(activationid string) *GetActivationByIDParams {
	o.SetActivationid(activationid)
	return o
}

// SetActivationid adds the activationid to the get activation by Id params
func (o *GetActivationByIDParams) SetActivationid(activationid string) {
	o.Activationid = activationid
}

// WithNamespace adds the namespace to the get activation by Id params
func (o *GetActivationByIDParams) WithNamespace(namespace string) *GetActivationByIDParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get activation by Id params
func (o *GetActivationByIDParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetActivationByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param activationid
	if err := r.SetPathParam("activationid", o.Activationid); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}