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

// checks if the GETParcelsParcelId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETParcelsParcelId200Response{}

// GETParcelsParcelId200Response struct for GETParcelsParcelId200Response
type GETParcelsParcelId200Response struct {
	Data *GETParcelsParcelId200ResponseData `json:"data,omitempty"`
}

// NewGETParcelsParcelId200Response instantiates a new GETParcelsParcelId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETParcelsParcelId200Response() *GETParcelsParcelId200Response {
	this := GETParcelsParcelId200Response{}
	return &this
}

// NewGETParcelsParcelId200ResponseWithDefaults instantiates a new GETParcelsParcelId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETParcelsParcelId200ResponseWithDefaults() *GETParcelsParcelId200Response {
	this := GETParcelsParcelId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETParcelsParcelId200Response) GetData() GETParcelsParcelId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETParcelsParcelId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelsParcelId200Response) GetDataOk() (*GETParcelsParcelId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETParcelsParcelId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETParcelsParcelId200ResponseData and assigns it to the Data field.
func (o *GETParcelsParcelId200Response) SetData(v GETParcelsParcelId200ResponseData) {
	o.Data = &v
}

func (o GETParcelsParcelId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETParcelsParcelId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETParcelsParcelId200Response struct {
	value *GETParcelsParcelId200Response
	isSet bool
}

func (v NullableGETParcelsParcelId200Response) Get() *GETParcelsParcelId200Response {
	return v.value
}

func (v *NullableGETParcelsParcelId200Response) Set(val *GETParcelsParcelId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETParcelsParcelId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETParcelsParcelId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETParcelsParcelId200Response(val *GETParcelsParcelId200Response) *NullableGETParcelsParcelId200Response {
	return &NullableGETParcelsParcelId200Response{value: val, isSet: true}
}

func (v NullableGETParcelsParcelId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETParcelsParcelId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
