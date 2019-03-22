// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// File URL and type for raw data and analysis pipeline files
// swagger:model file
type File struct {

	// URL
	URL string `json:"URL,omitempty"`

	// file type
	FileType string `json:"fileType,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// study ID
	StudyID string `json:"studyID,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this file
func (m *File) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *File) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *File) UnmarshalBinary(b []byte) error {
	var res File
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
