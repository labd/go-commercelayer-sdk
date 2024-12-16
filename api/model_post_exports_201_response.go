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

// checks if the POSTExports201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTExports201Response{}

// POSTExports201Response struct for POSTExports201Response
type POSTExports201Response struct {
	Data *POSTExports201ResponseData `json:"data,omitempty"`
}

// NewPOSTExports201Response instantiates a new POSTExports201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExports201Response() *POSTExports201Response {
	this := POSTExports201Response{}
	return &this
}

// NewPOSTExports201ResponseWithDefaults instantiates a new POSTExports201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExports201ResponseWithDefaults() *POSTExports201Response {
	this := POSTExports201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTExports201Response) GetData() POSTExports201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTExports201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExports201Response) GetDataOk() (*POSTExports201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTExports201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTExports201ResponseData and assigns it to the Data field.
func (o *POSTExports201Response) SetData(v POSTExports201ResponseData) {
	o.Data = &v
}

func (o POSTExports201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTExports201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTExports201Response struct {
	value *POSTExports201Response
	isSet bool
}

func (v NullablePOSTExports201Response) Get() *POSTExports201Response {
	return v.value
}

func (v *NullablePOSTExports201Response) Set(val *POSTExports201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExports201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExports201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExports201Response(val *POSTExports201Response) *NullablePOSTExports201Response {
	return &NullablePOSTExports201Response{value: val, isSet: true}
}

func (v NullablePOSTExports201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExports201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
