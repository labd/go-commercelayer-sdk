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

// checks if the POSTBingGeocoders201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTBingGeocoders201Response{}

// POSTBingGeocoders201Response struct for POSTBingGeocoders201Response
type POSTBingGeocoders201Response struct {
	Data *POSTBingGeocoders201ResponseData `json:"data,omitempty"`
}

// NewPOSTBingGeocoders201Response instantiates a new POSTBingGeocoders201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTBingGeocoders201Response() *POSTBingGeocoders201Response {
	this := POSTBingGeocoders201Response{}
	return &this
}

// NewPOSTBingGeocoders201ResponseWithDefaults instantiates a new POSTBingGeocoders201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTBingGeocoders201ResponseWithDefaults() *POSTBingGeocoders201Response {
	this := POSTBingGeocoders201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTBingGeocoders201Response) GetData() POSTBingGeocoders201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTBingGeocoders201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBingGeocoders201Response) GetDataOk() (*POSTBingGeocoders201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTBingGeocoders201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTBingGeocoders201ResponseData and assigns it to the Data field.
func (o *POSTBingGeocoders201Response) SetData(v POSTBingGeocoders201ResponseData) {
	o.Data = &v
}

func (o POSTBingGeocoders201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTBingGeocoders201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTBingGeocoders201Response struct {
	value *POSTBingGeocoders201Response
	isSet bool
}

func (v NullablePOSTBingGeocoders201Response) Get() *POSTBingGeocoders201Response {
	return v.value
}

func (v *NullablePOSTBingGeocoders201Response) Set(val *POSTBingGeocoders201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTBingGeocoders201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTBingGeocoders201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTBingGeocoders201Response(val *POSTBingGeocoders201Response) *NullablePOSTBingGeocoders201Response {
	return &NullablePOSTBingGeocoders201Response{value: val, isSet: true}
}

func (v NullablePOSTBingGeocoders201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTBingGeocoders201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
