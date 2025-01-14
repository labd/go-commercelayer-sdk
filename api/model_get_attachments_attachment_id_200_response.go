/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.6.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETAttachmentsAttachmentId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETAttachmentsAttachmentId200Response{}

// GETAttachmentsAttachmentId200Response struct for GETAttachmentsAttachmentId200Response
type GETAttachmentsAttachmentId200Response struct {
	Data *GETAttachmentsAttachmentId200ResponseData `json:"data,omitempty"`
}

// NewGETAttachmentsAttachmentId200Response instantiates a new GETAttachmentsAttachmentId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAttachmentsAttachmentId200Response() *GETAttachmentsAttachmentId200Response {
	this := GETAttachmentsAttachmentId200Response{}
	return &this
}

// NewGETAttachmentsAttachmentId200ResponseWithDefaults instantiates a new GETAttachmentsAttachmentId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAttachmentsAttachmentId200ResponseWithDefaults() *GETAttachmentsAttachmentId200Response {
	this := GETAttachmentsAttachmentId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETAttachmentsAttachmentId200Response) GetData() GETAttachmentsAttachmentId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETAttachmentsAttachmentId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAttachmentsAttachmentId200Response) GetDataOk() (*GETAttachmentsAttachmentId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETAttachmentsAttachmentId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETAttachmentsAttachmentId200ResponseData and assigns it to the Data field.
func (o *GETAttachmentsAttachmentId200Response) SetData(v GETAttachmentsAttachmentId200ResponseData) {
	o.Data = &v
}

func (o GETAttachmentsAttachmentId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETAttachmentsAttachmentId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETAttachmentsAttachmentId200Response struct {
	value *GETAttachmentsAttachmentId200Response
	isSet bool
}

func (v NullableGETAttachmentsAttachmentId200Response) Get() *GETAttachmentsAttachmentId200Response {
	return v.value
}

func (v *NullableGETAttachmentsAttachmentId200Response) Set(val *GETAttachmentsAttachmentId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAttachmentsAttachmentId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAttachmentsAttachmentId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAttachmentsAttachmentId200Response(val *GETAttachmentsAttachmentId200Response) *NullableGETAttachmentsAttachmentId200Response {
	return &NullableGETAttachmentsAttachmentId200Response{value: val, isSet: true}
}

func (v NullableGETAttachmentsAttachmentId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAttachmentsAttachmentId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
