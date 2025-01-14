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

// checks if the PATCHLinksLinkId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHLinksLinkId200Response{}

// PATCHLinksLinkId200Response struct for PATCHLinksLinkId200Response
type PATCHLinksLinkId200Response struct {
	Data *PATCHLinksLinkId200ResponseData `json:"data,omitempty"`
}

// NewPATCHLinksLinkId200Response instantiates a new PATCHLinksLinkId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHLinksLinkId200Response() *PATCHLinksLinkId200Response {
	this := PATCHLinksLinkId200Response{}
	return &this
}

// NewPATCHLinksLinkId200ResponseWithDefaults instantiates a new PATCHLinksLinkId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHLinksLinkId200ResponseWithDefaults() *PATCHLinksLinkId200Response {
	this := PATCHLinksLinkId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHLinksLinkId200Response) GetData() PATCHLinksLinkId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHLinksLinkId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHLinksLinkId200Response) GetDataOk() (*PATCHLinksLinkId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHLinksLinkId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHLinksLinkId200ResponseData and assigns it to the Data field.
func (o *PATCHLinksLinkId200Response) SetData(v PATCHLinksLinkId200ResponseData) {
	o.Data = &v
}

func (o PATCHLinksLinkId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHLinksLinkId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHLinksLinkId200Response struct {
	value *PATCHLinksLinkId200Response
	isSet bool
}

func (v NullablePATCHLinksLinkId200Response) Get() *PATCHLinksLinkId200Response {
	return v.value
}

func (v *NullablePATCHLinksLinkId200Response) Set(val *PATCHLinksLinkId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHLinksLinkId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHLinksLinkId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHLinksLinkId200Response(val *PATCHLinksLinkId200Response) *NullablePATCHLinksLinkId200Response {
	return &NullablePATCHLinksLinkId200Response{value: val, isSet: true}
}

func (v NullablePATCHLinksLinkId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHLinksLinkId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
