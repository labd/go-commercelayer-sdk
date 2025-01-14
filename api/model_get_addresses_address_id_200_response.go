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

// checks if the GETAddressesAddressId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETAddressesAddressId200Response{}

// GETAddressesAddressId200Response struct for GETAddressesAddressId200Response
type GETAddressesAddressId200Response struct {
	Data *GETAddressesAddressId200ResponseData `json:"data,omitempty"`
}

// NewGETAddressesAddressId200Response instantiates a new GETAddressesAddressId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAddressesAddressId200Response() *GETAddressesAddressId200Response {
	this := GETAddressesAddressId200Response{}
	return &this
}

// NewGETAddressesAddressId200ResponseWithDefaults instantiates a new GETAddressesAddressId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAddressesAddressId200ResponseWithDefaults() *GETAddressesAddressId200Response {
	this := GETAddressesAddressId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETAddressesAddressId200Response) GetData() GETAddressesAddressId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETAddressesAddressId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAddressesAddressId200Response) GetDataOk() (*GETAddressesAddressId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETAddressesAddressId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETAddressesAddressId200ResponseData and assigns it to the Data field.
func (o *GETAddressesAddressId200Response) SetData(v GETAddressesAddressId200ResponseData) {
	o.Data = &v
}

func (o GETAddressesAddressId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETAddressesAddressId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETAddressesAddressId200Response struct {
	value *GETAddressesAddressId200Response
	isSet bool
}

func (v NullableGETAddressesAddressId200Response) Get() *GETAddressesAddressId200Response {
	return v.value
}

func (v *NullableGETAddressesAddressId200Response) Set(val *GETAddressesAddressId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAddressesAddressId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAddressesAddressId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAddressesAddressId200Response(val *GETAddressesAddressId200Response) *NullableGETAddressesAddressId200Response {
	return &NullableGETAddressesAddressId200Response{value: val, isSet: true}
}

func (v NullableGETAddressesAddressId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAddressesAddressId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
