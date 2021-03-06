package packages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAlPackagesParams creates a new GetAlPackagesParams object
// with the default values initialized.
func NewGetAlPackagesParams() *GetAlPackagesParams {
	var ()
	return &GetAlPackagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlPackagesParamsWithTimeout creates a new GetAlPackagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAlPackagesParamsWithTimeout(timeout time.Duration) *GetAlPackagesParams {
	var ()
	return &GetAlPackagesParams{

		timeout: timeout,
	}
}

// NewGetAlPackagesParamsWithContext creates a new GetAlPackagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAlPackagesParamsWithContext(ctx context.Context) *GetAlPackagesParams {
	var ()
	return &GetAlPackagesParams{

		Context: ctx,
	}
}

/*GetAlPackagesParams contains all the parameters to send to the API endpoint
for the get al packages operation typically these are written to a http.Request
*/
type GetAlPackagesParams struct {

	/*Limit
	  Number of entities to include in the result.

	*/
	Limit *int64
	/*Namespace
	  The entity namespace

	*/
	Namespace string
	/*Public
	  Include publicly shared entitles in the result.

	*/
	Public *bool
	/*Skip
	  Number of entities to skip in the result.

	*/
	Skip *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get al packages params
func (o *GetAlPackagesParams) WithTimeout(timeout time.Duration) *GetAlPackagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get al packages params
func (o *GetAlPackagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get al packages params
func (o *GetAlPackagesParams) WithContext(ctx context.Context) *GetAlPackagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get al packages params
func (o *GetAlPackagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithLimit adds the limit to the get al packages params
func (o *GetAlPackagesParams) WithLimit(limit *int64) *GetAlPackagesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get al packages params
func (o *GetAlPackagesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the get al packages params
func (o *GetAlPackagesParams) WithNamespace(namespace string) *GetAlPackagesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get al packages params
func (o *GetAlPackagesParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPublic adds the public to the get al packages params
func (o *GetAlPackagesParams) WithPublic(public *bool) *GetAlPackagesParams {
	o.SetPublic(public)
	return o
}

// SetPublic adds the public to the get al packages params
func (o *GetAlPackagesParams) SetPublic(public *bool) {
	o.Public = public
}

// WithSkip adds the skip to the get al packages params
func (o *GetAlPackagesParams) WithSkip(skip *int64) *GetAlPackagesParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the get al packages params
func (o *GetAlPackagesParams) SetSkip(skip *int64) {
	o.Skip = skip
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlPackagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Public != nil {

		// query param public
		var qrPublic bool
		if o.Public != nil {
			qrPublic = *o.Public
		}
		qPublic := swag.FormatBool(qrPublic)
		if qPublic != "" {
			if err := r.SetQueryParam("public", qPublic); err != nil {
				return err
			}
		}

	}

	if o.Skip != nil {

		// query param skip
		var qrSkip int64
		if o.Skip != nil {
			qrSkip = *o.Skip
		}
		qSkip := swag.FormatInt64(qrSkip)
		if qSkip != "" {
			if err := r.SetQueryParam("skip", qSkip); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
