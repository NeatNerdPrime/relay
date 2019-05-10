// Code generated by go-swagger; DO NOT EDIT.

package workflow_runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/puppetlabs/nebula/pkg/client/api/models"
)

// ShowWorkflowRunReader is a Reader for the ShowWorkflowRun structure.
type ShowWorkflowRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowWorkflowRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewShowWorkflowRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewShowWorkflowRunOK creates a ShowWorkflowRunOK with default headers values
func NewShowWorkflowRunOK() *ShowWorkflowRunOK {
	return &ShowWorkflowRunOK{}
}

/*ShowWorkflowRunOK handles this case with default header values.

An array of workflow runs
*/
type ShowWorkflowRunOK struct {
	Payload *models.WorkflowRun
}

func (o *ShowWorkflowRunOK) Error() string {
	return fmt.Sprintf("[GET /api/runs/{id}][%d] showWorkflowRunOK  %+v", 200, o.Payload)
}

func (o *ShowWorkflowRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkflowRun)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
