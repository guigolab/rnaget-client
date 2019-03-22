// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Study study
// swagger:model study
type Study struct {

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// parent project ID
	ParentProjectID string `json:"parentProjectID,omitempty"`

	// patient list
	PatientList []string `json:"patientList"`

	// sample list
	SampleList []string `json:"sampleList"`

	// tags
	Tags []string `json:"tags"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this study
func (m *Study) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Study) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Study) UnmarshalBinary(b []byte) error {
	var res Study
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
