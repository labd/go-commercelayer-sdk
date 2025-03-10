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

// checks if the MerchantUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantUpdate{}

// MerchantUpdate struct for MerchantUpdate
type MerchantUpdate struct {
	Data MerchantUpdateData `json:"data"`
}

// NewMerchantUpdate instantiates a new MerchantUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantUpdate(data MerchantUpdateData) *MerchantUpdate {
	this := MerchantUpdate{}
	this.Data = data
	return &this
}

// NewMerchantUpdateWithDefaults instantiates a new MerchantUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantUpdateWithDefaults() *MerchantUpdate {
	this := MerchantUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *MerchantUpdate) GetData() MerchantUpdateData {
	if o == nil {
		var ret MerchantUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MerchantUpdate) GetDataOk() (*MerchantUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MerchantUpdate) SetData(v MerchantUpdateData) {
	o.Data = v
}

func (o MerchantUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableMerchantUpdate struct {
	value *MerchantUpdate
	isSet bool
}

func (v NullableMerchantUpdate) Get() *MerchantUpdate {
	return v.value
}

func (v *NullableMerchantUpdate) Set(val *MerchantUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantUpdate(val *MerchantUpdate) *NullableMerchantUpdate {
	return &NullableMerchantUpdate{value: val, isSet: true}
}

func (v NullableMerchantUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
