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

// checks if the POSTPrices201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPrices201Response{}

// POSTPrices201Response struct for POSTPrices201Response
type POSTPrices201Response struct {
	Data *POSTPrices201ResponseData `json:"data,omitempty"`
}

// NewPOSTPrices201Response instantiates a new POSTPrices201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPrices201Response() *POSTPrices201Response {
	this := POSTPrices201Response{}
	return &this
}

// NewPOSTPrices201ResponseWithDefaults instantiates a new POSTPrices201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPrices201ResponseWithDefaults() *POSTPrices201Response {
	this := POSTPrices201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTPrices201Response) GetData() POSTPrices201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTPrices201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPrices201Response) GetDataOk() (*POSTPrices201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTPrices201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTPrices201ResponseData and assigns it to the Data field.
func (o *POSTPrices201Response) SetData(v POSTPrices201ResponseData) {
	o.Data = &v
}

func (o POSTPrices201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPrices201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTPrices201Response struct {
	value *POSTPrices201Response
	isSet bool
}

func (v NullablePOSTPrices201Response) Get() *POSTPrices201Response {
	return v.value
}

func (v *NullablePOSTPrices201Response) Set(val *POSTPrices201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPrices201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPrices201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPrices201Response(val *POSTPrices201Response) *NullablePOSTPrices201Response {
	return &NullablePOSTPrices201Response{value: val, isSet: true}
}

func (v NullablePOSTPrices201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPrices201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
