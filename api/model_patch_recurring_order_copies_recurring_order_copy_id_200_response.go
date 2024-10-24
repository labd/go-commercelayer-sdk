/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PATCHRecurringOrderCopiesRecurringOrderCopyId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHRecurringOrderCopiesRecurringOrderCopyId200Response{}

// PATCHRecurringOrderCopiesRecurringOrderCopyId200Response struct for PATCHRecurringOrderCopiesRecurringOrderCopyId200Response
type PATCHRecurringOrderCopiesRecurringOrderCopyId200Response struct {
	Data *PATCHRecurringOrderCopiesRecurringOrderCopyId200ResponseData `json:"data,omitempty"`
}

// NewPATCHRecurringOrderCopiesRecurringOrderCopyId200Response instantiates a new PATCHRecurringOrderCopiesRecurringOrderCopyId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHRecurringOrderCopiesRecurringOrderCopyId200Response() *PATCHRecurringOrderCopiesRecurringOrderCopyId200Response {
	this := PATCHRecurringOrderCopiesRecurringOrderCopyId200Response{}
	return &this
}

// NewPATCHRecurringOrderCopiesRecurringOrderCopyId200ResponseWithDefaults instantiates a new PATCHRecurringOrderCopiesRecurringOrderCopyId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHRecurringOrderCopiesRecurringOrderCopyId200ResponseWithDefaults() *PATCHRecurringOrderCopiesRecurringOrderCopyId200Response {
	this := PATCHRecurringOrderCopiesRecurringOrderCopyId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHRecurringOrderCopiesRecurringOrderCopyId200Response) GetData() PATCHRecurringOrderCopiesRecurringOrderCopyId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHRecurringOrderCopiesRecurringOrderCopyId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHRecurringOrderCopiesRecurringOrderCopyId200Response) GetDataOk() (*PATCHRecurringOrderCopiesRecurringOrderCopyId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHRecurringOrderCopiesRecurringOrderCopyId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHRecurringOrderCopiesRecurringOrderCopyId200ResponseData and assigns it to the Data field.
func (o *PATCHRecurringOrderCopiesRecurringOrderCopyId200Response) SetData(v PATCHRecurringOrderCopiesRecurringOrderCopyId200ResponseData) {
	o.Data = &v
}

func (o PATCHRecurringOrderCopiesRecurringOrderCopyId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHRecurringOrderCopiesRecurringOrderCopyId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHRecurringOrderCopiesRecurringOrderCopyId200Response struct {
	value *PATCHRecurringOrderCopiesRecurringOrderCopyId200Response
	isSet bool
}

func (v NullablePATCHRecurringOrderCopiesRecurringOrderCopyId200Response) Get() *PATCHRecurringOrderCopiesRecurringOrderCopyId200Response {
	return v.value
}

func (v *NullablePATCHRecurringOrderCopiesRecurringOrderCopyId200Response) Set(val *PATCHRecurringOrderCopiesRecurringOrderCopyId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHRecurringOrderCopiesRecurringOrderCopyId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHRecurringOrderCopiesRecurringOrderCopyId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHRecurringOrderCopiesRecurringOrderCopyId200Response(val *PATCHRecurringOrderCopiesRecurringOrderCopyId200Response) *NullablePATCHRecurringOrderCopiesRecurringOrderCopyId200Response {
	return &NullablePATCHRecurringOrderCopiesRecurringOrderCopyId200Response{value: val, isSet: true}
}

func (v NullablePATCHRecurringOrderCopiesRecurringOrderCopyId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHRecurringOrderCopiesRecurringOrderCopyId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}