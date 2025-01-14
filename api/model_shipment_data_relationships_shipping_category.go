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

// checks if the ShipmentDataRelationshipsShippingCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentDataRelationshipsShippingCategory{}

// ShipmentDataRelationshipsShippingCategory struct for ShipmentDataRelationshipsShippingCategory
type ShipmentDataRelationshipsShippingCategory struct {
	Data *ShipmentDataRelationshipsShippingCategoryData `json:"data,omitempty"`
}

// NewShipmentDataRelationshipsShippingCategory instantiates a new ShipmentDataRelationshipsShippingCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentDataRelationshipsShippingCategory() *ShipmentDataRelationshipsShippingCategory {
	this := ShipmentDataRelationshipsShippingCategory{}
	return &this
}

// NewShipmentDataRelationshipsShippingCategoryWithDefaults instantiates a new ShipmentDataRelationshipsShippingCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentDataRelationshipsShippingCategoryWithDefaults() *ShipmentDataRelationshipsShippingCategory {
	this := ShipmentDataRelationshipsShippingCategory{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ShipmentDataRelationshipsShippingCategory) GetData() ShipmentDataRelationshipsShippingCategoryData {
	if o == nil || IsNil(o.Data) {
		var ret ShipmentDataRelationshipsShippingCategoryData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationshipsShippingCategory) GetDataOk() (*ShipmentDataRelationshipsShippingCategoryData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ShipmentDataRelationshipsShippingCategory) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ShipmentDataRelationshipsShippingCategoryData and assigns it to the Data field.
func (o *ShipmentDataRelationshipsShippingCategory) SetData(v ShipmentDataRelationshipsShippingCategoryData) {
	o.Data = &v
}

func (o ShipmentDataRelationshipsShippingCategory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentDataRelationshipsShippingCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableShipmentDataRelationshipsShippingCategory struct {
	value *ShipmentDataRelationshipsShippingCategory
	isSet bool
}

func (v NullableShipmentDataRelationshipsShippingCategory) Get() *ShipmentDataRelationshipsShippingCategory {
	return v.value
}

func (v *NullableShipmentDataRelationshipsShippingCategory) Set(val *ShipmentDataRelationshipsShippingCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentDataRelationshipsShippingCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentDataRelationshipsShippingCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentDataRelationshipsShippingCategory(val *ShipmentDataRelationshipsShippingCategory) *NullableShipmentDataRelationshipsShippingCategory {
	return &NullableShipmentDataRelationshipsShippingCategory{value: val, isSet: true}
}

func (v NullableShipmentDataRelationshipsShippingCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentDataRelationshipsShippingCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
