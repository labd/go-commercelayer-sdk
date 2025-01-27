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

// checks if the AdyenPaymentCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdyenPaymentCreate{}

// AdyenPaymentCreate struct for AdyenPaymentCreate
type AdyenPaymentCreate struct {
	Data AdyenPaymentCreateData `json:"data"`
}

// NewAdyenPaymentCreate instantiates a new AdyenPaymentCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenPaymentCreate(data AdyenPaymentCreateData) *AdyenPaymentCreate {
	this := AdyenPaymentCreate{}
	this.Data = data
	return &this
}

// NewAdyenPaymentCreateWithDefaults instantiates a new AdyenPaymentCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenPaymentCreateWithDefaults() *AdyenPaymentCreate {
	this := AdyenPaymentCreate{}
	return &this
}

// GetData returns the Data field value
func (o *AdyenPaymentCreate) GetData() AdyenPaymentCreateData {
	if o == nil {
		var ret AdyenPaymentCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AdyenPaymentCreate) GetDataOk() (*AdyenPaymentCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AdyenPaymentCreate) SetData(v AdyenPaymentCreateData) {
	o.Data = v
}

func (o AdyenPaymentCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdyenPaymentCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAdyenPaymentCreate struct {
	value *AdyenPaymentCreate
	isSet bool
}

func (v NullableAdyenPaymentCreate) Get() *AdyenPaymentCreate {
	return v.value
}

func (v *NullableAdyenPaymentCreate) Set(val *AdyenPaymentCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenPaymentCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenPaymentCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenPaymentCreate(val *AdyenPaymentCreate) *NullableAdyenPaymentCreate {
	return &NullableAdyenPaymentCreate{value: val, isSet: true}
}

func (v NullableAdyenPaymentCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenPaymentCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
