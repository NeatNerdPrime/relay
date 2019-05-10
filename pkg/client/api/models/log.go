// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Log A log of a workflow run or workflow run action
// swagger:model Log
type Log struct {

	// Boolean indicating if the log is complete or not. If not complete, clients should re-request a socket connection
	Complete bool `json:"complete,omitempty"`

	// Log content. If the log is incomplete this should be excluded or nulled
	Content string `json:"content,omitempty"`
}

// Validate validates this log
func (m *Log) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Log) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Log) UnmarshalBinary(b []byte) error {
	var res Log
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
