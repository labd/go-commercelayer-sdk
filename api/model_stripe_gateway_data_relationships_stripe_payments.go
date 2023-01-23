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

// StripeGatewayDataRelationshipsStripePayments struct for StripeGatewayDataRelationshipsStripePayments
type StripeGatewayDataRelationshipsStripePayments struct {
	Data *StripeGatewayDataRelationshipsStripePaymentsData `json:"data,omitempty"`
}

// NewStripeGatewayDataRelationshipsStripePayments instantiates a new StripeGatewayDataRelationshipsStripePayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeGatewayDataRelationshipsStripePayments() *StripeGatewayDataRelationshipsStripePayments {
	this := StripeGatewayDataRelationshipsStripePayments{}
	return &this
}

// NewStripeGatewayDataRelationshipsStripePaymentsWithDefaults instantiates a new StripeGatewayDataRelationshipsStripePayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeGatewayDataRelationshipsStripePaymentsWithDefaults() *StripeGatewayDataRelationshipsStripePayments {
	this := StripeGatewayDataRelationshipsStripePayments{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *StripeGatewayDataRelationshipsStripePayments) GetData() StripeGatewayDataRelationshipsStripePaymentsData {
	if o == nil || o.Data == nil {
		var ret StripeGatewayDataRelationshipsStripePaymentsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeGatewayDataRelationshipsStripePayments) GetDataOk() (*StripeGatewayDataRelationshipsStripePaymentsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StripeGatewayDataRelationshipsStripePayments) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given StripeGatewayDataRelationshipsStripePaymentsData and assigns it to the Data field.
func (o *StripeGatewayDataRelationshipsStripePayments) SetData(v StripeGatewayDataRelationshipsStripePaymentsData) {
	o.Data = &v
}

func (o StripeGatewayDataRelationshipsStripePayments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStripeGatewayDataRelationshipsStripePayments struct {
	value *StripeGatewayDataRelationshipsStripePayments
	isSet bool
}

func (v NullableStripeGatewayDataRelationshipsStripePayments) Get() *StripeGatewayDataRelationshipsStripePayments {
	return v.value
}

func (v *NullableStripeGatewayDataRelationshipsStripePayments) Set(val *StripeGatewayDataRelationshipsStripePayments) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeGatewayDataRelationshipsStripePayments) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeGatewayDataRelationshipsStripePayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeGatewayDataRelationshipsStripePayments(val *StripeGatewayDataRelationshipsStripePayments) *NullableStripeGatewayDataRelationshipsStripePayments {
	return &NullableStripeGatewayDataRelationshipsStripePayments{value: val, isSet: true}
}

func (v NullableStripeGatewayDataRelationshipsStripePayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeGatewayDataRelationshipsStripePayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
