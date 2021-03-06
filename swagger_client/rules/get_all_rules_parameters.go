package rules

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

// NewGetAllRulesParams creates a new GetAllRulesParams object
// with the default values initialized.
func NewGetAllRulesParams() *GetAllRulesParams {
	var ()
	return &GetAllRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllRulesParamsWithTimeout creates a new GetAllRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllRulesParamsWithTimeout(timeout time.Duration) *GetAllRulesParams {
	var ()
	return &GetAllRulesParams{

		timeout: timeout,
	}
}

// NewGetAllRulesParamsWithContext creates a new GetAllRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllRulesParamsWithContext(ctx context.Context) *GetAllRulesParams {
	var ()
	return &GetAllRulesParams{

		Context: ctx,
	}
}

/*GetAllRulesParams contains all the parameters to send to the API endpoint
for the get all rules operation typically these are written to a http.Request
*/
type GetAllRulesParams struct {

	/*Limit
	  Number of entities to include in the result.

	*/
	Limit *int64
	/*Namespace
	  The entity namespace

	*/
	Namespace string
	/*Skip
	  Number of entities to skip in the result.

	*/
	Skip *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all rules params
func (o *GetAllRulesParams) WithTimeout(timeout time.Duration) *GetAllRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all rules params
func (o *GetAllRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all rules params
func (o *GetAllRulesParams) WithContext(ctx context.Context) *GetAllRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all rules params
func (o *GetAllRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithLimit adds the limit to the get all rules params
func (o *GetAllRulesParams) WithLimit(limit *int64) *GetAllRulesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get all rules params
func (o *GetAllRulesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the get all rules params
func (o *GetAllRulesParams) WithNamespace(namespace string) *GetAllRulesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get all rules params
func (o *GetAllRulesParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithSkip adds the skip to the get all rules params
func (o *GetAllRulesParams) WithSkip(skip *int64) *GetAllRulesParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the get all rules params
func (o *GetAllRulesParams) SetSkip(skip *int64) {
	o.Skip = skip
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
