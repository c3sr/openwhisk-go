package activations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new activations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for activations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetNamespacesNamespaceActivationsActivationidLogs gets the logs for an activation

Get activation logs information.
*/
func (a *Client) GetNamespacesNamespaceActivationsActivationidLogs(params *GetNamespacesNamespaceActivationsActivationidLogsParams) (*GetNamespacesNamespaceActivationsActivationidLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNamespacesNamespaceActivationsActivationidLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNamespacesNamespaceActivationsActivationidLogs",
		Method:             "GET",
		PathPattern:        "/namespaces/{namespace}/activations/{activationid}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNamespacesNamespaceActivationsActivationidLogsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNamespacesNamespaceActivationsActivationidLogsOK), nil

}

/*
GetNamespacesNamespaceActivationsActivationidResult gets the result of an activation

Get activation result.
*/
func (a *Client) GetNamespacesNamespaceActivationsActivationidResult(params *GetNamespacesNamespaceActivationsActivationidResultParams) (*GetNamespacesNamespaceActivationsActivationidResultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNamespacesNamespaceActivationsActivationidResultParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNamespacesNamespaceActivationsActivationidResult",
		Method:             "GET",
		PathPattern:        "/namespaces/{namespace}/activations/{activationid}/result",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNamespacesNamespaceActivationsActivationidResultReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNamespacesNamespaceActivationsActivationidResultOK), nil

}

/*
GetActivationByID gets activation information

Get activation information.
*/
func (a *Client) GetActivationByID(params *GetActivationByIDParams) (*GetActivationByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActivationByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getActivationById",
		Method:             "GET",
		PathPattern:        "/namespaces/{namespace}/activations/{activationid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetActivationByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetActivationByIDOK), nil

}

/*
GetActivations gets activation ids

Get activation ids.
*/
func (a *Client) GetActivations(params *GetActivationsParams) (*GetActivationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActivationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getActivations",
		Method:             "GET",
		PathPattern:        "/namespaces/{namespace}/activations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetActivationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetActivationsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
