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

// GETPrices200Response struct for GETPrices200Response
type GETPrices200Response struct {
	Data []GETPrices200ResponseDataInner `json:"data,omitempty"`
}

// NewGETPrices200Response instantiates a new GETPrices200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPrices200Response() *GETPrices200Response {
	this := GETPrices200Response{}
	return &this
}

// NewGETPrices200ResponseWithDefaults instantiates a new GETPrices200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPrices200ResponseWithDefaults() *GETPrices200Response {
	this := GETPrices200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETPrices200Response) GetData() []GETPrices200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETPrices200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPrices200Response) GetDataOk() ([]GETPrices200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETPrices200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETPrices200ResponseDataInner and assigns it to the Data field.
func (o *GETPrices200Response) SetData(v []GETPrices200ResponseDataInner) {
	o.Data = v
}

func (o GETPrices200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETPrices200Response struct {
	value *GETPrices200Response
	isSet bool
}

func (v NullableGETPrices200Response) Get() *GETPrices200Response {
	return v.value
}

func (v *NullableGETPrices200Response) Set(val *GETPrices200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPrices200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPrices200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPrices200Response(val *GETPrices200Response) *NullableGETPrices200Response {
	return &NullableGETPrices200Response{value: val, isSet: true}
}

func (v NullableGETPrices200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPrices200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
