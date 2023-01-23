/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BraintreeGatewayUpdate struct for BraintreeGatewayUpdate
type BraintreeGatewayUpdate struct {
	Data BraintreeGatewayUpdateData `json:"data"`
}

// NewBraintreeGatewayUpdate instantiates a new BraintreeGatewayUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBraintreeGatewayUpdate(data BraintreeGatewayUpdateData) *BraintreeGatewayUpdate {
	this := BraintreeGatewayUpdate{}
	this.Data = data
	return &this
}

// NewBraintreeGatewayUpdateWithDefaults instantiates a new BraintreeGatewayUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBraintreeGatewayUpdateWithDefaults() *BraintreeGatewayUpdate {
	this := BraintreeGatewayUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *BraintreeGatewayUpdate) GetData() BraintreeGatewayUpdateData {
	if o == nil {
		var ret BraintreeGatewayUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BraintreeGatewayUpdate) GetDataOk() (*BraintreeGatewayUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BraintreeGatewayUpdate) SetData(v BraintreeGatewayUpdateData) {
	o.Data = v
}

func (o BraintreeGatewayUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBraintreeGatewayUpdate struct {
	value *BraintreeGatewayUpdate
	isSet bool
}

func (v NullableBraintreeGatewayUpdate) Get() *BraintreeGatewayUpdate {
	return v.value
}

func (v *NullableBraintreeGatewayUpdate) Set(val *BraintreeGatewayUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableBraintreeGatewayUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableBraintreeGatewayUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBraintreeGatewayUpdate(val *BraintreeGatewayUpdate) *NullableBraintreeGatewayUpdate {
	return &NullableBraintreeGatewayUpdate{value: val, isSet: true}
}

func (v NullableBraintreeGatewayUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBraintreeGatewayUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
