package actions

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

// NewGetAllActionsParams creates a new GetAllActionsParams object
// with the default values initialized.
func NewGetAllActionsParams() *GetAllActionsParams {
	var ()
	return &GetAllActionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllActionsParamsWithTimeout creates a new GetAllActionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllActionsParamsWithTimeout(timeout time.Duration) *GetAllActionsParams {
	var ()
	return &GetAllActionsParams{

		timeout: timeout,
	}
}

// NewGetAllActionsParamsWithContext creates a new GetAllActionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllActionsParamsWithContext(ctx context.Context) *GetAllActionsParams {
	var ()
	return &GetAllActionsParams{

		Context: ctx,
	}
}

/*GetAllActionsParams contains all the parameters to send to the API endpoint
for the get all actions operation typically these are written to a http.Request
*/
type GetAllActionsParams struct {

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

// WithTimeout adds the timeout to the get all actions params
func (o *GetAllActionsParams) WithTimeout(timeout time.Duration) *GetAllActionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all actions params
func (o *GetAllActionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all actions params
func (o *GetAllActionsParams) WithContext(ctx context.Context) *GetAllActionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all actions params
func (o *GetAllActionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithLimit adds the limit to the get all actions params
func (o *GetAllActionsParams) WithLimit(limit *int64) *GetAllActionsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get all actions params
func (o *GetAllActionsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the get all actions params
func (o *GetAllActionsParams) WithNamespace(namespace string) *GetAllActionsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get all actions params
func (o *GetAllActionsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithSkip adds the skip to the get all actions params
func (o *GetAllActionsParams) WithSkip(skip *int64) *GetAllActionsParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the get all actions params
func (o *GetAllActionsParams) SetSkip(skip *int64) {
	o.Skip = skip
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllActionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
