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

// checks if the POSTRecurringOrderCopies201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTRecurringOrderCopies201Response{}

// POSTRecurringOrderCopies201Response struct for POSTRecurringOrderCopies201Response
type POSTRecurringOrderCopies201Response struct {
	Data *POSTRecurringOrderCopies201ResponseData `json:"data,omitempty"`
}

// NewPOSTRecurringOrderCopies201Response instantiates a new POSTRecurringOrderCopies201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTRecurringOrderCopies201Response() *POSTRecurringOrderCopies201Response {
	this := POSTRecurringOrderCopies201Response{}
	return &this
}

// NewPOSTRecurringOrderCopies201ResponseWithDefaults instantiates a new POSTRecurringOrderCopies201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTRecurringOrderCopies201ResponseWithDefaults() *POSTRecurringOrderCopies201Response {
	this := POSTRecurringOrderCopies201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTRecurringOrderCopies201Response) GetData() POSTRecurringOrderCopies201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTRecurringOrderCopies201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTRecurringOrderCopies201Response) GetDataOk() (*POSTRecurringOrderCopies201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTRecurringOrderCopies201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTRecurringOrderCopies201ResponseData and assigns it to the Data field.
func (o *POSTRecurringOrderCopies201Response) SetData(v POSTRecurringOrderCopies201ResponseData) {
	o.Data = &v
}

func (o POSTRecurringOrderCopies201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTRecurringOrderCopies201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTRecurringOrderCopies201Response struct {
	value *POSTRecurringOrderCopies201Response
	isSet bool
}

func (v NullablePOSTRecurringOrderCopies201Response) Get() *POSTRecurringOrderCopies201Response {
	return v.value
}

func (v *NullablePOSTRecurringOrderCopies201Response) Set(val *POSTRecurringOrderCopies201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTRecurringOrderCopies201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTRecurringOrderCopies201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTRecurringOrderCopies201Response(val *POSTRecurringOrderCopies201Response) *NullablePOSTRecurringOrderCopies201Response {
	return &NullablePOSTRecurringOrderCopies201Response{value: val, isSet: true}
}

func (v NullablePOSTRecurringOrderCopies201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTRecurringOrderCopies201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
