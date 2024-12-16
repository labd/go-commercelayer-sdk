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

// checks if the POSTLineItems201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTLineItems201Response{}

// POSTLineItems201Response struct for POSTLineItems201Response
type POSTLineItems201Response struct {
	Data *POSTLineItems201ResponseData `json:"data,omitempty"`
}

// NewPOSTLineItems201Response instantiates a new POSTLineItems201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTLineItems201Response() *POSTLineItems201Response {
	this := POSTLineItems201Response{}
	return &this
}

// NewPOSTLineItems201ResponseWithDefaults instantiates a new POSTLineItems201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTLineItems201ResponseWithDefaults() *POSTLineItems201Response {
	this := POSTLineItems201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTLineItems201Response) GetData() POSTLineItems201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTLineItems201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201Response) GetDataOk() (*POSTLineItems201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTLineItems201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTLineItems201ResponseData and assigns it to the Data field.
func (o *POSTLineItems201Response) SetData(v POSTLineItems201ResponseData) {
	o.Data = &v
}

func (o POSTLineItems201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTLineItems201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTLineItems201Response struct {
	value *POSTLineItems201Response
	isSet bool
}

func (v NullablePOSTLineItems201Response) Get() *POSTLineItems201Response {
	return v.value
}

func (v *NullablePOSTLineItems201Response) Set(val *POSTLineItems201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTLineItems201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTLineItems201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTLineItems201Response(val *POSTLineItems201Response) *NullablePOSTLineItems201Response {
	return &NullablePOSTLineItems201Response{value: val, isSet: true}
}

func (v NullablePOSTLineItems201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTLineItems201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
