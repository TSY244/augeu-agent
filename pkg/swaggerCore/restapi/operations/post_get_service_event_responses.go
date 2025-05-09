// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"augeu/public/pkg/swaggerCore/models"
)

// PostGetServiceEventOKCode is the HTTP code returned for type PostGetServiceEventOK
const PostGetServiceEventOKCode int = 200

/*
PostGetServiceEventOK 成功返回服务事件数组

swagger:response postGetServiceEventOK
*/
type PostGetServiceEventOK struct {

	/*
	  In: Body
	*/
	Payload []*models.ServiceInfo `json:"body,omitempty"`
}

// NewPostGetServiceEventOK creates PostGetServiceEventOK with default headers values
func NewPostGetServiceEventOK() *PostGetServiceEventOK {

	return &PostGetServiceEventOK{}
}

// WithPayload adds the payload to the post get service event o k response
func (o *PostGetServiceEventOK) WithPayload(payload []*models.ServiceInfo) *PostGetServiceEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post get service event o k response
func (o *PostGetServiceEventOK) SetPayload(payload []*models.ServiceInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostGetServiceEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.ServiceInfo, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostGetServiceEventBadRequestCode is the HTTP code returned for type PostGetServiceEventBadRequest
const PostGetServiceEventBadRequestCode int = 400

/*
PostGetServiceEventBadRequest 输入参数错误

swagger:response postGetServiceEventBadRequest
*/
type PostGetServiceEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.BadRequestError `json:"body,omitempty"`
}

// NewPostGetServiceEventBadRequest creates PostGetServiceEventBadRequest with default headers values
func NewPostGetServiceEventBadRequest() *PostGetServiceEventBadRequest {

	return &PostGetServiceEventBadRequest{}
}

// WithPayload adds the payload to the post get service event bad request response
func (o *PostGetServiceEventBadRequest) WithPayload(payload *models.BadRequestError) *PostGetServiceEventBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post get service event bad request response
func (o *PostGetServiceEventBadRequest) SetPayload(payload *models.BadRequestError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostGetServiceEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostGetServiceEventForbiddenCode is the HTTP code returned for type PostGetServiceEventForbidden
const PostGetServiceEventForbiddenCode int = 403

/*
PostGetServiceEventForbidden 没有权限

swagger:response postGetServiceEventForbidden
*/
type PostGetServiceEventForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.UnauthorizedError `json:"body,omitempty"`
}

// NewPostGetServiceEventForbidden creates PostGetServiceEventForbidden with default headers values
func NewPostGetServiceEventForbidden() *PostGetServiceEventForbidden {

	return &PostGetServiceEventForbidden{}
}

// WithPayload adds the payload to the post get service event forbidden response
func (o *PostGetServiceEventForbidden) WithPayload(payload *models.UnauthorizedError) *PostGetServiceEventForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post get service event forbidden response
func (o *PostGetServiceEventForbidden) SetPayload(payload *models.UnauthorizedError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostGetServiceEventForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostGetServiceEventInternalServerErrorCode is the HTTP code returned for type PostGetServiceEventInternalServerError
const PostGetServiceEventInternalServerErrorCode int = 500

/*
PostGetServiceEventInternalServerError 内部错误

swagger:response postGetServiceEventInternalServerError
*/
type PostGetServiceEventInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ActionFailure `json:"body,omitempty"`
}

// NewPostGetServiceEventInternalServerError creates PostGetServiceEventInternalServerError with default headers values
func NewPostGetServiceEventInternalServerError() *PostGetServiceEventInternalServerError {

	return &PostGetServiceEventInternalServerError{}
}

// WithPayload adds the payload to the post get service event internal server error response
func (o *PostGetServiceEventInternalServerError) WithPayload(payload *models.ActionFailure) *PostGetServiceEventInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post get service event internal server error response
func (o *PostGetServiceEventInternalServerError) SetPayload(payload *models.ActionFailure) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostGetServiceEventInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
