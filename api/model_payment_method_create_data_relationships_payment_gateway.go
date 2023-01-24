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

// PaymentMethodCreateDataRelationshipsPaymentGateway struct for PaymentMethodCreateDataRelationshipsPaymentGateway
type PaymentMethodCreateDataRelationshipsPaymentGateway struct {
	Data AdyenPaymentDataRelationshipsPaymentGatewayData `json:"data"`
}

// NewPaymentMethodCreateDataRelationshipsPaymentGateway instantiates a new PaymentMethodCreateDataRelationshipsPaymentGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodCreateDataRelationshipsPaymentGateway(data AdyenPaymentDataRelationshipsPaymentGatewayData) *PaymentMethodCreateDataRelationshipsPaymentGateway {
	this := PaymentMethodCreateDataRelationshipsPaymentGateway{}
	this.Data = data
	return &this
}

// NewPaymentMethodCreateDataRelationshipsPaymentGatewayWithDefaults instantiates a new PaymentMethodCreateDataRelationshipsPaymentGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodCreateDataRelationshipsPaymentGatewayWithDefaults() *PaymentMethodCreateDataRelationshipsPaymentGateway {
	this := PaymentMethodCreateDataRelationshipsPaymentGateway{}
	return &this
}

// GetData returns the Data field value
func (o *PaymentMethodCreateDataRelationshipsPaymentGateway) GetData() AdyenPaymentDataRelationshipsPaymentGatewayData {
	if o == nil {
		var ret AdyenPaymentDataRelationshipsPaymentGatewayData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCreateDataRelationshipsPaymentGateway) GetDataOk() (*AdyenPaymentDataRelationshipsPaymentGatewayData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PaymentMethodCreateDataRelationshipsPaymentGateway) SetData(v AdyenPaymentDataRelationshipsPaymentGatewayData) {
	o.Data = v
}

func (o PaymentMethodCreateDataRelationshipsPaymentGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethodCreateDataRelationshipsPaymentGateway struct {
	value *PaymentMethodCreateDataRelationshipsPaymentGateway
	isSet bool
}

func (v NullablePaymentMethodCreateDataRelationshipsPaymentGateway) Get() *PaymentMethodCreateDataRelationshipsPaymentGateway {
	return v.value
}

func (v *NullablePaymentMethodCreateDataRelationshipsPaymentGateway) Set(val *PaymentMethodCreateDataRelationshipsPaymentGateway) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCreateDataRelationshipsPaymentGateway) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCreateDataRelationshipsPaymentGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCreateDataRelationshipsPaymentGateway(val *PaymentMethodCreateDataRelationshipsPaymentGateway) *NullablePaymentMethodCreateDataRelationshipsPaymentGateway {
	return &NullablePaymentMethodCreateDataRelationshipsPaymentGateway{value: val, isSet: true}
}

func (v NullablePaymentMethodCreateDataRelationshipsPaymentGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCreateDataRelationshipsPaymentGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}