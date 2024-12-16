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

// checks if the ShippingZoneCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingZoneCreate{}

// ShippingZoneCreate struct for ShippingZoneCreate
type ShippingZoneCreate struct {
	Data ShippingZoneCreateData `json:"data"`
}

// NewShippingZoneCreate instantiates a new ShippingZoneCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingZoneCreate(data ShippingZoneCreateData) *ShippingZoneCreate {
	this := ShippingZoneCreate{}
	this.Data = data
	return &this
}

// NewShippingZoneCreateWithDefaults instantiates a new ShippingZoneCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingZoneCreateWithDefaults() *ShippingZoneCreate {
	this := ShippingZoneCreate{}
	return &this
}

// GetData returns the Data field value
func (o *ShippingZoneCreate) GetData() ShippingZoneCreateData {
	if o == nil {
		var ret ShippingZoneCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ShippingZoneCreate) GetDataOk() (*ShippingZoneCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ShippingZoneCreate) SetData(v ShippingZoneCreateData) {
	o.Data = v
}

func (o ShippingZoneCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingZoneCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableShippingZoneCreate struct {
	value *ShippingZoneCreate
	isSet bool
}

func (v NullableShippingZoneCreate) Get() *ShippingZoneCreate {
	return v.value
}

func (v *NullableShippingZoneCreate) Set(val *ShippingZoneCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingZoneCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingZoneCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingZoneCreate(val *ShippingZoneCreate) *NullableShippingZoneCreate {
	return &NullableShippingZoneCreate{value: val, isSet: true}
}

func (v NullableShippingZoneCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingZoneCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
