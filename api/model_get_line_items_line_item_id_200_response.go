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

// checks if the GETLineItemsLineItemId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETLineItemsLineItemId200Response{}

// GETLineItemsLineItemId200Response struct for GETLineItemsLineItemId200Response
type GETLineItemsLineItemId200Response struct {
	Data *GETLineItemsLineItemId200ResponseData `json:"data,omitempty"`
}

// NewGETLineItemsLineItemId200Response instantiates a new GETLineItemsLineItemId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLineItemsLineItemId200Response() *GETLineItemsLineItemId200Response {
	this := GETLineItemsLineItemId200Response{}
	return &this
}

// NewGETLineItemsLineItemId200ResponseWithDefaults instantiates a new GETLineItemsLineItemId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLineItemsLineItemId200ResponseWithDefaults() *GETLineItemsLineItemId200Response {
	this := GETLineItemsLineItemId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETLineItemsLineItemId200Response) GetData() GETLineItemsLineItemId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETLineItemsLineItemId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItemsLineItemId200Response) GetDataOk() (*GETLineItemsLineItemId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETLineItemsLineItemId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETLineItemsLineItemId200ResponseData and assigns it to the Data field.
func (o *GETLineItemsLineItemId200Response) SetData(v GETLineItemsLineItemId200ResponseData) {
	o.Data = &v
}

func (o GETLineItemsLineItemId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETLineItemsLineItemId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETLineItemsLineItemId200Response struct {
	value *GETLineItemsLineItemId200Response
	isSet bool
}

func (v NullableGETLineItemsLineItemId200Response) Get() *GETLineItemsLineItemId200Response {
	return v.value
}

func (v *NullableGETLineItemsLineItemId200Response) Set(val *GETLineItemsLineItemId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLineItemsLineItemId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLineItemsLineItemId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLineItemsLineItemId200Response(val *GETLineItemsLineItemId200Response) *NullableGETLineItemsLineItemId200Response {
	return &NullableGETLineItemsLineItemId200Response{value: val, isSet: true}
}

func (v NullableGETLineItemsLineItemId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLineItemsLineItemId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
