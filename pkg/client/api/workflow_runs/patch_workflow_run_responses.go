// Code generated by go-swagger; DO NOT EDIT.

package workflow_runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/puppetlabs/nebula-cli/pkg/client/api/models"
)

// PatchWorkflowRunReader is a Reader for the PatchWorkflowRun structure.
type PatchWorkflowRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchWorkflowRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchWorkflowRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchWorkflowRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchWorkflowRunOK creates a PatchWorkflowRunOK with default headers values
func NewPatchWorkflowRunOK() *PatchWorkflowRunOK {
	return &PatchWorkflowRunOK{}
}

/*PatchWorkflowRunOK handles this case with default header values.

Metadata about the workflow run
*/
type PatchWorkflowRunOK struct {
	Payload *PatchWorkflowRunOKBody
}

func (o *PatchWorkflowRunOK) Error() string {
	return fmt.Sprintf("[PATCH /api/workflows/{workflowName}/runs/{workflowRunNumber}][%d] patchWorkflowRunOK  %+v", 200, o.Payload)
}

func (o *PatchWorkflowRunOK) GetPayload() *PatchWorkflowRunOKBody {
	return o.Payload
}

func (o *PatchWorkflowRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PatchWorkflowRunOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkflowRunDefault creates a PatchWorkflowRunDefault with default headers values
func NewPatchWorkflowRunDefault(code int) *PatchWorkflowRunDefault {
	return &PatchWorkflowRunDefault{
		_statusCode: code,
	}
}

/*PatchWorkflowRunDefault handles this case with default header values.

An error occurred
*/
type PatchWorkflowRunDefault struct {
	_statusCode int

	Payload *PatchWorkflowRunDefaultBody
}

// Code gets the status code for the patch workflow run default response
func (o *PatchWorkflowRunDefault) Code() int {
	return o._statusCode
}

func (o *PatchWorkflowRunDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/workflows/{workflowName}/runs/{workflowRunNumber}][%d] patchWorkflowRun default  %+v", o._statusCode, o.Payload)
}

func (o *PatchWorkflowRunDefault) GetPayload() *PatchWorkflowRunDefaultBody {
	return o.Payload
}

func (o *PatchWorkflowRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PatchWorkflowRunDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchWorkflowRunBody patch workflow run body
swagger:model PatchWorkflowRunBody
*/
type PatchWorkflowRunBody struct {

	// operation
	Operation *models.WorkflowRunOperation `json:"operation,omitempty"`
}

// Validate validates this patch workflow run body
func (o *PatchWorkflowRunBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchWorkflowRunBody) validateOperation(formats strfmt.Registry) error {

	if swag.IsZero(o.Operation) { // not required
		return nil
	}

	if o.Operation != nil {
		if err := o.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "operation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchWorkflowRunBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchWorkflowRunBody) UnmarshalBinary(b []byte) error {
	var res PatchWorkflowRunBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchWorkflowRunDefaultBody Error response
swagger:model PatchWorkflowRunDefaultBody
*/
type PatchWorkflowRunDefaultBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`
}

// Validate validates this patch workflow run default body
func (o *PatchWorkflowRunDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchWorkflowRunDefaultBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("patchWorkflowRun default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchWorkflowRunDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchWorkflowRunDefaultBody) UnmarshalBinary(b []byte) error {
	var res PatchWorkflowRunDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchWorkflowRunOKBody patch workflow run o k body
swagger:model PatchWorkflowRunOKBody
*/
type PatchWorkflowRunOKBody struct {
	models.Entity

	// run
	Run *models.WorkflowRun `json:"run,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatchWorkflowRunOKBody) UnmarshalJSON(raw []byte) error {
	// PatchWorkflowRunOKBodyAO0
	var patchWorkflowRunOKBodyAO0 models.Entity
	if err := swag.ReadJSON(raw, &patchWorkflowRunOKBodyAO0); err != nil {
		return err
	}
	o.Entity = patchWorkflowRunOKBodyAO0

	// PatchWorkflowRunOKBodyAO1
	var dataPatchWorkflowRunOKBodyAO1 struct {
		Run *models.WorkflowRun `json:"run,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPatchWorkflowRunOKBodyAO1); err != nil {
		return err
	}

	o.Run = dataPatchWorkflowRunOKBodyAO1.Run

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatchWorkflowRunOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	patchWorkflowRunOKBodyAO0, err := swag.WriteJSON(o.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patchWorkflowRunOKBodyAO0)

	var dataPatchWorkflowRunOKBodyAO1 struct {
		Run *models.WorkflowRun `json:"run,omitempty"`
	}

	dataPatchWorkflowRunOKBodyAO1.Run = o.Run

	jsonDataPatchWorkflowRunOKBodyAO1, errPatchWorkflowRunOKBodyAO1 := swag.WriteJSON(dataPatchWorkflowRunOKBodyAO1)
	if errPatchWorkflowRunOKBodyAO1 != nil {
		return nil, errPatchWorkflowRunOKBodyAO1
	}
	_parts = append(_parts, jsonDataPatchWorkflowRunOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patch workflow run o k body
func (o *PatchWorkflowRunOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Entity
	if err := o.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRun(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchWorkflowRunOKBody) validateRun(formats strfmt.Registry) error {

	if swag.IsZero(o.Run) { // not required
		return nil
	}

	if o.Run != nil {
		if err := o.Run.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("patchWorkflowRunOK" + "." + "run")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchWorkflowRunOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchWorkflowRunOKBody) UnmarshalBinary(b []byte) error {
	var res PatchWorkflowRunOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}