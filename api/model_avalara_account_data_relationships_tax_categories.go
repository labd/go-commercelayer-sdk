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

// checks if the AvalaraAccountDataRelationshipsTaxCategories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvalaraAccountDataRelationshipsTaxCategories{}

// AvalaraAccountDataRelationshipsTaxCategories struct for AvalaraAccountDataRelationshipsTaxCategories
type AvalaraAccountDataRelationshipsTaxCategories struct {
	Data *AvalaraAccountDataRelationshipsTaxCategoriesData `json:"data,omitempty"`
}

// NewAvalaraAccountDataRelationshipsTaxCategories instantiates a new AvalaraAccountDataRelationshipsTaxCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvalaraAccountDataRelationshipsTaxCategories() *AvalaraAccountDataRelationshipsTaxCategories {
	this := AvalaraAccountDataRelationshipsTaxCategories{}
	return &this
}

// NewAvalaraAccountDataRelationshipsTaxCategoriesWithDefaults instantiates a new AvalaraAccountDataRelationshipsTaxCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvalaraAccountDataRelationshipsTaxCategoriesWithDefaults() *AvalaraAccountDataRelationshipsTaxCategories {
	this := AvalaraAccountDataRelationshipsTaxCategories{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AvalaraAccountDataRelationshipsTaxCategories) GetData() AvalaraAccountDataRelationshipsTaxCategoriesData {
	if o == nil || IsNil(o.Data) {
		var ret AvalaraAccountDataRelationshipsTaxCategoriesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountDataRelationshipsTaxCategories) GetDataOk() (*AvalaraAccountDataRelationshipsTaxCategoriesData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AvalaraAccountDataRelationshipsTaxCategories) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AvalaraAccountDataRelationshipsTaxCategoriesData and assigns it to the Data field.
func (o *AvalaraAccountDataRelationshipsTaxCategories) SetData(v AvalaraAccountDataRelationshipsTaxCategoriesData) {
	o.Data = &v
}

func (o AvalaraAccountDataRelationshipsTaxCategories) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvalaraAccountDataRelationshipsTaxCategories) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAvalaraAccountDataRelationshipsTaxCategories struct {
	value *AvalaraAccountDataRelationshipsTaxCategories
	isSet bool
}

func (v NullableAvalaraAccountDataRelationshipsTaxCategories) Get() *AvalaraAccountDataRelationshipsTaxCategories {
	return v.value
}

func (v *NullableAvalaraAccountDataRelationshipsTaxCategories) Set(val *AvalaraAccountDataRelationshipsTaxCategories) {
	v.value = val
	v.isSet = true
}

func (v NullableAvalaraAccountDataRelationshipsTaxCategories) IsSet() bool {
	return v.isSet
}

func (v *NullableAvalaraAccountDataRelationshipsTaxCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvalaraAccountDataRelationshipsTaxCategories(val *AvalaraAccountDataRelationshipsTaxCategories) *NullableAvalaraAccountDataRelationshipsTaxCategories {
	return &NullableAvalaraAccountDataRelationshipsTaxCategories{value: val, isSet: true}
}

func (v NullableAvalaraAccountDataRelationshipsTaxCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvalaraAccountDataRelationshipsTaxCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
