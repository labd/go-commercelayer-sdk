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

// checks if the ShipmentUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentUpdate{}

// ShipmentUpdate struct for ShipmentUpdate
type ShipmentUpdate struct {
	Data ShipmentUpdateData `json:"data"`
}

// NewShipmentUpdate instantiates a new ShipmentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentUpdate(data ShipmentUpdateData) *ShipmentUpdate {
	this := ShipmentUpdate{}
	this.Data = data
	return &this
}

// NewShipmentUpdateWithDefaults instantiates a new ShipmentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentUpdateWithDefaults() *ShipmentUpdate {
	this := ShipmentUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *ShipmentUpdate) GetData() ShipmentUpdateData {
	if o == nil {
		var ret ShipmentUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ShipmentUpdate) GetDataOk() (*ShipmentUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ShipmentUpdate) SetData(v ShipmentUpdateData) {
	o.Data = v
}

func (o ShipmentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableShipmentUpdate struct {
	value *ShipmentUpdate
	isSet bool
}

func (v NullableShipmentUpdate) Get() *ShipmentUpdate {
	return v.value
}

func (v *NullableShipmentUpdate) Set(val *ShipmentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentUpdate(val *ShipmentUpdate) *NullableShipmentUpdate {
	return &NullableShipmentUpdate{value: val, isSet: true}
}

func (v NullableShipmentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
