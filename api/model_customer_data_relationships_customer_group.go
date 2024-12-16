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

// checks if the CustomerDataRelationshipsCustomerGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerDataRelationshipsCustomerGroup{}

// CustomerDataRelationshipsCustomerGroup struct for CustomerDataRelationshipsCustomerGroup
type CustomerDataRelationshipsCustomerGroup struct {
	Data *CustomerDataRelationshipsCustomerGroupData `json:"data,omitempty"`
}

// NewCustomerDataRelationshipsCustomerGroup instantiates a new CustomerDataRelationshipsCustomerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerDataRelationshipsCustomerGroup() *CustomerDataRelationshipsCustomerGroup {
	this := CustomerDataRelationshipsCustomerGroup{}
	return &this
}

// NewCustomerDataRelationshipsCustomerGroupWithDefaults instantiates a new CustomerDataRelationshipsCustomerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerDataRelationshipsCustomerGroupWithDefaults() *CustomerDataRelationshipsCustomerGroup {
	this := CustomerDataRelationshipsCustomerGroup{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomerDataRelationshipsCustomerGroup) GetData() CustomerDataRelationshipsCustomerGroupData {
	if o == nil || IsNil(o.Data) {
		var ret CustomerDataRelationshipsCustomerGroupData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationshipsCustomerGroup) GetDataOk() (*CustomerDataRelationshipsCustomerGroupData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomerDataRelationshipsCustomerGroup) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CustomerDataRelationshipsCustomerGroupData and assigns it to the Data field.
func (o *CustomerDataRelationshipsCustomerGroup) SetData(v CustomerDataRelationshipsCustomerGroupData) {
	o.Data = &v
}

func (o CustomerDataRelationshipsCustomerGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerDataRelationshipsCustomerGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
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
