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

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeletePackageParams creates a new DeletePackageParams object
// with the default values initialized.
func NewDeletePackageParams() *DeletePackageParams {
	var ()
	return &DeletePackageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePackageParamsWithTimeout creates a new DeletePackageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePackageParamsWithTimeout(timeout time.Duration) *DeletePackageParams {
	var ()
	return &DeletePackageParams{

		timeout: timeout,
	}
}

// NewDeletePackageParamsWithContext creates a new DeletePackageParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePackageParamsWithContext(ctx context.Context) *DeletePackageParams {
	var ()
	return &DeletePackageParams{

		Context: ctx,
	}
}

/*DeletePackageParams contains all the parameters to send to the API endpoint
for the delete package operation typically these are written to a http.Request
*/
type DeletePackageParams struct {

	/*Namespace
	  The entity namespace

	*/
	Namespace string
	/*PackageName
	  Name of package

	*/
	PackageName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete package params
func (o *DeletePackageParams) WithTimeout(timeout time.Duration) *DeletePackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete package params
func (o *DeletePackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete package params
func (o *DeletePackageParams) WithContext(ctx context.Context) *DeletePackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete package params
func (o *DeletePackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithNamespace adds the namespace to the delete package params
func (o *DeletePackageParams) WithNamespace(namespace string) *DeletePackageParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete package params
func (o *DeletePackageParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPackageName adds the packageName to the delete package params
func (o *DeletePackageParams) WithPackageName(packageName string) *DeletePackageParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the delete package params
func (o *DeletePackageParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param packageName
	if err := r.SetPathParam("packageName", o.PackageName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
