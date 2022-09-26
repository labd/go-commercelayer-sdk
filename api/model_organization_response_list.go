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

// OrganizationResponseList struct for OrganizationResponseList
type OrganizationResponseList struct {
	Data []Data `json:"data,omitempty"`
}

// NewOrganizationResponseList instantiates a new OrganizationResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationResponseList() *OrganizationResponseList {
	this := OrganizationResponseList{}
	return &this
}

// NewOrganizationResponseListWithDefaults instantiates a new OrganizationResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationResponseListWithDefaults() *OrganizationResponseList {
	this := OrganizationResponseList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrganizationResponseList) GetData() []Data {
	if o == nil || o.Data == nil {
		var ret []Data
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationResponseList) GetDataOk() ([]Data, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrganizationResponseList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Data and assigns it to the Data field.
func (o *OrganizationResponseList) SetData(v []Data) {
	o.Data = v
}

func (o OrganizationResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationResponseList struct {
	value *OrganizationResponseList
	isSet bool
}

func (v NullableOrganizationResponseList) Get() *OrganizationResponseList {
	return v.value
}

func (v *NullableOrganizationResponseList) Set(val *OrganizationResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationResponseList(val *OrganizationResponseList) *NullableOrganizationResponseList {
	return &NullableOrganizationResponseList{value: val, isSet: true}
}

func (v NullableOrganizationResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
