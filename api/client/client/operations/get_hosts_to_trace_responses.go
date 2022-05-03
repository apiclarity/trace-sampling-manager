// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/openclarity/trace-sampling-manager/api/client/models"
)

// GetHostsToTraceReader is a Reader for the GetHostsToTrace structure.
type GetHostsToTraceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHostsToTraceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHostsToTraceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHostsToTraceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHostsToTraceOK creates a GetHostsToTraceOK with default headers values
func NewGetHostsToTraceOK() *GetHostsToTraceOK {
	return &GetHostsToTraceOK{}
}

/* GetHostsToTraceOK describes a response with status code 200, with default header values.

Success
*/
type GetHostsToTraceOK struct {
	Payload *GetHostsToTraceOKBody
}

func (o *GetHostsToTraceOK) Error() string {
	return fmt.Sprintf("[GET /hostsToTrace][%d] getHostsToTraceOK  %+v", 200, o.Payload)
}
func (o *GetHostsToTraceOK) GetPayload() *GetHostsToTraceOKBody {
	return o.Payload
}

func (o *GetHostsToTraceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetHostsToTraceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostsToTraceDefault creates a GetHostsToTraceDefault with default headers values
func NewGetHostsToTraceDefault(code int) *GetHostsToTraceDefault {
	return &GetHostsToTraceDefault{
		_statusCode: code,
	}
}

/* GetHostsToTraceDefault describes a response with status code -1, with default header values.

unknown error
*/
type GetHostsToTraceDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the get hosts to trace default response
func (o *GetHostsToTraceDefault) Code() int {
	return o._statusCode
}

func (o *GetHostsToTraceDefault) Error() string {
	return fmt.Sprintf("[GET /hostsToTrace][%d] GetHostsToTrace default  %+v", o._statusCode, o.Payload)
}
func (o *GetHostsToTraceDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *GetHostsToTraceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetHostsToTraceOKBody get hosts to trace o k body
swagger:model GetHostsToTraceOKBody
*/
type GetHostsToTraceOKBody struct {

	// List of hosts to trace in the format of hostname:port
	Hosts []string `json:"hosts"`
}

// Validate validates this get hosts to trace o k body
func (o *GetHostsToTraceOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get hosts to trace o k body based on context it is used
func (o *GetHostsToTraceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetHostsToTraceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHostsToTraceOKBody) UnmarshalBinary(b []byte) error {
	var res GetHostsToTraceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
