package rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new rules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for rules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteRule deletes a rule

Delete a rule
*/
func (a *Client) DeleteRule(params *DeleteRuleParams) (*DeleteRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRule",
		Method:             "DELETE",
		PathPattern:        "/namespaces/{namespace}/rules/{ruleName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteRuleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRuleOK), nil

}

/*
GetAllRules gets all rules

Get all rules
*/
func (a *Client) GetAllRules(params *GetAllRulesParams) (*GetAllRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllRules",
		Method:             "GET",
		PathPattern:        "/namespaces/{namespace}/rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllRulesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllRulesOK), nil

}

/*
GetRuleByName gets rule information

Get rule information
*/
func (a *Client) GetRuleByName(params *GetRuleByNameParams) (*GetRuleByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRuleByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRuleByName",
		Method:             "GET",
		PathPattern:        "/namespaces/{namespace}/rules/{ruleName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRuleByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRuleByNameOK), nil

}

/*
SetState enables or disable a rule

Enable or disable a rule
*/
func (a *Client) SetState(params *SetStateParams) (*SetStateOK, *SetStateAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setState",
		Method:             "POST",
		PathPattern:        "/namespaces/{namespace}/rules/{ruleName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetStateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *SetStateOK:
		return value, nil, nil
	case *SetStateAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
UpdateRule updates a rule

Update a rule
*/
func (a *Client) UpdateRule(params *UpdateRuleParams) (*UpdateRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateRule",
		Method:             "PUT",
		PathPattern:        "/namespaces/{namespace}/rules/{ruleName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateRuleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateRuleOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}