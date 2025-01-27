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

// checks if the PATCHLineItemsLineItemId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHLineItemsLineItemId200Response{}

// PATCHLineItemsLineItemId200Response struct for PATCHLineItemsLineItemId200Response
type PATCHLineItemsLineItemId200Response struct {
	Data *PATCHLineItemsLineItemId200ResponseData `json:"data,omitempty"`
}

// NewPATCHLineItemsLineItemId200Response instantiates a new PATCHLineItemsLineItemId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHLineItemsLineItemId200Response() *PATCHLineItemsLineItemId200Response {
	this := PATCHLineItemsLineItemId200Response{}
	return &this
}

// NewPATCHLineItemsLineItemId200ResponseWithDefaults instantiates a new PATCHLineItemsLineItemId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHLineItemsLineItemId200ResponseWithDefaults() *PATCHLineItemsLineItemId200Response {
	this := PATCHLineItemsLineItemId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHLineItemsLineItemId200Response) GetData() PATCHLineItemsLineItemId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHLineItemsLineItemId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHLineItemsLineItemId200Response) GetDataOk() (*PATCHLineItemsLineItemId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHLineItemsLineItemId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHLineItemsLineItemId200ResponseData and assigns it to the Data field.
func (o *PATCHLineItemsLineItemId200Response) SetData(v PATCHLineItemsLineItemId200ResponseData) {
	o.Data = &v
}

func (o PATCHLineItemsLineItemId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHLineItemsLineItemId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHLineItemsLineItemId200Response struct {
	value *PATCHLineItemsLineItemId200Response
	isSet bool
}

func (v NullablePATCHLineItemsLineItemId200Response) Get() *PATCHLineItemsLineItemId200Response {
	return v.value
}

func (v *NullablePATCHLineItemsLineItemId200Response) Set(val *PATCHLineItemsLineItemId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHLineItemsLineItemId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHLineItemsLineItemId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHLineItemsLineItemId200Response(val *PATCHLineItemsLineItemId200Response) *NullablePATCHLineItemsLineItemId200Response {
	return &NullablePATCHLineItemsLineItemId200Response{value: val, isSet: true}
}

func (v NullablePATCHLineItemsLineItemId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHLineItemsLineItemId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
