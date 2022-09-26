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

// ApplicationResponse struct for ApplicationResponse
type ApplicationResponse struct {
	Data *ApplicationResponseData `json:"data,omitempty"`
}

// NewApplicationResponse instantiates a new ApplicationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationResponse() *ApplicationResponse {
	this := ApplicationResponse{}
	return &this
}

// NewApplicationResponseWithDefaults instantiates a new ApplicationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationResponseWithDefaults() *ApplicationResponse {
	this := ApplicationResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ApplicationResponse) GetData() ApplicationResponseData {
	if o == nil || o.Data == nil {
		var ret ApplicationResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetDataOk() (*ApplicationResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ApplicationResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ApplicationResponseData and assigns it to the Data field.
func (o *ApplicationResponse) SetData(v ApplicationResponseData) {
	o.Data = &v
}

func (o ApplicationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationResponse struct {
	value *ApplicationResponse
	isSet bool
}

func (v NullableApplicationResponse) Get() *ApplicationResponse {
	return v.value
}

func (v *NullableApplicationResponse) Set(val *ApplicationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationResponse(val *ApplicationResponse) *NullableApplicationResponse {
	return &NullableApplicationResponse{value: val, isSet: true}
}

func (v NullableApplicationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
