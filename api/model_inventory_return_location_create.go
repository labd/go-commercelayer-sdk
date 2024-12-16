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

// checks if the InventoryReturnLocationCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryReturnLocationCreate{}

// InventoryReturnLocationCreate struct for InventoryReturnLocationCreate
type InventoryReturnLocationCreate struct {
	Data InventoryReturnLocationCreateData `json:"data"`
}

// NewInventoryReturnLocationCreate instantiates a new InventoryReturnLocationCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryReturnLocationCreate(data InventoryReturnLocationCreateData) *InventoryReturnLocationCreate {
	this := InventoryReturnLocationCreate{}
	this.Data = data
	return &this
}

// NewInventoryReturnLocationCreateWithDefaults instantiates a new InventoryReturnLocationCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryReturnLocationCreateWithDefaults() *InventoryReturnLocationCreate {
	this := InventoryReturnLocationCreate{}
	return &this
}

// GetData returns the Data field value
func (o *InventoryReturnLocationCreate) GetData() InventoryReturnLocationCreateData {
	if o == nil {
		var ret InventoryReturnLocationCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationCreate) GetDataOk() (*InventoryReturnLocationCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InventoryReturnLocationCreate) SetData(v InventoryReturnLocationCreateData) {
	o.Data = v
}

func (o InventoryReturnLocationCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryReturnLocationCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableInventoryReturnLocationCreate struct {
	value *InventoryReturnLocationCreate
	isSet bool
}

func (v NullableInventoryReturnLocationCreate) Get() *InventoryReturnLocationCreate {
	return v.value
}

func (v *NullableInventoryReturnLocationCreate) Set(val *InventoryReturnLocationCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryReturnLocationCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryReturnLocationCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryReturnLocationCreate(val *InventoryReturnLocationCreate) *NullableInventoryReturnLocationCreate {
	return &NullableInventoryReturnLocationCreate{value: val, isSet: true}
}

func (v NullableInventoryReturnLocationCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryReturnLocationCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
