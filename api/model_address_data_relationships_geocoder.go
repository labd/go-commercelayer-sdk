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

// checks if the AddressDataRelationshipsGeocoder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressDataRelationshipsGeocoder{}

// AddressDataRelationshipsGeocoder struct for AddressDataRelationshipsGeocoder
type AddressDataRelationshipsGeocoder struct {
	Data *AddressDataRelationshipsGeocoderData `json:"data,omitempty"`
}

// NewAddressDataRelationshipsGeocoder instantiates a new AddressDataRelationshipsGeocoder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressDataRelationshipsGeocoder() *AddressDataRelationshipsGeocoder {
	this := AddressDataRelationshipsGeocoder{}
	return &this
}

// NewAddressDataRelationshipsGeocoderWithDefaults instantiates a new AddressDataRelationshipsGeocoder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressDataRelationshipsGeocoderWithDefaults() *AddressDataRelationshipsGeocoder {
	this := AddressDataRelationshipsGeocoder{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AddressDataRelationshipsGeocoder) GetData() AddressDataRelationshipsGeocoderData {
	if o == nil || IsNil(o.Data) {
		var ret AddressDataRelationshipsGeocoderData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressDataRelationshipsGeocoder) GetDataOk() (*AddressDataRelationshipsGeocoderData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AddressDataRelationshipsGeocoder) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AddressDataRelationshipsGeocoderData and assigns it to the Data field.
func (o *AddressDataRelationshipsGeocoder) SetData(v AddressDataRelationshipsGeocoderData) {
	o.Data = &v
}

func (o AddressDataRelationshipsGeocoder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressDataRelationshipsGeocoder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAddressDataRelationshipsGeocoder struct {
	value *AddressDataRelationshipsGeocoder
	isSet bool
}

func (v NullableAddressDataRelationshipsGeocoder) Get() *AddressDataRelationshipsGeocoder {
	return v.value
}

func (v *NullableAddressDataRelationshipsGeocoder) Set(val *AddressDataRelationshipsGeocoder) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressDataRelationshipsGeocoder) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressDataRelationshipsGeocoder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressDataRelationshipsGeocoder(val *AddressDataRelationshipsGeocoder) *NullableAddressDataRelationshipsGeocoder {
	return &NullableAddressDataRelationshipsGeocoder{value: val, isSet: true}
}

func (v NullableAddressDataRelationshipsGeocoder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressDataRelationshipsGeocoder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
