/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PATCHBingGeocodersBingGeocoderId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHBingGeocodersBingGeocoderId200Response{}

// PATCHBingGeocodersBingGeocoderId200Response struct for PATCHBingGeocodersBingGeocoderId200Response
type PATCHBingGeocodersBingGeocoderId200Response struct {
	Data *PATCHBingGeocodersBingGeocoderId200ResponseData `json:"data,omitempty"`
}

// NewPATCHBingGeocodersBingGeocoderId200Response instantiates a new PATCHBingGeocodersBingGeocoderId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHBingGeocodersBingGeocoderId200Response() *PATCHBingGeocodersBingGeocoderId200Response {
	this := PATCHBingGeocodersBingGeocoderId200Response{}
	return &this
}

// NewPATCHBingGeocodersBingGeocoderId200ResponseWithDefaults instantiates a new PATCHBingGeocodersBingGeocoderId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHBingGeocodersBingGeocoderId200ResponseWithDefaults() *PATCHBingGeocodersBingGeocoderId200Response {
	this := PATCHBingGeocodersBingGeocoderId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHBingGeocodersBingGeocoderId200Response) GetData() PATCHBingGeocodersBingGeocoderId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHBingGeocodersBingGeocoderId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHBingGeocodersBingGeocoderId200Response) GetDataOk() (*PATCHBingGeocodersBingGeocoderId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHBingGeocodersBingGeocoderId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHBingGeocodersBingGeocoderId200ResponseData and assigns it to the Data field.
func (o *PATCHBingGeocodersBingGeocoderId200Response) SetData(v PATCHBingGeocodersBingGeocoderId200ResponseData) {
	o.Data = &v
}

func (o PATCHBingGeocodersBingGeocoderId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHBingGeocodersBingGeocoderId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHBingGeocodersBingGeocoderId200Response struct {
	value *PATCHBingGeocodersBingGeocoderId200Response
	isSet bool
}

func (v NullablePATCHBingGeocodersBingGeocoderId200Response) Get() *PATCHBingGeocodersBingGeocoderId200Response {
	return v.value
}

func (v *NullablePATCHBingGeocodersBingGeocoderId200Response) Set(val *PATCHBingGeocodersBingGeocoderId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHBingGeocodersBingGeocoderId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHBingGeocodersBingGeocoderId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHBingGeocodersBingGeocoderId200Response(val *PATCHBingGeocodersBingGeocoderId200Response) *NullablePATCHBingGeocodersBingGeocoderId200Response {
	return &NullablePATCHBingGeocodersBingGeocoderId200Response{value: val, isSet: true}
}

func (v NullablePATCHBingGeocodersBingGeocoderId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHBingGeocodersBingGeocoderId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
