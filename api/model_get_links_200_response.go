/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETLinks200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETLinks200Response{}

// GETLinks200Response struct for GETLinks200Response
type GETLinks200Response struct {
	Data interface{} `json:"data,omitempty"`
}

// NewGETLinks200Response instantiates a new GETLinks200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLinks200Response() *GETLinks200Response {
	this := GETLinks200Response{}
	return &this
}

// NewGETLinks200ResponseWithDefaults instantiates a new GETLinks200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLinks200ResponseWithDefaults() *GETLinks200Response {
	this := GETLinks200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETLinks200Response) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETLinks200Response) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETLinks200Response) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *GETLinks200Response) SetData(v interface{}) {
	o.Data = v
}

func (o GETLinks200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETLinks200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETLinks200Response struct {
	value *GETLinks200Response
	isSet bool
}

func (v NullableGETLinks200Response) Get() *GETLinks200Response {
	return v.value
}

func (v *NullableGETLinks200Response) Set(val *GETLinks200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLinks200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLinks200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLinks200Response(val *GETLinks200Response) *NullableGETLinks200Response {
	return &NullableGETLinks200Response{value: val, isSet: true}
}

func (v NullableGETLinks200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLinks200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}