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

// checks if the CustomerDataRelationshipsCustomerSubscriptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerDataRelationshipsCustomerSubscriptions{}

// CustomerDataRelationshipsCustomerSubscriptions struct for CustomerDataRelationshipsCustomerSubscriptions
type CustomerDataRelationshipsCustomerSubscriptions struct {
	Data *CustomerDataRelationshipsCustomerSubscriptionsData `json:"data,omitempty"`
}

// NewCustomerDataRelationshipsCustomerSubscriptions instantiates a new CustomerDataRelationshipsCustomerSubscriptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerDataRelationshipsCustomerSubscriptions() *CustomerDataRelationshipsCustomerSubscriptions {
	this := CustomerDataRelationshipsCustomerSubscriptions{}
	return &this
}

// NewCustomerDataRelationshipsCustomerSubscriptionsWithDefaults instantiates a new CustomerDataRelationshipsCustomerSubscriptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerDataRelationshipsCustomerSubscriptionsWithDefaults() *CustomerDataRelationshipsCustomerSubscriptions {
	this := CustomerDataRelationshipsCustomerSubscriptions{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomerDataRelationshipsCustomerSubscriptions) GetData() CustomerDataRelationshipsCustomerSubscriptionsData {
	if o == nil || IsNil(o.Data) {
		var ret CustomerDataRelationshipsCustomerSubscriptionsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationshipsCustomerSubscriptions) GetDataOk() (*CustomerDataRelationshipsCustomerSubscriptionsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomerDataRelationshipsCustomerSubscriptions) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CustomerDataRelationshipsCustomerSubscriptionsData and assigns it to the Data field.
func (o *CustomerDataRelationshipsCustomerSubscriptions) SetData(v CustomerDataRelationshipsCustomerSubscriptionsData) {
	o.Data = &v
}

func (o CustomerDataRelationshipsCustomerSubscriptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerDataRelationshipsCustomerSubscriptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCustomerDataRelationshipsCustomerSubscriptions struct {
	value *CustomerDataRelationshipsCustomerSubscriptions
	isSet bool
}

func (v NullableCustomerDataRelationshipsCustomerSubscriptions) Get() *CustomerDataRelationshipsCustomerSubscriptions {
	return v.value
}

func (v *NullableCustomerDataRelationshipsCustomerSubscriptions) Set(val *CustomerDataRelationshipsCustomerSubscriptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerDataRelationshipsCustomerSubscriptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerDataRelationshipsCustomerSubscriptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerDataRelationshipsCustomerSubscriptions(val *CustomerDataRelationshipsCustomerSubscriptions) *NullableCustomerDataRelationshipsCustomerSubscriptions {
	return &NullableCustomerDataRelationshipsCustomerSubscriptions{value: val, isSet: true}
}

func (v NullableCustomerDataRelationshipsCustomerSubscriptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerDataRelationshipsCustomerSubscriptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
