/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AdyenGatewayCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdyenGatewayCreate{}

// AdyenGatewayCreate struct for AdyenGatewayCreate
type AdyenGatewayCreate struct {
	Data AdyenGatewayCreateData `json:"data"`
}

// NewAdyenGatewayCreate instantiates a new AdyenGatewayCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenGatewayCreate(data AdyenGatewayCreateData) *AdyenGatewayCreate {
	this := AdyenGatewayCreate{}
	this.Data = data
	return &this
}

// NewAdyenGatewayCreateWithDefaults instantiates a new AdyenGatewayCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenGatewayCreateWithDefaults() *AdyenGatewayCreate {
	this := AdyenGatewayCreate{}
	return &this
}

// GetData returns the Data field value
func (o *AdyenGatewayCreate) GetData() AdyenGatewayCreateData {
	if o == nil {
		var ret AdyenGatewayCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AdyenGatewayCreate) GetDataOk() (*AdyenGatewayCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AdyenGatewayCreate) SetData(v AdyenGatewayCreateData) {
	o.Data = v
}

func (o AdyenGatewayCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdyenGatewayCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAdyenGatewayCreate struct {
	value *AdyenGatewayCreate
	isSet bool
}

func (v NullableAdyenGatewayCreate) Get() *AdyenGatewayCreate {
	return v.value
}

func (v *NullableAdyenGatewayCreate) Set(val *AdyenGatewayCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenGatewayCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenGatewayCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenGatewayCreate(val *AdyenGatewayCreate) *NullableAdyenGatewayCreate {
	return &NullableAdyenGatewayCreate{value: val, isSet: true}
}

func (v NullableAdyenGatewayCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenGatewayCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
