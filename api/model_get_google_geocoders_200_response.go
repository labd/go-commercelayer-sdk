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

// GETGoogleGeocoders200Response struct for GETGoogleGeocoders200Response
type GETGoogleGeocoders200Response struct {
	Data []GETGoogleGeocoders200ResponseDataInner `json:"data,omitempty"`
}

// NewGETGoogleGeocoders200Response instantiates a new GETGoogleGeocoders200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETGoogleGeocoders200Response() *GETGoogleGeocoders200Response {
	this := GETGoogleGeocoders200Response{}
	return &this
}

// NewGETGoogleGeocoders200ResponseWithDefaults instantiates a new GETGoogleGeocoders200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETGoogleGeocoders200ResponseWithDefaults() *GETGoogleGeocoders200Response {
	this := GETGoogleGeocoders200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETGoogleGeocoders200Response) GetData() []GETGoogleGeocoders200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETGoogleGeocoders200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETGoogleGeocoders200Response) GetDataOk() ([]GETGoogleGeocoders200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETGoogleGeocoders200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETGoogleGeocoders200ResponseDataInner and assigns it to the Data field.
func (o *GETGoogleGeocoders200Response) SetData(v []GETGoogleGeocoders200ResponseDataInner) {
	o.Data = v
}

func (o GETGoogleGeocoders200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETGoogleGeocoders200Response struct {
	value *GETGoogleGeocoders200Response
	isSet bool
}

func (v NullableGETGoogleGeocoders200Response) Get() *GETGoogleGeocoders200Response {
	return v.value
}

func (v *NullableGETGoogleGeocoders200Response) Set(val *GETGoogleGeocoders200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETGoogleGeocoders200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETGoogleGeocoders200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETGoogleGeocoders200Response(val *GETGoogleGeocoders200Response) *NullableGETGoogleGeocoders200Response {
	return &NullableGETGoogleGeocoders200Response{value: val, isSet: true}
}

func (v NullableGETGoogleGeocoders200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETGoogleGeocoders200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
