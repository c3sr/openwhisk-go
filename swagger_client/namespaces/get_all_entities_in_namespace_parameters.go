package namespaces

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

// NewGetAllEntitiesInNamespaceParams creates a new GetAllEntitiesInNamespaceParams object
// with the default values initialized.
func NewGetAllEntitiesInNamespaceParams() *GetAllEntitiesInNamespaceParams {
	var ()
	return &GetAllEntitiesInNamespaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllEntitiesInNamespaceParamsWithTimeout creates a new GetAllEntitiesInNamespaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllEntitiesInNamespaceParamsWithTimeout(timeout time.Duration) *GetAllEntitiesInNamespaceParams {
	var ()
	return &GetAllEntitiesInNamespaceParams{

		timeout: timeout,
	}
}

// NewGetAllEntitiesInNamespaceParamsWithContext creates a new GetAllEntitiesInNamespaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllEntitiesInNamespaceParamsWithContext(ctx context.Context) *GetAllEntitiesInNamespaceParams {
	var ()
	return &GetAllEntitiesInNamespaceParams{

		Context: ctx,
	}
}

/*GetAllEntitiesInNamespaceParams contains all the parameters to send to the API endpoint
for the get all entities in namespace operation typically these are written to a http.Request
*/
type GetAllEntitiesInNamespaceParams struct {

	/*Namespace
	  The namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all entities in namespace params
func (o *GetAllEntitiesInNamespaceParams) WithTimeout(timeout time.Duration) *GetAllEntitiesInNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all entities in namespace params
func (o *GetAllEntitiesInNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all entities in namespace params
func (o *GetAllEntitiesInNamespaceParams) WithContext(ctx context.Context) *GetAllEntitiesInNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all entities in namespace params
func (o *GetAllEntitiesInNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithNamespace adds the namespace to the get all entities in namespace params
func (o *GetAllEntitiesInNamespaceParams) WithNamespace(namespace string) *GetAllEntitiesInNamespaceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get all entities in namespace params
func (o *GetAllEntitiesInNamespaceParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllEntitiesInNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
