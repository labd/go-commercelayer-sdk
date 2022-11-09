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

// GETGiftCards200Response struct for GETGiftCards200Response
type GETGiftCards200Response struct {
	Data []GETGiftCards200ResponseDataInner `json:"data,omitempty"`
}

// NewGETGiftCards200Response instantiates a new GETGiftCards200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETGiftCards200Response() *GETGiftCards200Response {
	this := GETGiftCards200Response{}
	return &this
}

// NewGETGiftCards200ResponseWithDefaults instantiates a new GETGiftCards200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETGiftCards200ResponseWithDefaults() *GETGiftCards200Response {
	this := GETGiftCards200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETGiftCards200Response) GetData() []GETGiftCards200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETGiftCards200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETGiftCards200Response) GetDataOk() ([]GETGiftCards200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETGiftCards200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETGiftCards200ResponseDataInner and assigns it to the Data field.
func (o *GETGiftCards200Response) SetData(v []GETGiftCards200ResponseDataInner) {
	o.Data = v
}

func (o GETGiftCards200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETGiftCards200Response struct {
	value *GETGiftCards200Response
	isSet bool
}

func (v NullableGETGiftCards200Response) Get() *GETGiftCards200Response {
	return v.value
}

func (v *NullableGETGiftCards200Response) Set(val *GETGiftCards200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETGiftCards200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETGiftCards200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETGiftCards200Response(val *GETGiftCards200Response) *NullableGETGiftCards200Response {
	return &NullableGETGiftCards200Response{value: val, isSet: true}
}

func (v NullableGETGiftCards200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETGiftCards200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
