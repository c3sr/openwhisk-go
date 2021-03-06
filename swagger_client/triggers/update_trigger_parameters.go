package triggers

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

	"github.com/c3sr/openwhisk-go/swagger_models"
)

// NewUpdateTriggerParams creates a new UpdateTriggerParams object
// with the default values initialized.
func NewUpdateTriggerParams() *UpdateTriggerParams {
	var ()
	return &UpdateTriggerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTriggerParamsWithTimeout creates a new UpdateTriggerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateTriggerParamsWithTimeout(timeout time.Duration) *UpdateTriggerParams {
	var ()
	return &UpdateTriggerParams{

		timeout: timeout,
	}
}

// NewUpdateTriggerParamsWithContext creates a new UpdateTriggerParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateTriggerParamsWithContext(ctx context.Context) *UpdateTriggerParams {
	var ()
	return &UpdateTriggerParams{

		Context: ctx,
	}
}

/*UpdateTriggerParams contains all the parameters to send to the API endpoint
for the update trigger operation typically these are written to a http.Request
*/
type UpdateTriggerParams struct {

	/*Namespace
	  The entity namespace

	*/
	Namespace string
	/*Overwrite
	  Overwrite item if it exists. Default is false.

	*/
	Overwrite *string
	/*Trigger
	  The trigger being updated

	*/
	Trigger *swagger_models.TriggerPut
	/*TriggerName
	  Name of trigger to update

	*/
	TriggerName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update trigger params
func (o *UpdateTriggerParams) WithTimeout(timeout time.Duration) *UpdateTriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update trigger params
func (o *UpdateTriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update trigger params
func (o *UpdateTriggerParams) WithContext(ctx context.Context) *UpdateTriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update trigger params
func (o *UpdateTriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithNamespace adds the namespace to the update trigger params
func (o *UpdateTriggerParams) WithNamespace(namespace string) *UpdateTriggerParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update trigger params
func (o *UpdateTriggerParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOverwrite adds the overwrite to the update trigger params
func (o *UpdateTriggerParams) WithOverwrite(overwrite *string) *UpdateTriggerParams {
	o.SetOverwrite(overwrite)
	return o
}

// SetOverwrite adds the overwrite to the update trigger params
func (o *UpdateTriggerParams) SetOverwrite(overwrite *string) {
	o.Overwrite = overwrite
}

// WithTrigger adds the trigger to the update trigger params
func (o *UpdateTriggerParams) WithTrigger(trigger *swagger_models.TriggerPut) *UpdateTriggerParams {
	o.SetTrigger(trigger)
	return o
}

// SetTrigger adds the trigger to the update trigger params
func (o *UpdateTriggerParams) SetTrigger(trigger *swagger_models.TriggerPut) {
	o.Trigger = trigger
}

// WithTriggerName adds the triggerName to the update trigger params
func (o *UpdateTriggerParams) WithTriggerName(triggerName string) *UpdateTriggerParams {
	o.SetTriggerName(triggerName)
	return o
}

// SetTriggerName adds the triggerName to the update trigger params
func (o *UpdateTriggerParams) SetTriggerName(triggerName string) {
	o.TriggerName = triggerName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Overwrite != nil {

		// query param overwrite
		var qrOverwrite string
		if o.Overwrite != nil {
			qrOverwrite = *o.Overwrite
		}
		qOverwrite := qrOverwrite
		if qOverwrite != "" {
			if err := r.SetQueryParam("overwrite", qOverwrite); err != nil {
				return err
			}
		}

	}

	if o.Trigger == nil {
		o.Trigger = new(swagger_models.TriggerPut)
	}

	if err := r.SetBodyParam(o.Trigger); err != nil {
		return err
	}

	// path param triggerName
	if err := r.SetPathParam("triggerName", o.TriggerName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
