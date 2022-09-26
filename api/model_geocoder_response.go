/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GeocoderResponse struct for GeocoderResponse
type GeocoderResponse struct {
	Data *GeocoderResponseData `json:"data,omitempty"`
}

// NewGeocoderResponse instantiates a new GeocoderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeocoderResponse() *GeocoderResponse {
	this := GeocoderResponse{}
	return &this
}

// NewGeocoderResponseWithDefaults instantiates a new GeocoderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeocoderResponseWithDefaults() *GeocoderResponse {
	this := GeocoderResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GeocoderResponse) GetData() GeocoderResponseData {
	if o == nil || o.Data == nil {
		var ret GeocoderResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeocoderResponse) GetDataOk() (*GeocoderResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GeocoderResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GeocoderResponseData and assigns it to the Data field.
func (o *GeocoderResponse) SetData(v GeocoderResponseData) {
	o.Data = &v
}

func (o GeocoderResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGeocoderResponse struct {
	value *GeocoderResponse
	isSet bool
}

func (v NullableGeocoderResponse) Get() *GeocoderResponse {
	return v.value
}

func (v *NullableGeocoderResponse) Set(val *GeocoderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGeocoderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGeocoderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeocoderResponse(val *GeocoderResponse) *NullableGeocoderResponse {
	return &NullableGeocoderResponse{value: val, isSet: true}
}

func (v NullableGeocoderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeocoderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
