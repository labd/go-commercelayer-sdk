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

// GETBingGeocodersBingGeocoderId200Response struct for GETBingGeocodersBingGeocoderId200Response
type GETBingGeocodersBingGeocoderId200Response struct {
	Data *GETBingGeocoders200ResponseDataInner `json:"data,omitempty"`
}

// NewGETBingGeocodersBingGeocoderId200Response instantiates a new GETBingGeocodersBingGeocoderId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETBingGeocodersBingGeocoderId200Response() *GETBingGeocodersBingGeocoderId200Response {
	this := GETBingGeocodersBingGeocoderId200Response{}
	return &this
}

// NewGETBingGeocodersBingGeocoderId200ResponseWithDefaults instantiates a new GETBingGeocodersBingGeocoderId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETBingGeocodersBingGeocoderId200ResponseWithDefaults() *GETBingGeocodersBingGeocoderId200Response {
	this := GETBingGeocodersBingGeocoderId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETBingGeocodersBingGeocoderId200Response) GetData() GETBingGeocoders200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETBingGeocoders200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBingGeocodersBingGeocoderId200Response) GetDataOk() (*GETBingGeocoders200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETBingGeocodersBingGeocoderId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETBingGeocoders200ResponseDataInner and assigns it to the Data field.
func (o *GETBingGeocodersBingGeocoderId200Response) SetData(v GETBingGeocoders200ResponseDataInner) {
	o.Data = &v
}

func (o GETBingGeocodersBingGeocoderId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETBingGeocodersBingGeocoderId200Response struct {
	value *GETBingGeocodersBingGeocoderId200Response
	isSet bool
}

func (v NullableGETBingGeocodersBingGeocoderId200Response) Get() *GETBingGeocodersBingGeocoderId200Response {
	return v.value
}

func (v *NullableGETBingGeocodersBingGeocoderId200Response) Set(val *GETBingGeocodersBingGeocoderId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETBingGeocodersBingGeocoderId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETBingGeocodersBingGeocoderId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETBingGeocodersBingGeocoderId200Response(val *GETBingGeocodersBingGeocoderId200Response) *NullableGETBingGeocodersBingGeocoderId200Response {
	return &NullableGETBingGeocodersBingGeocoderId200Response{value: val, isSet: true}
}

func (v NullableGETBingGeocodersBingGeocoderId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETBingGeocodersBingGeocoderId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}