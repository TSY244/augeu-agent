// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"augeu/public/pkg/swaggerCore/models"
)

// PostUploadUserInfoOKCode is the HTTP code returned for type PostUploadUserInfoOK
const PostUploadUserInfoOKCode int = 200

/*
PostUploadUserInfoOK 上传成功

swagger:response postUploadUserInfoOK
*/
type PostUploadUserInfoOK struct {

	/*
	  In: Body
	*/
	Payload *models.SuccessResponse `json:"body,omitempty"`
}

// NewPostUploadUserInfoOK creates PostUploadUserInfoOK with default headers values
func NewPostUploadUserInfoOK() *PostUploadUserInfoOK {

	return &PostUploadUserInfoOK{}
}

// WithPayload adds the payload to the post upload user info o k response
func (o *PostUploadUserInfoOK) WithPayload(payload *models.SuccessResponse) *PostUploadUserInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post upload user info o k response
func (o *PostUploadUserInfoOK) SetPayload(payload *models.SuccessResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUploadUserInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUploadUserInfoBadRequestCode is the HTTP code returned for type PostUploadUserInfoBadRequest
const PostUploadUserInfoBadRequestCode int = 400

/*
PostUploadUserInfoBadRequest 输入参数错误

swagger:response postUploadUserInfoBadRequest
*/
type PostUploadUserInfoBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.BadRequestError `json:"body,omitempty"`
}

// NewPostUploadUserInfoBadRequest creates PostUploadUserInfoBadRequest with default headers values
func NewPostUploadUserInfoBadRequest() *PostUploadUserInfoBadRequest {

	return &PostUploadUserInfoBadRequest{}
}

// WithPayload adds the payload to the post upload user info bad request response
func (o *PostUploadUserInfoBadRequest) WithPayload(payload *models.BadRequestError) *PostUploadUserInfoBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post upload user info bad request response
func (o *PostUploadUserInfoBadRequest) SetPayload(payload *models.BadRequestError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUploadUserInfoBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUploadUserInfoForbiddenCode is the HTTP code returned for type PostUploadUserInfoForbidden
const PostUploadUserInfoForbiddenCode int = 403

/*
PostUploadUserInfoForbidden 没有权限

swagger:response postUploadUserInfoForbidden
*/
type PostUploadUserInfoForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.UnauthorizedError `json:"body,omitempty"`
}

// NewPostUploadUserInfoForbidden creates PostUploadUserInfoForbidden with default headers values
func NewPostUploadUserInfoForbidden() *PostUploadUserInfoForbidden {

	return &PostUploadUserInfoForbidden{}
}

// WithPayload adds the payload to the post upload user info forbidden response
func (o *PostUploadUserInfoForbidden) WithPayload(payload *models.UnauthorizedError) *PostUploadUserInfoForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post upload user info forbidden response
func (o *PostUploadUserInfoForbidden) SetPayload(payload *models.UnauthorizedError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUploadUserInfoForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUploadUserInfoInternalServerErrorCode is the HTTP code returned for type PostUploadUserInfoInternalServerError
const PostUploadUserInfoInternalServerErrorCode int = 500

/*
PostUploadUserInfoInternalServerError 内部错误

swagger:response postUploadUserInfoInternalServerError
*/
type PostUploadUserInfoInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ActionFailure `json:"body,omitempty"`
}

// NewPostUploadUserInfoInternalServerError creates PostUploadUserInfoInternalServerError with default headers values
func NewPostUploadUserInfoInternalServerError() *PostUploadUserInfoInternalServerError {

	return &PostUploadUserInfoInternalServerError{}
}

// WithPayload adds the payload to the post upload user info internal server error response
func (o *PostUploadUserInfoInternalServerError) WithPayload(payload *models.ActionFailure) *PostUploadUserInfoInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post upload user info internal server error response
func (o *PostUploadUserInfoInternalServerError) SetPayload(payload *models.ActionFailure) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUploadUserInfoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
