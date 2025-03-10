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

// checks if the InventoryReturnLocationUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryReturnLocationUpdate{}

// InventoryReturnLocationUpdate struct for InventoryReturnLocationUpdate
type InventoryReturnLocationUpdate struct {
	Data InventoryReturnLocationUpdateData `json:"data"`
}

// NewInventoryReturnLocationUpdate instantiates a new InventoryReturnLocationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryReturnLocationUpdate(data InventoryReturnLocationUpdateData) *InventoryReturnLocationUpdate {
	this := InventoryReturnLocationUpdate{}
	this.Data = data
	return &this
}

// NewInventoryReturnLocationUpdateWithDefaults instantiates a new InventoryReturnLocationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryReturnLocationUpdateWithDefaults() *InventoryReturnLocationUpdate {
	this := InventoryReturnLocationUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *InventoryReturnLocationUpdate) GetData() InventoryReturnLocationUpdateData {
	if o == nil {
		var ret InventoryReturnLocationUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationUpdate) GetDataOk() (*InventoryReturnLocationUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InventoryReturnLocationUpdate) SetData(v InventoryReturnLocationUpdateData) {
	o.Data = v
}

func (o InventoryReturnLocationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryReturnLocationUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableInventoryReturnLocationUpdate struct {
	value *InventoryReturnLocationUpdate
	isSet bool
}

func (v NullableInventoryReturnLocationUpdate) Get() *InventoryReturnLocationUpdate {
	return v.value
}

func (v *NullableInventoryReturnLocationUpdate) Set(val *InventoryReturnLocationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryReturnLocationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryReturnLocationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryReturnLocationUpdate(val *InventoryReturnLocationUpdate) *NullableInventoryReturnLocationUpdate {
	return &NullableInventoryReturnLocationUpdate{value: val, isSet: true}
}

func (v NullableInventoryReturnLocationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryReturnLocationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
