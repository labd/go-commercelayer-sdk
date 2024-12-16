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

// checks if the AxervePaymentCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AxervePaymentCreate{}

// AxervePaymentCreate struct for AxervePaymentCreate
type AxervePaymentCreate struct {
	Data AxervePaymentCreateData `json:"data"`
}

// NewAxervePaymentCreate instantiates a new AxervePaymentCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAxervePaymentCreate(data AxervePaymentCreateData) *AxervePaymentCreate {
	this := AxervePaymentCreate{}
	this.Data = data
	return &this
}

// NewAxervePaymentCreateWithDefaults instantiates a new AxervePaymentCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAxervePaymentCreateWithDefaults() *AxervePaymentCreate {
	this := AxervePaymentCreate{}
	return &this
}

// GetData returns the Data field value
func (o *AxervePaymentCreate) GetData() AxervePaymentCreateData {
	if o == nil {
		var ret AxervePaymentCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AxervePaymentCreate) GetDataOk() (*AxervePaymentCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AxervePaymentCreate) SetData(v AxervePaymentCreateData) {
	o.Data = v
}

func (o AxervePaymentCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AxervePaymentCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAxervePaymentCreate struct {
	value *AxervePaymentCreate
	isSet bool
}

func (v NullableAxervePaymentCreate) Get() *AxervePaymentCreate {
	return v.value
}

func (v *NullableAxervePaymentCreate) Set(val *AxervePaymentCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAxervePaymentCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAxervePaymentCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAxervePaymentCreate(val *AxervePaymentCreate) *NullableAxervePaymentCreate {
	return &NullableAxervePaymentCreate{value: val, isSet: true}
}

func (v NullableAxervePaymentCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAxervePaymentCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
