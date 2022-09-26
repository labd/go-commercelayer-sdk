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

// PaymentGatewayResponse struct for PaymentGatewayResponse
type PaymentGatewayResponse struct {
	Data *PaymentGatewayResponseData `json:"data,omitempty"`
}

// NewPaymentGatewayResponse instantiates a new PaymentGatewayResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentGatewayResponse() *PaymentGatewayResponse {
	this := PaymentGatewayResponse{}
	return &this
}

// NewPaymentGatewayResponseWithDefaults instantiates a new PaymentGatewayResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentGatewayResponseWithDefaults() *PaymentGatewayResponse {
	this := PaymentGatewayResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PaymentGatewayResponse) GetData() PaymentGatewayResponseData {
	if o == nil || o.Data == nil {
		var ret PaymentGatewayResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayResponse) GetDataOk() (*PaymentGatewayResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PaymentGatewayResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PaymentGatewayResponseData and assigns it to the Data field.
func (o *PaymentGatewayResponse) SetData(v PaymentGatewayResponseData) {
	o.Data = &v
}

func (o PaymentGatewayResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentGatewayResponse struct {
	value *PaymentGatewayResponse
	isSet bool
}

func (v NullablePaymentGatewayResponse) Get() *PaymentGatewayResponse {
	return v.value
}

func (v *NullablePaymentGatewayResponse) Set(val *PaymentGatewayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentGatewayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentGatewayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentGatewayResponse(val *PaymentGatewayResponse) *NullablePaymentGatewayResponse {
	return &NullablePaymentGatewayResponse{value: val, isSet: true}
}

func (v NullablePaymentGatewayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentGatewayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
