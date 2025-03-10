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

// checks if the POSTShippingWeightTiers201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTShippingWeightTiers201Response{}

// POSTShippingWeightTiers201Response struct for POSTShippingWeightTiers201Response
type POSTShippingWeightTiers201Response struct {
	Data *POSTShippingWeightTiers201ResponseData `json:"data,omitempty"`
}

// NewPOSTShippingWeightTiers201Response instantiates a new POSTShippingWeightTiers201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShippingWeightTiers201Response() *POSTShippingWeightTiers201Response {
	this := POSTShippingWeightTiers201Response{}
	return &this
}

// NewPOSTShippingWeightTiers201ResponseWithDefaults instantiates a new POSTShippingWeightTiers201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShippingWeightTiers201ResponseWithDefaults() *POSTShippingWeightTiers201Response {
	this := POSTShippingWeightTiers201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTShippingWeightTiers201Response) GetData() POSTShippingWeightTiers201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTShippingWeightTiers201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingWeightTiers201Response) GetDataOk() (*POSTShippingWeightTiers201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTShippingWeightTiers201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTShippingWeightTiers201ResponseData and assigns it to the Data field.
func (o *POSTShippingWeightTiers201Response) SetData(v POSTShippingWeightTiers201ResponseData) {
	o.Data = &v
}

func (o POSTShippingWeightTiers201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTShippingWeightTiers201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTShippingWeightTiers201Response struct {
	value *POSTShippingWeightTiers201Response
	isSet bool
}

func (v NullablePOSTShippingWeightTiers201Response) Get() *POSTShippingWeightTiers201Response {
	return v.value
}

func (v *NullablePOSTShippingWeightTiers201Response) Set(val *POSTShippingWeightTiers201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShippingWeightTiers201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShippingWeightTiers201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShippingWeightTiers201Response(val *POSTShippingWeightTiers201Response) *NullablePOSTShippingWeightTiers201Response {
	return &NullablePOSTShippingWeightTiers201Response{value: val, isSet: true}
}

func (v NullablePOSTShippingWeightTiers201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShippingWeightTiers201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
