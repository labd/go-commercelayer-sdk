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

// checks if the GETVoidsVoidId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETVoidsVoidId200Response{}

// GETVoidsVoidId200Response struct for GETVoidsVoidId200Response
type GETVoidsVoidId200Response struct {
	Data *GETVoidsVoidId200ResponseData `json:"data,omitempty"`
}

// NewGETVoidsVoidId200Response instantiates a new GETVoidsVoidId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETVoidsVoidId200Response() *GETVoidsVoidId200Response {
	this := GETVoidsVoidId200Response{}
	return &this
}

// NewGETVoidsVoidId200ResponseWithDefaults instantiates a new GETVoidsVoidId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETVoidsVoidId200ResponseWithDefaults() *GETVoidsVoidId200Response {
	this := GETVoidsVoidId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETVoidsVoidId200Response) GetData() GETVoidsVoidId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETVoidsVoidId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETVoidsVoidId200Response) GetDataOk() (*GETVoidsVoidId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETVoidsVoidId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETVoidsVoidId200ResponseData and assigns it to the Data field.
func (o *GETVoidsVoidId200Response) SetData(v GETVoidsVoidId200ResponseData) {
	o.Data = &v
}

func (o GETVoidsVoidId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETVoidsVoidId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETVoidsVoidId200Response struct {
	value *GETVoidsVoidId200Response
	isSet bool
}

func (v NullableGETVoidsVoidId200Response) Get() *GETVoidsVoidId200Response {
	return v.value
}

func (v *NullableGETVoidsVoidId200Response) Set(val *GETVoidsVoidId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETVoidsVoidId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETVoidsVoidId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETVoidsVoidId200Response(val *GETVoidsVoidId200Response) *NullableGETVoidsVoidId200Response {
	return &NullableGETVoidsVoidId200Response{value: val, isSet: true}
}

func (v NullableGETVoidsVoidId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETVoidsVoidId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
