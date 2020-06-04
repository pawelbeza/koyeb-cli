// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new activity API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for activity API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAccountActivities(params *GetAccountActivitiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountActivitiesOK, error)

	ListActivities(params *ListActivitiesParams, authInfo runtime.ClientAuthInfoWriter) (*ListActivitiesOK, error)

	ListNotifications(params *ListNotificationsParams, authInfo runtime.ClientAuthInfoWriter) (*ListNotificationsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAccountActivities get account activities API
*/
func (a *Client) GetAccountActivities(params *GetAccountActivitiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountActivitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountActivitiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAccountActivities",
		Method:             "GET",
		PathPattern:        "/v1/account/activities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAccountActivitiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountActivitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAccountActivities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListActivities list activities API
*/
func (a *Client) ListActivities(params *ListActivitiesParams, authInfo runtime.ClientAuthInfoWriter) (*ListActivitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListActivitiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListActivities",
		Method:             "GET",
		PathPattern:        "/v1/activities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListActivitiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListActivitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListActivities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListNotifications list notifications API
*/
func (a *Client) ListNotifications(params *ListNotificationsParams, authInfo runtime.ClientAuthInfoWriter) (*ListNotificationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNotificationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListNotifications",
		Method:             "GET",
		PathPattern:        "/v1/notifications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListNotificationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListNotificationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListNotifications: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
