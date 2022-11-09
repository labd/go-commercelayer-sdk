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

// ManualGatewayCreate struct for ManualGatewayCreate
type ManualGatewayCreate struct {
	Data ManualGatewayCreateData `json:"data"`
}

// NewManualGatewayCreate instantiates a new ManualGatewayCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualGatewayCreate(data ManualGatewayCreateData) *ManualGatewayCreate {
	this := ManualGatewayCreate{}
	this.Data = data
	return &this
}

// NewManualGatewayCreateWithDefaults instantiates a new ManualGatewayCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualGatewayCreateWithDefaults() *ManualGatewayCreate {
	this := ManualGatewayCreate{}
	return &this
}

// GetData returns the Data field value
func (o *ManualGatewayCreate) GetData() ManualGatewayCreateData {
	if o == nil {
		var ret ManualGatewayCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ManualGatewayCreate) GetDataOk() (*ManualGatewayCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ManualGatewayCreate) SetData(v ManualGatewayCreateData) {
	o.Data = v
}

func (o ManualGatewayCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableManualGatewayCreate struct {
	value *ManualGatewayCreate
	isSet bool
}

func (v NullableManualGatewayCreate) Get() *ManualGatewayCreate {
	return v.value
}

func (v *NullableManualGatewayCreate) Set(val *ManualGatewayCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableManualGatewayCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableManualGatewayCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualGatewayCreate(val *ManualGatewayCreate) *NullableManualGatewayCreate {
	return &NullableManualGatewayCreate{value: val, isSet: true}
}

func (v NullableManualGatewayCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualGatewayCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
