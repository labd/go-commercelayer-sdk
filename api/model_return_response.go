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

// ReturnResponse struct for ReturnResponse
type ReturnResponse struct {
	Data *ReturnResponseData `json:"data,omitempty"`
}

// NewReturnResponse instantiates a new ReturnResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnResponse() *ReturnResponse {
	this := ReturnResponse{}
	return &this
}

// NewReturnResponseWithDefaults instantiates a new ReturnResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnResponseWithDefaults() *ReturnResponse {
	this := ReturnResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ReturnResponse) GetData() ReturnResponseData {
	if o == nil || o.Data == nil {
		var ret ReturnResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnResponse) GetDataOk() (*ReturnResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ReturnResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ReturnResponseData and assigns it to the Data field.
func (o *ReturnResponse) SetData(v ReturnResponseData) {
	o.Data = &v
}

func (o ReturnResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableReturnResponse struct {
	value *ReturnResponse
	isSet bool
}

func (v NullableReturnResponse) Get() *ReturnResponse {
	return v.value
}

func (v *NullableReturnResponse) Set(val *ReturnResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnResponse(val *ReturnResponse) *NullableReturnResponse {
	return &NullableReturnResponse{value: val, isSet: true}
}

func (v NullableReturnResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
