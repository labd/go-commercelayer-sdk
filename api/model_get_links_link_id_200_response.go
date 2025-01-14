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

// checks if the GETLinksLinkId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETLinksLinkId200Response{}

// GETLinksLinkId200Response struct for GETLinksLinkId200Response
type GETLinksLinkId200Response struct {
	Data *GETLinksLinkId200ResponseData `json:"data,omitempty"`
}

// NewGETLinksLinkId200Response instantiates a new GETLinksLinkId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLinksLinkId200Response() *GETLinksLinkId200Response {
	this := GETLinksLinkId200Response{}
	return &this
}

// NewGETLinksLinkId200ResponseWithDefaults instantiates a new GETLinksLinkId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLinksLinkId200ResponseWithDefaults() *GETLinksLinkId200Response {
	this := GETLinksLinkId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETLinksLinkId200Response) GetData() GETLinksLinkId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETLinksLinkId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLinksLinkId200Response) GetDataOk() (*GETLinksLinkId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETLinksLinkId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETLinksLinkId200ResponseData and assigns it to the Data field.
func (o *GETLinksLinkId200Response) SetData(v GETLinksLinkId200ResponseData) {
	o.Data = &v
}

func (o GETLinksLinkId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETLinksLinkId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETLinksLinkId200Response struct {
	value *GETLinksLinkId200Response
	isSet bool
}

func (v NullableGETLinksLinkId200Response) Get() *GETLinksLinkId200Response {
	return v.value
}

func (v *NullableGETLinksLinkId200Response) Set(val *GETLinksLinkId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLinksLinkId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLinksLinkId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLinksLinkId200Response(val *GETLinksLinkId200Response) *NullableGETLinksLinkId200Response {
	return &NullableGETLinksLinkId200Response{value: val, isSet: true}
}

func (v NullableGETLinksLinkId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLinksLinkId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
