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

// checks if the POSTOrders201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrders201Response{}

// POSTOrders201Response struct for POSTOrders201Response
type POSTOrders201Response struct {
	Data *POSTOrders201ResponseData `json:"data,omitempty"`
}

// NewPOSTOrders201Response instantiates a new POSTOrders201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrders201Response() *POSTOrders201Response {
	this := POSTOrders201Response{}
	return &this
}

// NewPOSTOrders201ResponseWithDefaults instantiates a new POSTOrders201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrders201ResponseWithDefaults() *POSTOrders201Response {
	this := POSTOrders201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTOrders201Response) GetData() POSTOrders201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTOrders201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201Response) GetDataOk() (*POSTOrders201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTOrders201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTOrders201ResponseData and assigns it to the Data field.
func (o *POSTOrders201Response) SetData(v POSTOrders201ResponseData) {
	o.Data = &v
}

func (o POSTOrders201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrders201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTOrders201Response struct {
	value *POSTOrders201Response
	isSet bool
}

func (v NullablePOSTOrders201Response) Get() *POSTOrders201Response {
	return v.value
}

func (v *NullablePOSTOrders201Response) Set(val *POSTOrders201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrders201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrders201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrders201Response(val *POSTOrders201Response) *NullablePOSTOrders201Response {
	return &NullablePOSTOrders201Response{value: val, isSet: true}
}

func (v NullablePOSTOrders201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrders201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
