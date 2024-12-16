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

// checks if the ShippingMethodDataRelationshipsShippingWeightTiers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingMethodDataRelationshipsShippingWeightTiers{}

// ShippingMethodDataRelationshipsShippingWeightTiers struct for ShippingMethodDataRelationshipsShippingWeightTiers
type ShippingMethodDataRelationshipsShippingWeightTiers struct {
	Data *ShippingMethodDataRelationshipsShippingWeightTiersData `json:"data,omitempty"`
}

// NewShippingMethodDataRelationshipsShippingWeightTiers instantiates a new ShippingMethodDataRelationshipsShippingWeightTiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodDataRelationshipsShippingWeightTiers() *ShippingMethodDataRelationshipsShippingWeightTiers {
	this := ShippingMethodDataRelationshipsShippingWeightTiers{}
	return &this
}

// NewShippingMethodDataRelationshipsShippingWeightTiersWithDefaults instantiates a new ShippingMethodDataRelationshipsShippingWeightTiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodDataRelationshipsShippingWeightTiersWithDefaults() *ShippingMethodDataRelationshipsShippingWeightTiers {
	this := ShippingMethodDataRelationshipsShippingWeightTiers{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ShippingMethodDataRelationshipsShippingWeightTiers) GetData() ShippingMethodDataRelationshipsShippingWeightTiersData {
	if o == nil || IsNil(o.Data) {
		var ret ShippingMethodDataRelationshipsShippingWeightTiersData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataRelationshipsShippingWeightTiers) GetDataOk() (*ShippingMethodDataRelationshipsShippingWeightTiersData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ShippingMethodDataRelationshipsShippingWeightTiers) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ShippingMethodDataRelationshipsShippingWeightTiersData and assigns it to the Data field.
func (o *ShippingMethodDataRelationshipsShippingWeightTiers) SetData(v ShippingMethodDataRelationshipsShippingWeightTiersData) {
	o.Data = &v
}

func (o ShippingMethodDataRelationshipsShippingWeightTiers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingMethodDataRelationshipsShippingWeightTiers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableShippingMethodDataRelationshipsShippingWeightTiers struct {
	value *ShippingMethodDataRelationshipsShippingWeightTiers
	isSet bool
}

func (v NullableShippingMethodDataRelationshipsShippingWeightTiers) Get() *ShippingMethodDataRelationshipsShippingWeightTiers {
	return v.value
}

func (v *NullableShippingMethodDataRelationshipsShippingWeightTiers) Set(val *ShippingMethodDataRelationshipsShippingWeightTiers) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodDataRelationshipsShippingWeightTiers) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodDataRelationshipsShippingWeightTiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodDataRelationshipsShippingWeightTiers(val *ShippingMethodDataRelationshipsShippingWeightTiers) *NullableShippingMethodDataRelationshipsShippingWeightTiers {
	return &NullableShippingMethodDataRelationshipsShippingWeightTiers{value: val, isSet: true}
}

func (v NullableShippingMethodDataRelationshipsShippingWeightTiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodDataRelationshipsShippingWeightTiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
