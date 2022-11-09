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

// CustomerDataRelationshipsCustomerGroup struct for CustomerDataRelationshipsCustomerGroup
type CustomerDataRelationshipsCustomerGroup struct {
	Data CustomerDataRelationshipsCustomerGroupData `json:"data"`
}

// NewCustomerDataRelationshipsCustomerGroup instantiates a new CustomerDataRelationshipsCustomerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerDataRelationshipsCustomerGroup(data CustomerDataRelationshipsCustomerGroupData) *CustomerDataRelationshipsCustomerGroup {
	this := CustomerDataRelationshipsCustomerGroup{}
	this.Data = data
	return &this
}

// NewCustomerDataRelationshipsCustomerGroupWithDefaults instantiates a new CustomerDataRelationshipsCustomerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerDataRelationshipsCustomerGroupWithDefaults() *CustomerDataRelationshipsCustomerGroup {
	this := CustomerDataRelationshipsCustomerGroup{}
	return &this
}

// GetData returns the Data field value
func (o *CustomerDataRelationshipsCustomerGroup) GetData() CustomerDataRelationshipsCustomerGroupData {
	if o == nil {
		var ret CustomerDataRelationshipsCustomerGroupData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationshipsCustomerGroup) GetDataOk() (*CustomerDataRelationshipsCustomerGroupData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomerDataRelationshipsCustomerGroup) SetData(v CustomerDataRelationshipsCustomerGroupData) {
	o.Data = v
}

func (o CustomerDataRelationshipsCustomerGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerDataRelationshipsCustomerGroup struct {
	value *CustomerDataRelationshipsCustomerGroup
	isSet bool
}

func (v NullableCustomerDataRelationshipsCustomerGroup) Get() *CustomerDataRelationshipsCustomerGroup {
	return v.value
}

func (v *NullableCustomerDataRelationshipsCustomerGroup) Set(val *CustomerDataRelationshipsCustomerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerDataRelationshipsCustomerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerDataRelationshipsCustomerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerDataRelationshipsCustomerGroup(val *CustomerDataRelationshipsCustomerGroup) *NullableCustomerDataRelationshipsCustomerGroup {
	return &NullableCustomerDataRelationshipsCustomerGroup{value: val, isSet: true}
}

func (v NullableCustomerDataRelationshipsCustomerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerDataRelationshipsCustomerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
