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

// checks if the POSTCleanups201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCleanups201Response{}

// POSTCleanups201Response struct for POSTCleanups201Response
type POSTCleanups201Response struct {
	Data *POSTCleanups201ResponseData `json:"data,omitempty"`
}

// NewPOSTCleanups201Response instantiates a new POSTCleanups201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCleanups201Response() *POSTCleanups201Response {
	this := POSTCleanups201Response{}
	return &this
}

// NewPOSTCleanups201ResponseWithDefaults instantiates a new POSTCleanups201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCleanups201ResponseWithDefaults() *POSTCleanups201Response {
	this := POSTCleanups201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCleanups201Response) GetData() POSTCleanups201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTCleanups201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCleanups201Response) GetDataOk() (*POSTCleanups201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCleanups201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCleanups201ResponseData and assigns it to the Data field.
func (o *POSTCleanups201Response) SetData(v POSTCleanups201ResponseData) {
	o.Data = &v
}

func (o POSTCleanups201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCleanups201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTCleanups201Response struct {
	value *POSTCleanups201Response
	isSet bool
}

func (v NullablePOSTCleanups201Response) Get() *POSTCleanups201Response {
	return v.value
}

func (v *NullablePOSTCleanups201Response) Set(val *POSTCleanups201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCleanups201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCleanups201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCleanups201Response(val *POSTCleanups201Response) *NullablePOSTCleanups201Response {
	return &NullablePOSTCleanups201Response{value: val, isSet: true}
}

func (v NullablePOSTCleanups201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCleanups201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
