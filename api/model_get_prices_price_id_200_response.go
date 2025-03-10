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

// checks if the GETPricesPriceId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETPricesPriceId200Response{}

// GETPricesPriceId200Response struct for GETPricesPriceId200Response
type GETPricesPriceId200Response struct {
	Data *GETPricesPriceId200ResponseData `json:"data,omitempty"`
}

// NewGETPricesPriceId200Response instantiates a new GETPricesPriceId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPricesPriceId200Response() *GETPricesPriceId200Response {
	this := GETPricesPriceId200Response{}
	return &this
}

// NewGETPricesPriceId200ResponseWithDefaults instantiates a new GETPricesPriceId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPricesPriceId200ResponseWithDefaults() *GETPricesPriceId200Response {
	this := GETPricesPriceId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETPricesPriceId200Response) GetData() GETPricesPriceId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETPricesPriceId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPricesPriceId200Response) GetDataOk() (*GETPricesPriceId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETPricesPriceId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETPricesPriceId200ResponseData and assigns it to the Data field.
func (o *GETPricesPriceId200Response) SetData(v GETPricesPriceId200ResponseData) {
	o.Data = &v
}

func (o GETPricesPriceId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETPricesPriceId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETPricesPriceId200Response struct {
	value *GETPricesPriceId200Response
	isSet bool
}

func (v NullableGETPricesPriceId200Response) Get() *GETPricesPriceId200Response {
	return v.value
}

func (v *NullableGETPricesPriceId200Response) Set(val *GETPricesPriceId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPricesPriceId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPricesPriceId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPricesPriceId200Response(val *GETPricesPriceId200Response) *NullableGETPricesPriceId200Response {
	return &NullableGETPricesPriceId200Response{value: val, isSet: true}
}

func (v NullableGETPricesPriceId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPricesPriceId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
