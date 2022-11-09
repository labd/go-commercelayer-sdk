/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETTransactions200ResponseDataInnerRelationships struct for GETTransactions200ResponseDataInnerRelationships
type GETTransactions200ResponseDataInnerRelationships struct {
	Order *GETAdyenPayments200ResponseDataInnerRelationshipsOrder `json:"order,omitempty"`
}

// NewGETTransactions200ResponseDataInnerRelationships instantiates a new GETTransactions200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETTransactions200ResponseDataInnerRelationships() *GETTransactions200ResponseDataInnerRelationships {
	this := GETTransactions200ResponseDataInnerRelationships{}
	return &this
}

// NewGETTransactions200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETTransactions200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETTransactions200ResponseDataInnerRelationshipsWithDefaults() *GETTransactions200ResponseDataInnerRelationships {
	this := GETTransactions200ResponseDataInnerRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *GETTransactions200ResponseDataInnerRelationships) GetOrder() GETAdyenPayments200ResponseDataInnerRelationshipsOrder {
	if o == nil || o.Order == nil {
		var ret GETAdyenPayments200ResponseDataInnerRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTransactions200ResponseDataInnerRelationships) GetOrderOk() (*GETAdyenPayments200ResponseDataInnerRelationshipsOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *GETTransactions200ResponseDataInnerRelationships) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given GETAdyenPayments200ResponseDataInnerRelationshipsOrder and assigns it to the Order field.
func (o *GETTransactions200ResponseDataInnerRelationships) SetOrder(v GETAdyenPayments200ResponseDataInnerRelationshipsOrder) {
	o.Order = &v
}

func (o GETTransactions200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	return json.Marshal(toSerialize)
}

type NullableGETTransactions200ResponseDataInnerRelationships struct {
	value *GETTransactions200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETTransactions200ResponseDataInnerRelationships) Get() *GETTransactions200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETTransactions200ResponseDataInnerRelationships) Set(val *GETTransactions200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETTransactions200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETTransactions200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETTransactions200ResponseDataInnerRelationships(val *GETTransactions200ResponseDataInnerRelationships) *NullableGETTransactions200ResponseDataInnerRelationships {
	return &NullableGETTransactions200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETTransactions200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETTransactions200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
