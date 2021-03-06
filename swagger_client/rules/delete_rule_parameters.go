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
)

// NewDeleteRuleParams creates a new DeleteRuleParams object
// with the default values initialized.
func NewDeleteRuleParams() *DeleteRuleParams {
	var ()
	return &DeleteRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRuleParamsWithTimeout creates a new DeleteRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRuleParamsWithTimeout(timeout time.Duration) *DeleteRuleParams {
	var ()
	return &DeleteRuleParams{

		timeout: timeout,
	}
}

// NewDeleteRuleParamsWithContext creates a new DeleteRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRuleParamsWithContext(ctx context.Context) *DeleteRuleParams {
	var ()
	return &DeleteRuleParams{

		Context: ctx,
	}
}

/*DeleteRuleParams contains all the parameters to send to the API endpoint
for the delete rule operation typically these are written to a http.Request
*/
type DeleteRuleParams struct {

	/*Namespace
	  The entity namespace

	*/
	Namespace string
	/*RuleName
	  Name of rule to delete

	*/
	RuleName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete rule params
func (o *DeleteRuleParams) WithTimeout(timeout time.Duration) *DeleteRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete rule params
func (o *DeleteRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete rule params
func (o *DeleteRuleParams) WithContext(ctx context.Context) *DeleteRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete rule params
func (o *DeleteRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithNamespace adds the namespace to the delete rule params
func (o *DeleteRuleParams) WithNamespace(namespace string) *DeleteRuleParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete rule params
func (o *DeleteRuleParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithRuleName adds the ruleName to the delete rule params
func (o *DeleteRuleParams) WithRuleName(ruleName string) *DeleteRuleParams {
	o.SetRuleName(ruleName)
	return o
}

// SetRuleName adds the ruleName to the delete rule params
func (o *DeleteRuleParams) SetRuleName(ruleName string) {
	o.RuleName = ruleName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
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
