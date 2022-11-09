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

// PATCHGoogleGeocodersGoogleGeocoderId200Response struct for PATCHGoogleGeocodersGoogleGeocoderId200Response
type PATCHGoogleGeocodersGoogleGeocoderId200Response struct {
	Data *PATCHGoogleGeocodersGoogleGeocoderId200ResponseData `json:"data,omitempty"`
}

// NewPATCHGoogleGeocodersGoogleGeocoderId200Response instantiates a new PATCHGoogleGeocodersGoogleGeocoderId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHGoogleGeocodersGoogleGeocoderId200Response() *PATCHGoogleGeocodersGoogleGeocoderId200Response {
	this := PATCHGoogleGeocodersGoogleGeocoderId200Response{}
	return &this
}

// NewPATCHGoogleGeocodersGoogleGeocoderId200ResponseWithDefaults instantiates a new PATCHGoogleGeocodersGoogleGeocoderId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHGoogleGeocodersGoogleGeocoderId200ResponseWithDefaults() *PATCHGoogleGeocodersGoogleGeocoderId200Response {
	this := PATCHGoogleGeocodersGoogleGeocoderId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHGoogleGeocodersGoogleGeocoderId200Response) GetData() PATCHGoogleGeocodersGoogleGeocoderId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHGoogleGeocodersGoogleGeocoderId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGoogleGeocodersGoogleGeocoderId200Response) GetDataOk() (*PATCHGoogleGeocodersGoogleGeocoderId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHGoogleGeocodersGoogleGeocoderId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHGoogleGeocodersGoogleGeocoderId200ResponseData and assigns it to the Data field.
func (o *PATCHGoogleGeocodersGoogleGeocoderId200Response) SetData(v PATCHGoogleGeocodersGoogleGeocoderId200ResponseData) {
	o.Data = &v
}

func (o PATCHGoogleGeocodersGoogleGeocoderId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHGoogleGeocodersGoogleGeocoderId200Response struct {
	value *PATCHGoogleGeocodersGoogleGeocoderId200Response
	isSet bool
}

func (v NullablePATCHGoogleGeocodersGoogleGeocoderId200Response) Get() *PATCHGoogleGeocodersGoogleGeocoderId200Response {
	return v.value
}

func (v *NullablePATCHGoogleGeocodersGoogleGeocoderId200Response) Set(val *PATCHGoogleGeocodersGoogleGeocoderId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHGoogleGeocodersGoogleGeocoderId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHGoogleGeocodersGoogleGeocoderId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHGoogleGeocodersGoogleGeocoderId200Response(val *PATCHGoogleGeocodersGoogleGeocoderId200Response) *NullablePATCHGoogleGeocodersGoogleGeocoderId200Response {
	return &NullablePATCHGoogleGeocodersGoogleGeocoderId200Response{value: val, isSet: true}
}

func (v NullablePATCHGoogleGeocodersGoogleGeocoderId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHGoogleGeocodersGoogleGeocoderId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
