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

// checks if the Parcel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Parcel{}

// Parcel struct for Parcel
type Parcel struct {
	Data *ParcelData `json:"data,omitempty"`
}

// NewParcel instantiates a new Parcel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcel() *Parcel {
	this := Parcel{}
	return &this
}

// NewParcelWithDefaults instantiates a new Parcel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelWithDefaults() *Parcel {
	this := Parcel{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Parcel) GetData() ParcelData {
	if o == nil || IsNil(o.Data) {
		var ret ParcelData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Parcel) GetDataOk() (*ParcelData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Parcel) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ParcelData and assigns it to the Data field.
func (o *Parcel) SetData(v ParcelData) {
	o.Data = &v
}

func (o Parcel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Parcel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableParcel struct {
	value *Parcel
	isSet bool
}

func (v NullableParcel) Get() *Parcel {
	return v.value
}

func (v *NullableParcel) Set(val *Parcel) {
	v.value = val
	v.isSet = true
}

func (v NullableParcel) IsSet() bool {
	return v.isSet
}

func (v *NullableParcel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcel(val *Parcel) *NullableParcel {
	return &NullableParcel{value: val, isSet: true}
}

func (v NullableParcel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
