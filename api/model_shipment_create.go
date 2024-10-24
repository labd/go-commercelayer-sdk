/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ShipmentCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentCreate{}

// ShipmentCreate struct for ShipmentCreate
type ShipmentCreate struct {
	Data ShipmentCreateData `json:"data"`
}

// NewShipmentCreate instantiates a new ShipmentCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentCreate(data ShipmentCreateData) *ShipmentCreate {
	this := ShipmentCreate{}
	this.Data = data
	return &this
}

// NewShipmentCreateWithDefaults instantiates a new ShipmentCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentCreateWithDefaults() *ShipmentCreate {
	this := ShipmentCreate{}
	return &this
}

// GetData returns the Data field value
func (o *ShipmentCreate) GetData() ShipmentCreateData {
	if o == nil {
		var ret ShipmentCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ShipmentCreate) GetDataOk() (*ShipmentCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ShipmentCreate) SetData(v ShipmentCreateData) {
	o.Data = v
}

func (o ShipmentCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableShipmentCreate struct {
	value *ShipmentCreate
	isSet bool
}

func (v NullableShipmentCreate) Get() *ShipmentCreate {
	return v.value
}

func (v *NullableShipmentCreate) Set(val *ShipmentCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentCreate(val *ShipmentCreate) *NullableShipmentCreate {
	return &NullableShipmentCreate{value: val, isSet: true}
}

func (v NullableShipmentCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}