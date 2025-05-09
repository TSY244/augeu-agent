// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"augeu/public/pkg/swaggerCore/models"
)

// PostGetLoginEventOKCode is the HTTP code returned for type PostGetLoginEventOK
const PostGetLoginEventOKCode int = 200

/*
PostGetLoginEventOK 成功返回事件数组

swagger:response postGetLoginEventOK
*/
type PostGetLoginEventOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetLoginEventResponse `json:"body,omitempty"`
}

// NewPostGetLoginEventOK creates PostGetLoginEventOK with default headers values
func NewPostGetLoginEventOK() *PostGetLoginEventOK {

	return &PostGetLoginEventOK{}
}

// WithPayload adds the payload to the post get login event o k response
func (o *PostGetLoginEventOK) WithPayload(payload *models.GetLoginEventResponse) *PostGetLoginEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post get login event o k response
func (o *PostGetLoginEventOK) SetPayload(payload *models.GetLoginEventResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostGetLoginEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostGetLoginEventBadRequestCode is the HTTP code returned for type PostGetLoginEventBadRequest
const PostGetLoginEventBadRequestCode int = 400

/*
PostGetLoginEventBadRequest 输入参数错误

swagger:response postGetLoginEventBadRequest
*/
type PostGetLoginEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.BadRequestError `json:"body,omitempty"`
}

// NewPostGetLoginEventBadRequest creates PostGetLoginEventBadRequest with default headers values
func NewPostGetLoginEventBadRequest() *PostGetLoginEventBadRequest {

	return &PostGetLoginEventBadRequest{}
}

// WithPayload adds the payload to the post get login event bad request response
func (o *PostGetLoginEventBadRequest) WithPayload(payload *models.BadRequestError) *PostGetLoginEventBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post get login event bad request response
func (o *PostGetLoginEventBadRequest) SetPayload(payload *models.BadRequestError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostGetLoginEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostGetLoginEventForbiddenCode is the HTTP code returned for type PostGetLoginEventForbidden
const PostGetLoginEventForbiddenCode int = 403

/*
PostGetLoginEventForbidden 没有权限

swagger:response postGetLoginEventForbidden
*/
type PostGetLoginEventForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.UnauthorizedError `json:"body,omitempty"`
}

// NewPostGetLoginEventForbidden creates PostGetLoginEventForbidden with default headers values
func NewPostGetLoginEventForbidden() *PostGetLoginEventForbidden {

	return &PostGetLoginEventForbidden{}
}

// WithPayload adds the payload to the post get login event forbidden response
func (o *PostGetLoginEventForbidden) WithPayload(payload *models.UnauthorizedError) *PostGetLoginEventForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post get login event forbidden response
func (o *PostGetLoginEventForbidden) SetPayload(payload *models.UnauthorizedError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostGetLoginEventForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostGetLoginEventInternalServerErrorCode is the HTTP code returned for type PostGetLoginEventInternalServerError
const PostGetLoginEventInternalServerErrorCode int = 500

/*
PostGetLoginEventInternalServerError 内部错误

swagger:response postGetLoginEventInternalServerError
*/
type PostGetLoginEventInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ActionFailure `json:"body,omitempty"`
}

// NewPostGetLoginEventInternalServerError creates PostGetLoginEventInternalServerError with default headers values
func NewPostGetLoginEventInternalServerError() *PostGetLoginEventInternalServerError {

	return &PostGetLoginEventInternalServerError{}
}

// WithPayload adds the payload to the post get login event internal server error response
func (o *PostGetLoginEventInternalServerError) WithPayload(payload *models.ActionFailure) *PostGetLoginEventInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post get login event internal server error response
func (o *PostGetLoginEventInternalServerError) SetPayload(payload *models.ActionFailure) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostGetLoginEventInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
