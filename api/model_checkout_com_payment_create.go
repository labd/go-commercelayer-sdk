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

// checks if the CheckoutComPaymentCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutComPaymentCreate{}

// CheckoutComPaymentCreate struct for CheckoutComPaymentCreate
type CheckoutComPaymentCreate struct {
	Data CheckoutComPaymentCreateData `json:"data"`
}

// NewCheckoutComPaymentCreate instantiates a new CheckoutComPaymentCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComPaymentCreate(data CheckoutComPaymentCreateData) *CheckoutComPaymentCreate {
	this := CheckoutComPaymentCreate{}
	this.Data = data
	return &this
}

// NewCheckoutComPaymentCreateWithDefaults instantiates a new CheckoutComPaymentCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComPaymentCreateWithDefaults() *CheckoutComPaymentCreate {
	this := CheckoutComPaymentCreate{}
	return &this
}

// GetData returns the Data field value
func (o *CheckoutComPaymentCreate) GetData() CheckoutComPaymentCreateData {
	if o == nil {
		var ret CheckoutComPaymentCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentCreate) GetDataOk() (*CheckoutComPaymentCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CheckoutComPaymentCreate) SetData(v CheckoutComPaymentCreateData) {
	o.Data = v
}

func (o CheckoutComPaymentCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutComPaymentCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCheckoutComPaymentCreate struct {
	value *CheckoutComPaymentCreate
	isSet bool
}

func (v NullableCheckoutComPaymentCreate) Get() *CheckoutComPaymentCreate {
	return v.value
}

func (v *NullableCheckoutComPaymentCreate) Set(val *CheckoutComPaymentCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComPaymentCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComPaymentCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComPaymentCreate(val *CheckoutComPaymentCreate) *NullableCheckoutComPaymentCreate {
	return &NullableCheckoutComPaymentCreate{value: val, isSet: true}
}

func (v NullableCheckoutComPaymentCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComPaymentCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
