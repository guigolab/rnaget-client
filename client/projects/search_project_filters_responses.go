// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/guigolab/rnaget-client/models"
)

// SearchProjectFiltersReader is a Reader for the SearchProjectFilters structure.
type SearchProjectFiltersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchProjectFiltersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchProjectFiltersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSearchProjectFiltersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchProjectFiltersOK creates a SearchProjectFiltersOK with default headers values
func NewSearchProjectFiltersOK() *SearchProjectFiltersOK {
	return &SearchProjectFiltersOK{}
}

/*SearchProjectFiltersOK handles this case with default header values.

successful operation
*/
type SearchProjectFiltersOK struct {
	Payload []*models.SearchFilter
}

func (o *SearchProjectFiltersOK) Error() string {
	return fmt.Sprintf("[GET /projects/search/filters][%d] searchProjectFiltersOK  %+v", 200, o.Payload)
}

func (o *SearchProjectFiltersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchProjectFiltersBadRequest creates a SearchProjectFiltersBadRequest with default headers values
func NewSearchProjectFiltersBadRequest() *SearchProjectFiltersBadRequest {
	return &SearchProjectFiltersBadRequest{}
}

/*SearchProjectFiltersBadRequest handles this case with default header values.

Error
*/
type SearchProjectFiltersBadRequest struct {
}

func (o *SearchProjectFiltersBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/search/filters][%d] searchProjectFiltersBadRequest ", 400)
}

func (o *SearchProjectFiltersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
