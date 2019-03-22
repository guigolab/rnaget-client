// Code generated by go-swagger; DO NOT EDIT.

package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/guigolab/rnaget-client/models"
)

// SearchFilesReader is a Reader for the SearchFiles structure.
type SearchFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSearchFilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchFilesOK creates a SearchFilesOK with default headers values
func NewSearchFilesOK() *SearchFilesOK {
	return &SearchFilesOK{}
}

/*SearchFilesOK handles this case with default header values.

Successful operation
*/
type SearchFilesOK struct {
	Payload []*models.File
}

func (o *SearchFilesOK) Error() string {
	return fmt.Sprintf("[GET /files/search][%d] searchFilesOK  %+v", 200, o.Payload)
}

func (o *SearchFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchFilesBadRequest creates a SearchFilesBadRequest with default headers values
func NewSearchFilesBadRequest() *SearchFilesBadRequest {
	return &SearchFilesBadRequest{}
}

/*SearchFilesBadRequest handles this case with default header values.

Error
*/
type SearchFilesBadRequest struct {
}

func (o *SearchFilesBadRequest) Error() string {
	return fmt.Sprintf("[GET /files/search][%d] searchFilesBadRequest ", 400)
}

func (o *SearchFilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
