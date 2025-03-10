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

// checks if the GETSkusSkuId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETSkusSkuId200Response{}

// GETSkusSkuId200Response struct for GETSkusSkuId200Response
type GETSkusSkuId200Response struct {
	Data *GETSkusSkuId200ResponseData `json:"data,omitempty"`
}

// NewGETSkusSkuId200Response instantiates a new GETSkusSkuId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSkusSkuId200Response() *GETSkusSkuId200Response {
	this := GETSkusSkuId200Response{}
	return &this
}

// NewGETSkusSkuId200ResponseWithDefaults instantiates a new GETSkusSkuId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSkusSkuId200ResponseWithDefaults() *GETSkusSkuId200Response {
	this := GETSkusSkuId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETSkusSkuId200Response) GetData() GETSkusSkuId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETSkusSkuId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkusSkuId200Response) GetDataOk() (*GETSkusSkuId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETSkusSkuId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETSkusSkuId200ResponseData and assigns it to the Data field.
func (o *GETSkusSkuId200Response) SetData(v GETSkusSkuId200ResponseData) {
	o.Data = &v
}

func (o GETSkusSkuId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETSkusSkuId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETSkusSkuId200Response struct {
	value *GETSkusSkuId200Response
	isSet bool
}

func (v NullableGETSkusSkuId200Response) Get() *GETSkusSkuId200Response {
	return v.value
}

func (v *NullableGETSkusSkuId200Response) Set(val *GETSkusSkuId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSkusSkuId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSkusSkuId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSkusSkuId200Response(val *GETSkusSkuId200Response) *NullableGETSkusSkuId200Response {
	return &NullableGETSkusSkuId200Response{value: val, isSet: true}
}

func (v NullableGETSkusSkuId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSkusSkuId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
