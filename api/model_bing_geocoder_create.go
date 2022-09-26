/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BingGeocoderCreate struct for BingGeocoderCreate
type BingGeocoderCreate struct {
	Data BingGeocoderCreateData `json:"data"`
}

// NewBingGeocoderCreate instantiates a new BingGeocoderCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBingGeocoderCreate(data BingGeocoderCreateData) *BingGeocoderCreate {
	this := BingGeocoderCreate{}
	this.Data = data
	return &this
}

// NewBingGeocoderCreateWithDefaults instantiates a new BingGeocoderCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBingGeocoderCreateWithDefaults() *BingGeocoderCreate {
	this := BingGeocoderCreate{}
	return &this
}

// GetData returns the Data field value
func (o *BingGeocoderCreate) GetData() BingGeocoderCreateData {
	if o == nil {
		var ret BingGeocoderCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BingGeocoderCreate) GetDataOk() (*BingGeocoderCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BingGeocoderCreate) SetData(v BingGeocoderCreateData) {
	o.Data = v
}

func (o BingGeocoderCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBingGeocoderCreate struct {
	value *BingGeocoderCreate
	isSet bool
}

func (v NullableBingGeocoderCreate) Get() *BingGeocoderCreate {
	return v.value
}

func (v *NullableBingGeocoderCreate) Set(val *BingGeocoderCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBingGeocoderCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBingGeocoderCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBingGeocoderCreate(val *BingGeocoderCreate) *NullableBingGeocoderCreate {
	return &NullableBingGeocoderCreate{value: val, isSet: true}
}

func (v NullableBingGeocoderCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBingGeocoderCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
