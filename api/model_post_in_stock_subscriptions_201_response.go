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

// checks if the POSTInStockSubscriptions201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTInStockSubscriptions201Response{}

// POSTInStockSubscriptions201Response struct for POSTInStockSubscriptions201Response
type POSTInStockSubscriptions201Response struct {
	Data *POSTInStockSubscriptions201ResponseData `json:"data,omitempty"`
}

// NewPOSTInStockSubscriptions201Response instantiates a new POSTInStockSubscriptions201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTInStockSubscriptions201Response() *POSTInStockSubscriptions201Response {
	this := POSTInStockSubscriptions201Response{}
	return &this
}

// NewPOSTInStockSubscriptions201ResponseWithDefaults instantiates a new POSTInStockSubscriptions201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTInStockSubscriptions201ResponseWithDefaults() *POSTInStockSubscriptions201Response {
	this := POSTInStockSubscriptions201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTInStockSubscriptions201Response) GetData() POSTInStockSubscriptions201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTInStockSubscriptions201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTInStockSubscriptions201Response) GetDataOk() (*POSTInStockSubscriptions201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTInStockSubscriptions201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTInStockSubscriptions201ResponseData and assigns it to the Data field.
func (o *POSTInStockSubscriptions201Response) SetData(v POSTInStockSubscriptions201ResponseData) {
	o.Data = &v
}

func (o POSTInStockSubscriptions201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTInStockSubscriptions201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTInStockSubscriptions201Response struct {
	value *POSTInStockSubscriptions201Response
	isSet bool
}

func (v NullablePOSTInStockSubscriptions201Response) Get() *POSTInStockSubscriptions201Response {
	return v.value
}

func (v *NullablePOSTInStockSubscriptions201Response) Set(val *POSTInStockSubscriptions201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTInStockSubscriptions201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTInStockSubscriptions201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTInStockSubscriptions201Response(val *POSTInStockSubscriptions201Response) *NullablePOSTInStockSubscriptions201Response {
	return &NullablePOSTInStockSubscriptions201Response{value: val, isSet: true}
}

func (v NullablePOSTInStockSubscriptions201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTInStockSubscriptions201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
