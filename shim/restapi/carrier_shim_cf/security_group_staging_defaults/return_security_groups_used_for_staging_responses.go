// Code generated by go-swagger; DO NOT EDIT.

package security_group_staging_defaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// ReturnSecurityGroupsUsedForStagingOKCode is the HTTP code returned for type ReturnSecurityGroupsUsedForStagingOK
const ReturnSecurityGroupsUsedForStagingOKCode int = 200

/*ReturnSecurityGroupsUsedForStagingOK successful response

swagger:response returnSecurityGroupsUsedForStagingOK
*/
type ReturnSecurityGroupsUsedForStagingOK struct {

	/*
	  In: Body
	*/
	Payload *models.ReturnSecurityGroupsUsedForStagingResponsePaged `json:"body,omitempty"`
}

// NewReturnSecurityGroupsUsedForStagingOK creates ReturnSecurityGroupsUsedForStagingOK with default headers values
func NewReturnSecurityGroupsUsedForStagingOK() *ReturnSecurityGroupsUsedForStagingOK {

	return &ReturnSecurityGroupsUsedForStagingOK{}
}

// WithPayload adds the payload to the return security groups used for staging o k response
func (o *ReturnSecurityGroupsUsedForStagingOK) WithPayload(payload *models.ReturnSecurityGroupsUsedForStagingResponsePaged) *ReturnSecurityGroupsUsedForStagingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the return security groups used for staging o k response
func (o *ReturnSecurityGroupsUsedForStagingOK) SetPayload(payload *models.ReturnSecurityGroupsUsedForStagingResponsePaged) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReturnSecurityGroupsUsedForStagingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
