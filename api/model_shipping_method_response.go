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

// ShippingMethodResponse struct for ShippingMethodResponse
type ShippingMethodResponse struct {
	Data *ShippingMethodResponseData `json:"data,omitempty"`
}

// NewShippingMethodResponse instantiates a new ShippingMethodResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodResponse() *ShippingMethodResponse {
	this := ShippingMethodResponse{}
	return &this
}

// NewShippingMethodResponseWithDefaults instantiates a new ShippingMethodResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodResponseWithDefaults() *ShippingMethodResponse {
	this := ShippingMethodResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ShippingMethodResponse) GetData() ShippingMethodResponseData {
	if o == nil || o.Data == nil {
		var ret ShippingMethodResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodResponse) GetDataOk() (*ShippingMethodResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ShippingMethodResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ShippingMethodResponseData and assigns it to the Data field.
func (o *ShippingMethodResponse) SetData(v ShippingMethodResponseData) {
	o.Data = &v
}

func (o ShippingMethodResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableShippingMethodResponse struct {
	value *ShippingMethodResponse
	isSet bool
}

func (v NullableShippingMethodResponse) Get() *ShippingMethodResponse {
	return v.value
}

func (v *NullableShippingMethodResponse) Set(val *ShippingMethodResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodResponse(val *ShippingMethodResponse) *NullableShippingMethodResponse {
	return &NullableShippingMethodResponse{value: val, isSet: true}
}

func (v NullableShippingMethodResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
