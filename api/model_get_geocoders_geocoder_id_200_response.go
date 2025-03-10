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

// checks if the GETGeocodersGeocoderId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETGeocodersGeocoderId200Response{}

// GETGeocodersGeocoderId200Response struct for GETGeocodersGeocoderId200Response
type GETGeocodersGeocoderId200Response struct {
	Data *GETGeocodersGeocoderId200ResponseData `json:"data,omitempty"`
}

// NewGETGeocodersGeocoderId200Response instantiates a new GETGeocodersGeocoderId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETGeocodersGeocoderId200Response() *GETGeocodersGeocoderId200Response {
	this := GETGeocodersGeocoderId200Response{}
	return &this
}

// NewGETGeocodersGeocoderId200ResponseWithDefaults instantiates a new GETGeocodersGeocoderId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETGeocodersGeocoderId200ResponseWithDefaults() *GETGeocodersGeocoderId200Response {
	this := GETGeocodersGeocoderId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETGeocodersGeocoderId200Response) GetData() GETGeocodersGeocoderId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETGeocodersGeocoderId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETGeocodersGeocoderId200Response) GetDataOk() (*GETGeocodersGeocoderId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETGeocodersGeocoderId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETGeocodersGeocoderId200ResponseData and assigns it to the Data field.
func (o *GETGeocodersGeocoderId200Response) SetData(v GETGeocodersGeocoderId200ResponseData) {
	o.Data = &v
}

func (o GETGeocodersGeocoderId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETGeocodersGeocoderId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETGeocodersGeocoderId200Response struct {
	value *GETGeocodersGeocoderId200Response
	isSet bool
}

func (v NullableGETGeocodersGeocoderId200Response) Get() *GETGeocodersGeocoderId200Response {
	return v.value
}

func (v *NullableGETGeocodersGeocoderId200Response) Set(val *GETGeocodersGeocoderId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETGeocodersGeocoderId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETGeocodersGeocoderId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETGeocodersGeocoderId200Response(val *GETGeocodersGeocoderId200Response) *NullableGETGeocodersGeocoderId200Response {
	return &NullableGETGeocodersGeocoderId200Response{value: val, isSet: true}
}

func (v NullableGETGeocodersGeocoderId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETGeocodersGeocoderId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
