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

// checks if the GETReturnLineItemsReturnLineItemId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETReturnLineItemsReturnLineItemId200Response{}

// GETReturnLineItemsReturnLineItemId200Response struct for GETReturnLineItemsReturnLineItemId200Response
type GETReturnLineItemsReturnLineItemId200Response struct {
	Data *GETReturnLineItemsReturnLineItemId200ResponseData `json:"data,omitempty"`
}

// NewGETReturnLineItemsReturnLineItemId200Response instantiates a new GETReturnLineItemsReturnLineItemId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETReturnLineItemsReturnLineItemId200Response() *GETReturnLineItemsReturnLineItemId200Response {
	this := GETReturnLineItemsReturnLineItemId200Response{}
	return &this
}

// NewGETReturnLineItemsReturnLineItemId200ResponseWithDefaults instantiates a new GETReturnLineItemsReturnLineItemId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETReturnLineItemsReturnLineItemId200ResponseWithDefaults() *GETReturnLineItemsReturnLineItemId200Response {
	this := GETReturnLineItemsReturnLineItemId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETReturnLineItemsReturnLineItemId200Response) GetData() GETReturnLineItemsReturnLineItemId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETReturnLineItemsReturnLineItemId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnLineItemsReturnLineItemId200Response) GetDataOk() (*GETReturnLineItemsReturnLineItemId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETReturnLineItemsReturnLineItemId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETReturnLineItemsReturnLineItemId200ResponseData and assigns it to the Data field.
func (o *GETReturnLineItemsReturnLineItemId200Response) SetData(v GETReturnLineItemsReturnLineItemId200ResponseData) {
	o.Data = &v
}

func (o GETReturnLineItemsReturnLineItemId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETReturnLineItemsReturnLineItemId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETReturnLineItemsReturnLineItemId200Response struct {
	value *GETReturnLineItemsReturnLineItemId200Response
	isSet bool
}

func (v NullableGETReturnLineItemsReturnLineItemId200Response) Get() *GETReturnLineItemsReturnLineItemId200Response {
	return v.value
}

func (v *NullableGETReturnLineItemsReturnLineItemId200Response) Set(val *GETReturnLineItemsReturnLineItemId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETReturnLineItemsReturnLineItemId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETReturnLineItemsReturnLineItemId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETReturnLineItemsReturnLineItemId200Response(val *GETReturnLineItemsReturnLineItemId200Response) *NullableGETReturnLineItemsReturnLineItemId200Response {
	return &NullableGETReturnLineItemsReturnLineItemId200Response{value: val, isSet: true}
}

func (v NullableGETReturnLineItemsReturnLineItemId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETReturnLineItemsReturnLineItemId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
