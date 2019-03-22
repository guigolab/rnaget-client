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

// GetProjectByIDReader is a Reader for the GetProjectByID structure.
type GetProjectByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProjectByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetProjectByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetProjectByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProjectByIDOK creates a GetProjectByIDOK with default headers values
func NewGetProjectByIDOK() *GetProjectByIDOK {
	return &GetProjectByIDOK{}
}

/*GetProjectByIDOK handles this case with default header values.

successful operation
*/
type GetProjectByIDOK struct {
	Payload *models.Project
}

func (o *GetProjectByIDOK) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectByIdOK  %+v", 200, o.Payload)
}

func (o *GetProjectByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectByIDBadRequest creates a GetProjectByIDBadRequest with default headers values
func NewGetProjectByIDBadRequest() *GetProjectByIDBadRequest {
	return &GetProjectByIDBadRequest{}
}

/*GetProjectByIDBadRequest handles this case with default header values.

Invalid ID supplied
*/
type GetProjectByIDBadRequest struct {
}

func (o *GetProjectByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectByIdBadRequest ", 400)
}

func (o *GetProjectByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectByIDNotFound creates a GetProjectByIDNotFound with default headers values
func NewGetProjectByIDNotFound() *GetProjectByIDNotFound {
	return &GetProjectByIDNotFound{}
}

/*GetProjectByIDNotFound handles this case with default header values.

Project not found
*/
type GetProjectByIDNotFound struct {
}

func (o *GetProjectByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectByIdNotFound ", 404)
}

func (o *GetProjectByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}