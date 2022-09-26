/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AttachmentResponse struct for AttachmentResponse
type AttachmentResponse struct {
	Data *AttachmentResponseData `json:"data,omitempty"`
}

// NewAttachmentResponse instantiates a new AttachmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentResponse() *AttachmentResponse {
	this := AttachmentResponse{}
	return &this
}

// NewAttachmentResponseWithDefaults instantiates a new AttachmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentResponseWithDefaults() *AttachmentResponse {
	this := AttachmentResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AttachmentResponse) GetData() AttachmentResponseData {
	if o == nil || o.Data == nil {
		var ret AttachmentResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentResponse) GetDataOk() (*AttachmentResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AttachmentResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given AttachmentResponseData and assigns it to the Data field.
func (o *AttachmentResponse) SetData(v AttachmentResponseData) {
	o.Data = &v
}

func (o AttachmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAttachmentResponse struct {
	value *AttachmentResponse
	isSet bool
}

func (v NullableAttachmentResponse) Get() *AttachmentResponse {
	return v.value
}

func (v *NullableAttachmentResponse) Set(val *AttachmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentResponse(val *AttachmentResponse) *NullableAttachmentResponse {
	return &NullableAttachmentResponse{value: val, isSet: true}
}

func (v NullableAttachmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
