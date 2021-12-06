// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/atrariksa/fastrogos/rula/models"
)

// DeleteUserReader is a Reader for the DeleteUser structure.
type DeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserOK creates a DeleteUserOK with default headers values
func NewDeleteUserOK() *DeleteUserOK {
	return &DeleteUserOK{}
}

/* DeleteUserOK describes a response with status code 200, with default header values.

OK
*/
type DeleteUserOK struct {
	Payload *models.ModelsResponse
}

func (o *DeleteUserOK) Error() string {
	return fmt.Sprintf("[DELETE /user/][%d] deleteUserOK  %+v", 200, o.Payload)
}
func (o *DeleteUserOK) GetPayload() *models.ModelsResponse {
	return o.Payload
}

func (o *DeleteUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserBadRequest creates a DeleteUserBadRequest with default headers values
func NewDeleteUserBadRequest() *DeleteUserBadRequest {
	return &DeleteUserBadRequest{}
}

/* DeleteUserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteUserBadRequest struct {
	Payload *models.ModelsResponse
}

func (o *DeleteUserBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /user/][%d] deleteUserBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteUserBadRequest) GetPayload() *models.ModelsResponse {
	return o.Payload
}

func (o *DeleteUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
