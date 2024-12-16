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

// checks if the CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments{}

// CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments struct for CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments
type CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments struct {
	Data CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData `json:"data"`
}

// NewCheckoutComGatewayCreateDataRelationshipsCheckoutComPayments instantiates a new CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComGatewayCreateDataRelationshipsCheckoutComPayments(data CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) *CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments {
	this := CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments{}
	this.Data = data
	return &this
}

// NewCheckoutComGatewayCreateDataRelationshipsCheckoutComPaymentsWithDefaults instantiates a new CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComGatewayCreateDataRelationshipsCheckoutComPaymentsWithDefaults() *CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments {
	this := CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments{}
	return &this
}

// GetData returns the Data field value
func (o *CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments) GetData() CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData {
	if o == nil {
		var ret CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments) GetDataOk() (*CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments) SetData(v CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) {
	o.Data = v
}

func (o CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCheckoutComGatewayCreateDataRelationshipsCheckoutComPayments struct {
	value *CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments
	isSet bool
}

func (v NullableCheckoutComGatewayCreateDataRelationshipsCheckoutComPayments) Get() *CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments {
	return v.value
}

func (v *NullableCheckoutComGatewayCreateDataRelationshipsCheckoutComPayments) Set(val *CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComGatewayCreateDataRelationshipsCheckoutComPayments) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComGatewayCreateDataRelationshipsCheckoutComPayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComGatewayCreateDataRelationshipsCheckoutComPayments(val *CheckoutComGatewayCreateDataRelationshipsCheckoutComPayments) *NullableCheckoutComGatewayCreateDataRelationshipsCheckoutComPayments {
	return &NullableCheckoutComGatewayCreateDataRelationshipsCheckoutComPayments{value: val, isSet: true}
}

func (v NullableCheckoutComGatewayCreateDataRelationshipsCheckoutComPayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComGatewayCreateDataRelationshipsCheckoutComPayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
