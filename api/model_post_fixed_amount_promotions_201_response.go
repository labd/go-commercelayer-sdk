/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// POSTFixedAmountPromotions201Response struct for POSTFixedAmountPromotions201Response
type POSTFixedAmountPromotions201Response struct {
	Data *POSTFixedAmountPromotions201ResponseData `json:"data,omitempty"`
}

// NewPOSTFixedAmountPromotions201Response instantiates a new POSTFixedAmountPromotions201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTFixedAmountPromotions201Response() *POSTFixedAmountPromotions201Response {
	this := POSTFixedAmountPromotions201Response{}
	return &this
}

// NewPOSTFixedAmountPromotions201ResponseWithDefaults instantiates a new POSTFixedAmountPromotions201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTFixedAmountPromotions201ResponseWithDefaults() *POSTFixedAmountPromotions201Response {
	this := POSTFixedAmountPromotions201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTFixedAmountPromotions201Response) GetData() POSTFixedAmountPromotions201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTFixedAmountPromotions201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTFixedAmountPromotions201Response) GetDataOk() (*POSTFixedAmountPromotions201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTFixedAmountPromotions201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTFixedAmountPromotions201ResponseData and assigns it to the Data field.
func (o *POSTFixedAmountPromotions201Response) SetData(v POSTFixedAmountPromotions201ResponseData) {
	o.Data = &v
}

func (o POSTFixedAmountPromotions201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTFixedAmountPromotions201Response struct {
	value *POSTFixedAmountPromotions201Response
	isSet bool
}

func (v NullablePOSTFixedAmountPromotions201Response) Get() *POSTFixedAmountPromotions201Response {
	return v.value
}

func (v *NullablePOSTFixedAmountPromotions201Response) Set(val *POSTFixedAmountPromotions201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTFixedAmountPromotions201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTFixedAmountPromotions201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTFixedAmountPromotions201Response(val *POSTFixedAmountPromotions201Response) *NullablePOSTFixedAmountPromotions201Response {
	return &NullablePOSTFixedAmountPromotions201Response{value: val, isSet: true}
}

func (v NullablePOSTFixedAmountPromotions201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTFixedAmountPromotions201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
