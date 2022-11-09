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

// PATCHReturnsReturnId200Response struct for PATCHReturnsReturnId200Response
type PATCHReturnsReturnId200Response struct {
	Data *PATCHReturnsReturnId200ResponseData `json:"data,omitempty"`
}

// NewPATCHReturnsReturnId200Response instantiates a new PATCHReturnsReturnId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHReturnsReturnId200Response() *PATCHReturnsReturnId200Response {
	this := PATCHReturnsReturnId200Response{}
	return &this
}

// NewPATCHReturnsReturnId200ResponseWithDefaults instantiates a new PATCHReturnsReturnId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHReturnsReturnId200ResponseWithDefaults() *PATCHReturnsReturnId200Response {
	this := PATCHReturnsReturnId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHReturnsReturnId200Response) GetData() PATCHReturnsReturnId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHReturnsReturnId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHReturnsReturnId200Response) GetDataOk() (*PATCHReturnsReturnId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHReturnsReturnId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHReturnsReturnId200ResponseData and assigns it to the Data field.
func (o *PATCHReturnsReturnId200Response) SetData(v PATCHReturnsReturnId200ResponseData) {
	o.Data = &v
}

func (o PATCHReturnsReturnId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHReturnsReturnId200Response struct {
	value *PATCHReturnsReturnId200Response
	isSet bool
}

func (v NullablePATCHReturnsReturnId200Response) Get() *PATCHReturnsReturnId200Response {
	return v.value
}

func (v *NullablePATCHReturnsReturnId200Response) Set(val *PATCHReturnsReturnId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHReturnsReturnId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHReturnsReturnId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHReturnsReturnId200Response(val *PATCHReturnsReturnId200Response) *NullablePATCHReturnsReturnId200Response {
	return &NullablePATCHReturnsReturnId200Response{value: val, isSet: true}
}

func (v NullablePATCHReturnsReturnId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHReturnsReturnId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
