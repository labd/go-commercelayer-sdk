/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// POSTInventoryModels201Response struct for POSTInventoryModels201Response
type POSTInventoryModels201Response struct {
	Data *POSTInventoryModels201ResponseData `json:"data,omitempty"`
}

// NewPOSTInventoryModels201Response instantiates a new POSTInventoryModels201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTInventoryModels201Response() *POSTInventoryModels201Response {
	this := POSTInventoryModels201Response{}
	return &this
}

// NewPOSTInventoryModels201ResponseWithDefaults instantiates a new POSTInventoryModels201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTInventoryModels201ResponseWithDefaults() *POSTInventoryModels201Response {
	this := POSTInventoryModels201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTInventoryModels201Response) GetData() POSTInventoryModels201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTInventoryModels201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTInventoryModels201Response) GetDataOk() (*POSTInventoryModels201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTInventoryModels201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTInventoryModels201ResponseData and assigns it to the Data field.
func (o *POSTInventoryModels201Response) SetData(v POSTInventoryModels201ResponseData) {
	o.Data = &v
}

func (o POSTInventoryModels201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTInventoryModels201Response struct {
	value *POSTInventoryModels201Response
	isSet bool
}

func (v NullablePOSTInventoryModels201Response) Get() *POSTInventoryModels201Response {
	return v.value
}

func (v *NullablePOSTInventoryModels201Response) Set(val *POSTInventoryModels201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTInventoryModels201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTInventoryModels201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTInventoryModels201Response(val *POSTInventoryModels201Response) *NullablePOSTInventoryModels201Response {
	return &NullablePOSTInventoryModels201Response{value: val, isSet: true}
}

func (v NullablePOSTInventoryModels201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTInventoryModels201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
