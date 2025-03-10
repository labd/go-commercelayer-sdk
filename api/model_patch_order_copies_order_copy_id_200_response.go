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

// checks if the PATCHOrderCopiesOrderCopyId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHOrderCopiesOrderCopyId200Response{}

// PATCHOrderCopiesOrderCopyId200Response struct for PATCHOrderCopiesOrderCopyId200Response
type PATCHOrderCopiesOrderCopyId200Response struct {
	Data *PATCHOrderCopiesOrderCopyId200ResponseData `json:"data,omitempty"`
}

// NewPATCHOrderCopiesOrderCopyId200Response instantiates a new PATCHOrderCopiesOrderCopyId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHOrderCopiesOrderCopyId200Response() *PATCHOrderCopiesOrderCopyId200Response {
	this := PATCHOrderCopiesOrderCopyId200Response{}
	return &this
}

// NewPATCHOrderCopiesOrderCopyId200ResponseWithDefaults instantiates a new PATCHOrderCopiesOrderCopyId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHOrderCopiesOrderCopyId200ResponseWithDefaults() *PATCHOrderCopiesOrderCopyId200Response {
	this := PATCHOrderCopiesOrderCopyId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHOrderCopiesOrderCopyId200Response) GetData() PATCHOrderCopiesOrderCopyId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHOrderCopiesOrderCopyId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHOrderCopiesOrderCopyId200Response) GetDataOk() (*PATCHOrderCopiesOrderCopyId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHOrderCopiesOrderCopyId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHOrderCopiesOrderCopyId200ResponseData and assigns it to the Data field.
func (o *PATCHOrderCopiesOrderCopyId200Response) SetData(v PATCHOrderCopiesOrderCopyId200ResponseData) {
	o.Data = &v
}

func (o PATCHOrderCopiesOrderCopyId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHOrderCopiesOrderCopyId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHOrderCopiesOrderCopyId200Response struct {
	value *PATCHOrderCopiesOrderCopyId200Response
	isSet bool
}

func (v NullablePATCHOrderCopiesOrderCopyId200Response) Get() *PATCHOrderCopiesOrderCopyId200Response {
	return v.value
}

func (v *NullablePATCHOrderCopiesOrderCopyId200Response) Set(val *PATCHOrderCopiesOrderCopyId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHOrderCopiesOrderCopyId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHOrderCopiesOrderCopyId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHOrderCopiesOrderCopyId200Response(val *PATCHOrderCopiesOrderCopyId200Response) *NullablePATCHOrderCopiesOrderCopyId200Response {
	return &NullablePATCHOrderCopiesOrderCopyId200Response{value: val, isSet: true}
}

func (v NullablePATCHOrderCopiesOrderCopyId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHOrderCopiesOrderCopyId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
