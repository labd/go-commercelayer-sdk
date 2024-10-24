/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SatispayGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SatispayGateway{}

// SatispayGateway struct for SatispayGateway
type SatispayGateway struct {
	Data *SatispayGatewayData `json:"data,omitempty"`
}

// NewSatispayGateway instantiates a new SatispayGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSatispayGateway() *SatispayGateway {
	this := SatispayGateway{}
	return &this
}

// NewSatispayGatewayWithDefaults instantiates a new SatispayGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSatispayGatewayWithDefaults() *SatispayGateway {
	this := SatispayGateway{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SatispayGateway) GetData() SatispayGatewayData {
	if o == nil || IsNil(o.Data) {
		var ret SatispayGatewayData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SatispayGateway) GetDataOk() (*SatispayGatewayData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SatispayGateway) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SatispayGatewayData and assigns it to the Data field.
func (o *SatispayGateway) SetData(v SatispayGatewayData) {
	o.Data = &v
}

func (o SatispayGateway) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SatispayGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSatispayGateway struct {
	value *SatispayGateway
	isSet bool
}

func (v NullableSatispayGateway) Get() *SatispayGateway {
	return v.value
}

func (v *NullableSatispayGateway) Set(val *SatispayGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableSatispayGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableSatispayGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSatispayGateway(val *SatispayGateway) *NullableSatispayGateway {
	return &NullableSatispayGateway{value: val, isSet: true}
}

func (v NullableSatispayGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSatispayGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}