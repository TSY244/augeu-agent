// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"augeu/public/pkg/swaggerCore/models"
)

// PostGetProcessEventOKCode is the HTTP code returned for type PostGetProcessEventOK
const PostGetProcessEventOKCode int = 200

/*
PostGetProcessEventOK 成功返回事件数组

swagger:response postGetProcessEventOK
*/
type PostGetProcessEventOK struct {

	/*
	  In: Body
	*/
	Payload []*models.EventCreateProcess `json:"body,omitempty"`
}

// NewPostGetProcessEventOK creates PostGetProcessEventOK with default headers values
func NewPostGetProcessEventOK() *PostGetProcessEventOK {

	return &PostGetProcessEventOK{}
}

// WithPayload adds the payload to the post get process event o k response
func (o *PostGetProcessEventOK) WithPayload(payload []*models.EventCreateProcess) *PostGetProcessEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post get process event o k response
func (o *PostGetProcessEventOK) SetPayload(payload []*models.EventCreateProcess) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostGetProcessEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.EventCreateProcess, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostGetProcessEventBadRequestCode is the HTTP code returned for type PostGetProcessEventBadRequest
const PostGetProcessEventBadRequestCode int = 400

/*
PostGetProcessEventBadRequest 输入参数错误

swagger:response postGetProcessEventBadRequest
*/
type PostGetProcessEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.BadRequestError `json:"body,omitempty"`
}

// NewPostGetProcessEventBadRequest creates PostGetProcessEventBadRequest with default headers values
func NewPostGetProcessEventBadRequest() *PostGetProcessEventBadRequest {

	return &PostGetProcessEventBadRequest{}
}

// WithPayload adds the payload to the post get process event bad request response
func (o *PostGetProcessEventBadRequest) WithPayload(payload *models.BadRequestError) *PostGetProcessEventBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post get process event bad request response
func (o *PostGetProcessEventBadRequest) SetPayload(payload *models.BadRequestError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostGetProcessEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostGetProcessEventForbiddenCode is the HTTP code returned for type PostGetProcessEventForbidden
const PostGetProcessEventForbiddenCode int = 403

/*
PostGetProcessEventForbidden 没有权限

swagger:response postGetProcessEventForbidden
*/
type PostGetProcessEventForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.UnauthorizedError `json:"body,omitempty"`
}

// NewPostGetProcessEventForbidden creates PostGetProcessEventForbidden with default headers values
func NewPostGetProcessEventForbidden() *PostGetProcessEventForbidden {

	return &PostGetProcessEventForbidden{}
}

// WithPayload adds the payload to the post get process event forbidden response
func (o *PostGetProcessEventForbidden) WithPayload(payload *models.UnauthorizedError) *PostGetProcessEventForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post get process event forbidden response
func (o *PostGetProcessEventForbidden) SetPayload(payload *models.UnauthorizedError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostGetProcessEventForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostGetProcessEventInternalServerErrorCode is the HTTP code returned for type PostGetProcessEventInternalServerError
const PostGetProcessEventInternalServerErrorCode int = 500

/*
PostGetProcessEventInternalServerError 内部错误

swagger:response postGetProcessEventInternalServerError
*/
type PostGetProcessEventInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ActionFailure `json:"body,omitempty"`
}

// NewPostGetProcessEventInternalServerError creates PostGetProcessEventInternalServerError with default headers values
func NewPostGetProcessEventInternalServerError() *PostGetProcessEventInternalServerError {

	return &PostGetProcessEventInternalServerError{}
}

// WithPayload adds the payload to the post get process event internal server error response
func (o *PostGetProcessEventInternalServerError) WithPayload(payload *models.ActionFailure) *PostGetProcessEventInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post get process event internal server error response
func (o *PostGetProcessEventInternalServerError) SetPayload(payload *models.ActionFailure) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostGetProcessEventInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
