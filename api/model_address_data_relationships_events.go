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

// checks if the AddressDataRelationshipsEvents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressDataRelationshipsEvents{}

// AddressDataRelationshipsEvents struct for AddressDataRelationshipsEvents
type AddressDataRelationshipsEvents struct {
	Data *AddressDataRelationshipsEventsData `json:"data,omitempty"`
}

// NewAddressDataRelationshipsEvents instantiates a new AddressDataRelationshipsEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressDataRelationshipsEvents() *AddressDataRelationshipsEvents {
	this := AddressDataRelationshipsEvents{}
	return &this
}

// NewAddressDataRelationshipsEventsWithDefaults instantiates a new AddressDataRelationshipsEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressDataRelationshipsEventsWithDefaults() *AddressDataRelationshipsEvents {
	this := AddressDataRelationshipsEvents{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AddressDataRelationshipsEvents) GetData() AddressDataRelationshipsEventsData {
	if o == nil || IsNil(o.Data) {
		var ret AddressDataRelationshipsEventsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataRelationshipsEvents) GetDataOk() (*AddressDataRelationshipsEventsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AddressDataRelationshipsEvents) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AddressDataRelationshipsEventsData and assigns it to the Data field.
func (o *AddressDataRelationshipsEvents) SetData(v AddressDataRelationshipsEventsData) {
	o.Data = &v
}

func (o AddressDataRelationshipsEvents) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressDataRelationshipsEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAddressDataRelationshipsEvents struct {
	value *AddressDataRelationshipsEvents
	isSet bool
}

func (v NullableAddressDataRelationshipsEvents) Get() *AddressDataRelationshipsEvents {
	return v.value
}

func (v *NullableAddressDataRelationshipsEvents) Set(val *AddressDataRelationshipsEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressDataRelationshipsEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressDataRelationshipsEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressDataRelationshipsEvents(val *AddressDataRelationshipsEvents) *NullableAddressDataRelationshipsEvents {
	return &NullableAddressDataRelationshipsEvents{value: val, isSet: true}
}

func (v NullableAddressDataRelationshipsEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressDataRelationshipsEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
