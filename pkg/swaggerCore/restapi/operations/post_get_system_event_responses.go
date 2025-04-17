// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"augeu/public/pkg/swaggerCore/models"
)

// PostGetSystemEventOKCode is the HTTP code returned for type PostGetSystemEventOK
const PostGetSystemEventOKCode int = 200

/*
PostGetSystemEventOK 成功返回事件数组

swagger:response postGetSystemEventOK
*/
type PostGetSystemEventOK struct {

	/*
	  In: Body
	*/
	Payload []*models.SystemEvent `json:"body,omitempty"`
}

// NewPostGetSystemEventOK creates PostGetSystemEventOK with default headers values
func NewPostGetSystemEventOK() *PostGetSystemEventOK {

	return &PostGetSystemEventOK{}
}

// WithPayload adds the payload to the post get system event o k response
func (o *PostGetSystemEventOK) WithPayload(payload []*models.SystemEvent) *PostGetSystemEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post get system event o k response
func (o *PostGetSystemEventOK) SetPayload(payload []*models.SystemEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostGetSystemEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.SystemEvent, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostGetSystemEventBadRequestCode is the HTTP code returned for type PostGetSystemEventBadRequest
const PostGetSystemEventBadRequestCode int = 400

/*
PostGetSystemEventBadRequest 输入参数错误

swagger:response postGetSystemEventBadRequest
*/
type PostGetSystemEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.BadRequestError `json:"body,omitempty"`
}

// NewPostGetSystemEventBadRequest creates PostGetSystemEventBadRequest with default headers values
func NewPostGetSystemEventBadRequest() *PostGetSystemEventBadRequest {

	return &PostGetSystemEventBadRequest{}
}

// WithPayload adds the payload to the post get system event bad request response
func (o *PostGetSystemEventBadRequest) WithPayload(payload *models.BadRequestError) *PostGetSystemEventBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post get system event bad request response
func (o *PostGetSystemEventBadRequest) SetPayload(payload *models.BadRequestError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostGetSystemEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostGetSystemEventForbiddenCode is the HTTP code returned for type PostGetSystemEventForbidden
const PostGetSystemEventForbiddenCode int = 403

/*
PostGetSystemEventForbidden 没有权限

swagger:response postGetSystemEventForbidden
*/
type PostGetSystemEventForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.UnauthorizedError `json:"body,omitempty"`
}

// NewPostGetSystemEventForbidden creates PostGetSystemEventForbidden with default headers values
func NewPostGetSystemEventForbidden() *PostGetSystemEventForbidden {

	return &PostGetSystemEventForbidden{}
}

// WithPayload adds the payload to the post get system event forbidden response
func (o *PostGetSystemEventForbidden) WithPayload(payload *models.UnauthorizedError) *PostGetSystemEventForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post get system event forbidden response
func (o *PostGetSystemEventForbidden) SetPayload(payload *models.UnauthorizedError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostGetSystemEventForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostGetSystemEventInternalServerErrorCode is the HTTP code returned for type PostGetSystemEventInternalServerError
const PostGetSystemEventInternalServerErrorCode int = 500

/*
PostGetSystemEventInternalServerError 内部错误

swagger:response postGetSystemEventInternalServerError
*/
type PostGetSystemEventInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ActionFailure `json:"body,omitempty"`
}

// NewPostGetSystemEventInternalServerError creates PostGetSystemEventInternalServerError with default headers values
func NewPostGetSystemEventInternalServerError() *PostGetSystemEventInternalServerError {

	return &PostGetSystemEventInternalServerError{}
}

// WithPayload adds the payload to the post get system event internal server error response
func (o *PostGetSystemEventInternalServerError) WithPayload(payload *models.ActionFailure) *PostGetSystemEventInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post get system event internal server error response
func (o *PostGetSystemEventInternalServerError) SetPayload(payload *models.ActionFailure) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostGetSystemEventInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
