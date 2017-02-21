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

	"github.com/c3sr/openwhisk-go/models"
)

// NewInvokeActionParams creates a new InvokeActionParams object
// with the default values initialized.
func NewInvokeActionParams() *InvokeActionParams {
	var ()
	return &InvokeActionParams{

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewInvokeActionParamsWithTimeout creates a new InvokeActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInvokeActionParamsWithTimeout(timeout time.Duration) *InvokeActionParams {
	var ()
	return &InvokeActionParams{

		requestTimeout: timeout,
	}
}

// NewInvokeActionParamsWithContext creates a new InvokeActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewInvokeActionParamsWithContext(ctx context.Context) *InvokeActionParams {
	var ()
	return &InvokeActionParams{

		Context: ctx,
	}
}

/*InvokeActionParams contains all the parameters to send to the API endpoint
for the invoke action operation typically these are written to a http.Request
*/
type InvokeActionParams struct {

	/*ActionName
	  Name of action

	*/
	ActionName string
	/*Blocking
	  Blocking or non-blocking invocation. Default is non-blocking.

	*/
	Blocking *string
	/*Namespace
	  The entity namespace

	*/
	Namespace string
	/*Payload
	  The parameters for the action being invoked

	*/
	Payload *models.KeyValue
	/*Result
	  Return only the result of a blocking activation. Default is false.

	*/
	Result *string
	/*Timeout
	  Wait no more than specified duration in milliseconds for a blocking response. Default value and max allowed timeout are 60000.

	*/
	Timeout *int64

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the invoke action params
func (o *InvokeActionParams) WithRequestTimeout(timeout time.Duration) *InvokeActionParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the invoke action params
func (o *InvokeActionParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the invoke action params
func (o *InvokeActionParams) WithContext(ctx context.Context) *InvokeActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invoke action params
func (o *InvokeActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithActionName adds the actionName to the invoke action params
func (o *InvokeActionParams) WithActionName(actionName string) *InvokeActionParams {
	o.SetActionName(actionName)
	return o
}

// SetActionName adds the actionName to the invoke action params
func (o *InvokeActionParams) SetActionName(actionName string) {
	o.ActionName = actionName
}

// WithBlocking adds the blocking to the invoke action params
func (o *InvokeActionParams) WithBlocking(blocking *string) *InvokeActionParams {
	o.SetBlocking(blocking)
	return o
}

// SetBlocking adds the blocking to the invoke action params
func (o *InvokeActionParams) SetBlocking(blocking *string) {
	o.Blocking = blocking
}

// WithNamespace adds the namespace to the invoke action params
func (o *InvokeActionParams) WithNamespace(namespace string) *InvokeActionParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the invoke action params
func (o *InvokeActionParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPayload adds the payload to the invoke action params
func (o *InvokeActionParams) WithPayload(payload *models.KeyValue) *InvokeActionParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the invoke action params
func (o *InvokeActionParams) SetPayload(payload *models.KeyValue) {
	o.Payload = payload
}

// WithResult adds the result to the invoke action params
func (o *InvokeActionParams) WithResult(result *string) *InvokeActionParams {
	o.SetResult(result)
	return o
}

// SetResult adds the result to the invoke action params
func (o *InvokeActionParams) SetResult(result *string) {
	o.Result = result
}

// WithTimeout adds the timeout to the invoke action params
func (o *InvokeActionParams) WithTimeout(timeout *int64) *InvokeActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invoke action params
func (o *InvokeActionParams) SetTimeout(timeout *int64) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *InvokeActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.requestTimeout)
	var res []error

	// path param actionName
	if err := r.SetPathParam("actionName", o.ActionName); err != nil {
		return err
	}

	if o.Blocking != nil {

		// query param blocking
		var qrBlocking string
		if o.Blocking != nil {
			qrBlocking = *o.Blocking
		}
		qBlocking := qrBlocking
		if qBlocking != "" {
			if err := r.SetQueryParam("blocking", qBlocking); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Payload == nil {
		o.Payload = new(models.KeyValue)
	}

	if err := r.SetBodyParam(o.Payload); err != nil {
		return err
	}

	if o.Result != nil {

		// query param result
		var qrResult string
		if o.Result != nil {
			qrResult = *o.Result
		}
		qResult := qrResult
		if qResult != "" {
			if err := r.SetQueryParam("result", qResult); err != nil {
				return err
			}
		}

	}

	if o.Timeout != nil {

		// query param timeout
		var qrTimeout int64
		if o.Timeout != nil {
			qrTimeout = *o.Timeout
		}
		qTimeout := swag.FormatInt64(qrTimeout)
		if qTimeout != "" {
			if err := r.SetQueryParam("timeout", qTimeout); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}