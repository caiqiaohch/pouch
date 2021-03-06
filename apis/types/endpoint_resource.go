// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EndpointResource NetworkResource is the body of the "get network" http response message
// swagger:model EndpointResource
type EndpointResource struct {

	// EndpointID represents the endpoint's id
	EndpointID string `json:"EndpointID,omitempty"`

	// IPv4Address represents the enpoint's ipv4 address
	IPV4Address string `json:"IPv4Address,omitempty"`

	// IPv4Address represents the enpoint's ipv6 address
	IPV6Address string `json:"IPv6Address,omitempty"`

	// MacAddress represents the enpoint's mac address
	MacAddress string `json:"MacAddress,omitempty"`

	// Name is the requested name of the network
	Name string `json:"Name,omitempty"`
}

// Validate validates this endpoint resource
func (m *EndpointResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EndpointResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointResource) UnmarshalBinary(b []byte) error {
	var res EndpointResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
