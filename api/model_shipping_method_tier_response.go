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

// ShippingMethodTierResponse struct for ShippingMethodTierResponse
type ShippingMethodTierResponse struct {
	Data *ShippingMethodTierResponseData `json:"data,omitempty"`
}

// NewShippingMethodTierResponse instantiates a new ShippingMethodTierResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodTierResponse() *ShippingMethodTierResponse {
	this := ShippingMethodTierResponse{}
	return &this
}

// NewShippingMethodTierResponseWithDefaults instantiates a new ShippingMethodTierResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodTierResponseWithDefaults() *ShippingMethodTierResponse {
	this := ShippingMethodTierResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ShippingMethodTierResponse) GetData() ShippingMethodTierResponseData {
	if o == nil || o.Data == nil {
		var ret ShippingMethodTierResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodTierResponse) GetDataOk() (*ShippingMethodTierResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ShippingMethodTierResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ShippingMethodTierResponseData and assigns it to the Data field.
func (o *ShippingMethodTierResponse) SetData(v ShippingMethodTierResponseData) {
	o.Data = &v
}

func (o ShippingMethodTierResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableShippingMethodTierResponse struct {
	value *ShippingMethodTierResponse
	isSet bool
}

func (v NullableShippingMethodTierResponse) Get() *ShippingMethodTierResponse {
	return v.value
}

func (v *NullableShippingMethodTierResponse) Set(val *ShippingMethodTierResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodTierResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodTierResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodTierResponse(val *ShippingMethodTierResponse) *NullableShippingMethodTierResponse {
	return &NullableShippingMethodTierResponse{value: val, isSet: true}
}

func (v NullableShippingMethodTierResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodTierResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
