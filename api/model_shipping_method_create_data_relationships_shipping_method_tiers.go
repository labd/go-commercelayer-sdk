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

// checks if the ShippingMethodCreateDataRelationshipsShippingMethodTiers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingMethodCreateDataRelationshipsShippingMethodTiers{}

// ShippingMethodCreateDataRelationshipsShippingMethodTiers struct for ShippingMethodCreateDataRelationshipsShippingMethodTiers
type ShippingMethodCreateDataRelationshipsShippingMethodTiers struct {
	Data ShippingMethodDataRelationshipsShippingMethodTiersData `json:"data"`
}

// NewShippingMethodCreateDataRelationshipsShippingMethodTiers instantiates a new ShippingMethodCreateDataRelationshipsShippingMethodTiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodCreateDataRelationshipsShippingMethodTiers(data ShippingMethodDataRelationshipsShippingMethodTiersData) *ShippingMethodCreateDataRelationshipsShippingMethodTiers {
	this := ShippingMethodCreateDataRelationshipsShippingMethodTiers{}
	this.Data = data
	return &this
}

// NewShippingMethodCreateDataRelationshipsShippingMethodTiersWithDefaults instantiates a new ShippingMethodCreateDataRelationshipsShippingMethodTiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodCreateDataRelationshipsShippingMethodTiersWithDefaults() *ShippingMethodCreateDataRelationshipsShippingMethodTiers {
	this := ShippingMethodCreateDataRelationshipsShippingMethodTiers{}
	return &this
}

// GetData returns the Data field value
func (o *ShippingMethodCreateDataRelationshipsShippingMethodTiers) GetData() ShippingMethodDataRelationshipsShippingMethodTiersData {
	if o == nil {
		var ret ShippingMethodDataRelationshipsShippingMethodTiersData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ShippingMethodCreateDataRelationshipsShippingMethodTiers) GetDataOk() (*ShippingMethodDataRelationshipsShippingMethodTiersData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ShippingMethodCreateDataRelationshipsShippingMethodTiers) SetData(v ShippingMethodDataRelationshipsShippingMethodTiersData) {
	o.Data = v
}

func (o ShippingMethodCreateDataRelationshipsShippingMethodTiers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingMethodCreateDataRelationshipsShippingMethodTiers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableShippingMethodCreateDataRelationshipsShippingMethodTiers struct {
	value *ShippingMethodCreateDataRelationshipsShippingMethodTiers
	isSet bool
}

func (v NullableShippingMethodCreateDataRelationshipsShippingMethodTiers) Get() *ShippingMethodCreateDataRelationshipsShippingMethodTiers {
	return v.value
}

func (v *NullableShippingMethodCreateDataRelationshipsShippingMethodTiers) Set(val *ShippingMethodCreateDataRelationshipsShippingMethodTiers) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodCreateDataRelationshipsShippingMethodTiers) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodCreateDataRelationshipsShippingMethodTiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodCreateDataRelationshipsShippingMethodTiers(val *ShippingMethodCreateDataRelationshipsShippingMethodTiers) *NullableShippingMethodCreateDataRelationshipsShippingMethodTiers {
	return &NullableShippingMethodCreateDataRelationshipsShippingMethodTiers{value: val, isSet: true}
}

func (v NullableShippingMethodCreateDataRelationshipsShippingMethodTiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodCreateDataRelationshipsShippingMethodTiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
