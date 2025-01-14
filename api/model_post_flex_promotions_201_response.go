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

// checks if the POSTFlexPromotions201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTFlexPromotions201Response{}

// POSTFlexPromotions201Response struct for POSTFlexPromotions201Response
type POSTFlexPromotions201Response struct {
	Data *POSTFlexPromotions201ResponseData `json:"data,omitempty"`
}

// NewPOSTFlexPromotions201Response instantiates a new POSTFlexPromotions201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTFlexPromotions201Response() *POSTFlexPromotions201Response {
	this := POSTFlexPromotions201Response{}
	return &this
}

// NewPOSTFlexPromotions201ResponseWithDefaults instantiates a new POSTFlexPromotions201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTFlexPromotions201ResponseWithDefaults() *POSTFlexPromotions201Response {
	this := POSTFlexPromotions201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTFlexPromotions201Response) GetData() POSTFlexPromotions201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTFlexPromotions201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTFlexPromotions201Response) GetDataOk() (*POSTFlexPromotions201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTFlexPromotions201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTFlexPromotions201ResponseData and assigns it to the Data field.
func (o *POSTFlexPromotions201Response) SetData(v POSTFlexPromotions201ResponseData) {
	o.Data = &v
}

func (o POSTFlexPromotions201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTFlexPromotions201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTFlexPromotions201Response struct {
	value *POSTFlexPromotions201Response
	isSet bool
}

func (v NullablePOSTFlexPromotions201Response) Get() *POSTFlexPromotions201Response {
	return v.value
}

func (v *NullablePOSTFlexPromotions201Response) Set(val *POSTFlexPromotions201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTFlexPromotions201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTFlexPromotions201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTFlexPromotions201Response(val *POSTFlexPromotions201Response) *NullablePOSTFlexPromotions201Response {
	return &NullablePOSTFlexPromotions201Response{value: val, isSet: true}
}

func (v NullablePOSTFlexPromotions201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTFlexPromotions201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
