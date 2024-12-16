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

// checks if the GETWireTransfersWireTransferId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETWireTransfersWireTransferId200Response{}

// GETWireTransfersWireTransferId200Response struct for GETWireTransfersWireTransferId200Response
type GETWireTransfersWireTransferId200Response struct {
	Data *GETWireTransfersWireTransferId200ResponseData `json:"data,omitempty"`
}

// NewGETWireTransfersWireTransferId200Response instantiates a new GETWireTransfersWireTransferId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETWireTransfersWireTransferId200Response() *GETWireTransfersWireTransferId200Response {
	this := GETWireTransfersWireTransferId200Response{}
	return &this
}

// NewGETWireTransfersWireTransferId200ResponseWithDefaults instantiates a new GETWireTransfersWireTransferId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETWireTransfersWireTransferId200ResponseWithDefaults() *GETWireTransfersWireTransferId200Response {
	this := GETWireTransfersWireTransferId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETWireTransfersWireTransferId200Response) GetData() GETWireTransfersWireTransferId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETWireTransfersWireTransferId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETWireTransfersWireTransferId200Response) GetDataOk() (*GETWireTransfersWireTransferId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETWireTransfersWireTransferId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETWireTransfersWireTransferId200ResponseData and assigns it to the Data field.
func (o *GETWireTransfersWireTransferId200Response) SetData(v GETWireTransfersWireTransferId200ResponseData) {
	o.Data = &v
}

func (o GETWireTransfersWireTransferId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETWireTransfersWireTransferId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETWireTransfersWireTransferId200Response struct {
	value *GETWireTransfersWireTransferId200Response
	isSet bool
}

func (v NullableGETWireTransfersWireTransferId200Response) Get() *GETWireTransfersWireTransferId200Response {
	return v.value
}

func (v *NullableGETWireTransfersWireTransferId200Response) Set(val *GETWireTransfersWireTransferId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETWireTransfersWireTransferId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETWireTransfersWireTransferId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETWireTransfersWireTransferId200Response(val *GETWireTransfersWireTransferId200Response) *NullableGETWireTransfersWireTransferId200Response {
	return &NullableGETWireTransfersWireTransferId200Response{value: val, isSet: true}
}

func (v NullableGETWireTransfersWireTransferId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETWireTransfersWireTransferId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
