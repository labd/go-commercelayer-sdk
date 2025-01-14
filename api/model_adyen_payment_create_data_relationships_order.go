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

// checks if the AdyenPaymentCreateDataRelationshipsOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdyenPaymentCreateDataRelationshipsOrder{}

// AdyenPaymentCreateDataRelationshipsOrder struct for AdyenPaymentCreateDataRelationshipsOrder
type AdyenPaymentCreateDataRelationshipsOrder struct {
	Data AdyenPaymentDataRelationshipsOrderData `json:"data"`
}

// NewAdyenPaymentCreateDataRelationshipsOrder instantiates a new AdyenPaymentCreateDataRelationshipsOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenPaymentCreateDataRelationshipsOrder(data AdyenPaymentDataRelationshipsOrderData) *AdyenPaymentCreateDataRelationshipsOrder {
	this := AdyenPaymentCreateDataRelationshipsOrder{}
	this.Data = data
	return &this
}

// NewAdyenPaymentCreateDataRelationshipsOrderWithDefaults instantiates a new AdyenPaymentCreateDataRelationshipsOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenPaymentCreateDataRelationshipsOrderWithDefaults() *AdyenPaymentCreateDataRelationshipsOrder {
	this := AdyenPaymentCreateDataRelationshipsOrder{}
	return &this
}

// GetData returns the Data field value
func (o *AdyenPaymentCreateDataRelationshipsOrder) GetData() AdyenPaymentDataRelationshipsOrderData {
	if o == nil {
		var ret AdyenPaymentDataRelationshipsOrderData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AdyenPaymentCreateDataRelationshipsOrder) GetDataOk() (*AdyenPaymentDataRelationshipsOrderData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AdyenPaymentCreateDataRelationshipsOrder) SetData(v AdyenPaymentDataRelationshipsOrderData) {
	o.Data = v
}

func (o AdyenPaymentCreateDataRelationshipsOrder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdyenPaymentCreateDataRelationshipsOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAdyenPaymentCreateDataRelationshipsOrder struct {
	value *AdyenPaymentCreateDataRelationshipsOrder
	isSet bool
}

func (v NullableAdyenPaymentCreateDataRelationshipsOrder) Get() *AdyenPaymentCreateDataRelationshipsOrder {
	return v.value
}

func (v *NullableAdyenPaymentCreateDataRelationshipsOrder) Set(val *AdyenPaymentCreateDataRelationshipsOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenPaymentCreateDataRelationshipsOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenPaymentCreateDataRelationshipsOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenPaymentCreateDataRelationshipsOrder(val *AdyenPaymentCreateDataRelationshipsOrder) *NullableAdyenPaymentCreateDataRelationshipsOrder {
	return &NullableAdyenPaymentCreateDataRelationshipsOrder{value: val, isSet: true}
}

func (v NullableAdyenPaymentCreateDataRelationshipsOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenPaymentCreateDataRelationshipsOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
