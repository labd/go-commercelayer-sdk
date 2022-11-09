/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// POSTAttachments201Response struct for POSTAttachments201Response
type POSTAttachments201Response struct {
	Data *POSTAttachments201ResponseData `json:"data,omitempty"`
}

// NewPOSTAttachments201Response instantiates a new POSTAttachments201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAttachments201Response() *POSTAttachments201Response {
	this := POSTAttachments201Response{}
	return &this
}

// NewPOSTAttachments201ResponseWithDefaults instantiates a new POSTAttachments201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAttachments201ResponseWithDefaults() *POSTAttachments201Response {
	this := POSTAttachments201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTAttachments201Response) GetData() POSTAttachments201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTAttachments201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAttachments201Response) GetDataOk() (*POSTAttachments201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTAttachments201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTAttachments201ResponseData and assigns it to the Data field.
func (o *POSTAttachments201Response) SetData(v POSTAttachments201ResponseData) {
	o.Data = &v
}

func (o POSTAttachments201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTAttachments201Response struct {
	value *POSTAttachments201Response
	isSet bool
}

func (v NullablePOSTAttachments201Response) Get() *POSTAttachments201Response {
	return v.value
}

func (v *NullablePOSTAttachments201Response) Set(val *POSTAttachments201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAttachments201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAttachments201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAttachments201Response(val *POSTAttachments201Response) *NullablePOSTAttachments201Response {
	return &NullablePOSTAttachments201Response{value: val, isSet: true}
}

func (v NullablePOSTAttachments201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAttachments201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
