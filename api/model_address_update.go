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

// checks if the AddressUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressUpdate{}

// AddressUpdate struct for AddressUpdate
type AddressUpdate struct {
	Data AddressUpdateData `json:"data"`
}

// NewAddressUpdate instantiates a new AddressUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressUpdate(data AddressUpdateData) *AddressUpdate {
	this := AddressUpdate{}
	this.Data = data
	return &this
}

// NewAddressUpdateWithDefaults instantiates a new AddressUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressUpdateWithDefaults() *AddressUpdate {
	this := AddressUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *AddressUpdate) GetData() AddressUpdateData {
	if o == nil {
		var ret AddressUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddressUpdate) GetDataOk() (*AddressUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AddressUpdate) SetData(v AddressUpdateData) {
	o.Data = v
}

func (o AddressUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAddressUpdate struct {
	value *AddressUpdate
	isSet bool
}

func (v NullableAddressUpdate) Get() *AddressUpdate {
	return v.value
}

func (v *NullableAddressUpdate) Set(val *AddressUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressUpdate(val *AddressUpdate) *NullableAddressUpdate {
	return &NullableAddressUpdate{value: val, isSet: true}
}

func (v NullableAddressUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
