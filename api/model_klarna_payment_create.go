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

// checks if the KlarnaPaymentCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KlarnaPaymentCreate{}

// KlarnaPaymentCreate struct for KlarnaPaymentCreate
type KlarnaPaymentCreate struct {
	Data KlarnaPaymentCreateData `json:"data"`
}

// NewKlarnaPaymentCreate instantiates a new KlarnaPaymentCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlarnaPaymentCreate(data KlarnaPaymentCreateData) *KlarnaPaymentCreate {
	this := KlarnaPaymentCreate{}
	this.Data = data
	return &this
}

// NewKlarnaPaymentCreateWithDefaults instantiates a new KlarnaPaymentCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlarnaPaymentCreateWithDefaults() *KlarnaPaymentCreate {
	this := KlarnaPaymentCreate{}
	return &this
}

// GetData returns the Data field value
func (o *KlarnaPaymentCreate) GetData() KlarnaPaymentCreateData {
	if o == nil {
		var ret KlarnaPaymentCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *KlarnaPaymentCreate) GetDataOk() (*KlarnaPaymentCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *KlarnaPaymentCreate) SetData(v KlarnaPaymentCreateData) {
	o.Data = v
}

func (o KlarnaPaymentCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlarnaPaymentCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableKlarnaPaymentCreate struct {
	value *KlarnaPaymentCreate
	isSet bool
}

func (v NullableKlarnaPaymentCreate) Get() *KlarnaPaymentCreate {
	return v.value
}

func (v *NullableKlarnaPaymentCreate) Set(val *KlarnaPaymentCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableKlarnaPaymentCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableKlarnaPaymentCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlarnaPaymentCreate(val *KlarnaPaymentCreate) *NullableKlarnaPaymentCreate {
	return &NullableKlarnaPaymentCreate{value: val, isSet: true}
}

func (v NullableKlarnaPaymentCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlarnaPaymentCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
