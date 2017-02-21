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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/c3sr/openwhisk-go/models"
)

// NewUpdateRuleParams creates a new UpdateRuleParams object
// with the default values initialized.
func NewUpdateRuleParams() *UpdateRuleParams {
	var ()
	return &UpdateRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRuleParamsWithTimeout creates a new UpdateRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateRuleParamsWithTimeout(timeout time.Duration) *UpdateRuleParams {
	var ()
	return &UpdateRuleParams{

		timeout: timeout,
	}
}

// NewUpdateRuleParamsWithContext creates a new UpdateRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateRuleParamsWithContext(ctx context.Context) *UpdateRuleParams {
	var ()
	return &UpdateRuleParams{

		Context: ctx,
	}
}

/*UpdateRuleParams contains all the parameters to send to the API endpoint
for the update rule operation typically these are written to a http.Request
*/
type UpdateRuleParams struct {

	/*Namespace
	  The entity namespace

	*/
	Namespace string
	/*Overwrite
	  Overwrite item if it exists. Default is false.

	*/
	Overwrite *string
	/*Rule
	  The rule being updated

	*/
	Rule *models.RulePut
	/*RuleName
	  Name of rule to update

	*/
	RuleName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update rule params
func (o *UpdateRuleParams) WithTimeout(timeout time.Duration) *UpdateRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update rule params
func (o *UpdateRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update rule params
func (o *UpdateRuleParams) WithContext(ctx context.Context) *UpdateRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update rule params
func (o *UpdateRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithNamespace adds the namespace to the update rule params
func (o *UpdateRuleParams) WithNamespace(namespace string) *UpdateRuleParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update rule params
func (o *UpdateRuleParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOverwrite adds the overwrite to the update rule params
func (o *UpdateRuleParams) WithOverwrite(overwrite *string) *UpdateRuleParams {
	o.SetOverwrite(overwrite)
	return o
}

// SetOverwrite adds the overwrite to the update rule params
func (o *UpdateRuleParams) SetOverwrite(overwrite *string) {
	o.Overwrite = overwrite
}

// WithRule adds the rule to the update rule params
func (o *UpdateRuleParams) WithRule(rule *models.RulePut) *UpdateRuleParams {
	o.SetRule(rule)
	return o
}

// SetRule adds the rule to the update rule params
func (o *UpdateRuleParams) SetRule(rule *models.RulePut) {
	o.Rule = rule
}

// WithRuleName adds the ruleName to the update rule params
func (o *UpdateRuleParams) WithRuleName(ruleName string) *UpdateRuleParams {
	o.SetRuleName(ruleName)
	return o
}

// SetRuleName adds the ruleName to the update rule params
func (o *UpdateRuleParams) SetRuleName(ruleName string) {
	o.RuleName = ruleName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Rule == nil {
		o.Rule = new(models.RulePut)
	}

	if err := r.SetBodyParam(o.Rule); err != nil {
		return err
	}

	// path param ruleName
	if err := r.SetPathParam("ruleName", o.RuleName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
