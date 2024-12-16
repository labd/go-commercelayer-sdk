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

// checks if the AddressDataRelationshipsEventsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressDataRelationshipsEventsData{}

// AddressDataRelationshipsEventsData struct for AddressDataRelationshipsEventsData
type AddressDataRelationshipsEventsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource's id
	Id interface{} `json:"id,omitempty"`
}

// NewAddressDataRelationshipsEventsData instantiates a new AddressDataRelationshipsEventsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressDataRelationshipsEventsData() *AddressDataRelationshipsEventsData {
	this := AddressDataRelationshipsEventsData{}
	return &this
}

// NewAddressDataRelationshipsEventsDataWithDefaults instantiates a new AddressDataRelationshipsEventsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressDataRelationshipsEventsDataWithDefaults() *AddressDataRelationshipsEventsData {
	this := AddressDataRelationshipsEventsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressDataRelationshipsEventsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressDataRelationshipsEventsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AddressDataRelationshipsEventsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *AddressDataRelationshipsEventsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressDataRelationshipsEventsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressDataRelationshipsEventsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AddressDataRelationshipsEventsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *AddressDataRelationshipsEventsData) SetId(v interface{}) {
	o.Id = v
}

func (o AddressDataRelationshipsEventsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressDataRelationshipsEventsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableAddressDataRelationshipsEventsData struct {
	value *AddressDataRelationshipsEventsData
	isSet bool
}

func (v NullableAddressDataRelationshipsEventsData) Get() *AddressDataRelationshipsEventsData {
	return v.value
}

func (v *NullableAddressDataRelationshipsEventsData) Set(val *AddressDataRelationshipsEventsData) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressDataRelationshipsEventsData) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressDataRelationshipsEventsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressDataRelationshipsEventsData(val *AddressDataRelationshipsEventsData) *NullableAddressDataRelationshipsEventsData {
	return &NullableAddressDataRelationshipsEventsData{value: val, isSet: true}
}

func (v NullableAddressDataRelationshipsEventsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressDataRelationshipsEventsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
