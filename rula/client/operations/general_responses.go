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

// GeneralReader is a Reader for the General structure.
type GeneralReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GeneralReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGeneralOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGeneralBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGeneralOK creates a GeneralOK with default headers values
func NewGeneralOK() *GeneralOK {
	return &GeneralOK{}
}

/* GeneralOK describes a response with status code 200, with default header values.

OK
*/
type GeneralOK struct {
	Payload *models.ModelsResponse
}

func (o *GeneralOK) Error() string {
	return fmt.Sprintf("[GET /][%d] generalOK  %+v", 200, o.Payload)
}
func (o *GeneralOK) GetPayload() *models.ModelsResponse {
	return o.Payload
}

func (o *GeneralOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeneralBadRequest creates a GeneralBadRequest with default headers values
func NewGeneralBadRequest() *GeneralBadRequest {
	return &GeneralBadRequest{}
}

/* GeneralBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GeneralBadRequest struct {
	Payload *models.ModelsResponse
}

func (o *GeneralBadRequest) Error() string {
	return fmt.Sprintf("[GET /][%d] generalBadRequest  %+v", 400, o.Payload)
}
func (o *GeneralBadRequest) GetPayload() *models.ModelsResponse {
	return o.Payload
}

func (o *GeneralBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
