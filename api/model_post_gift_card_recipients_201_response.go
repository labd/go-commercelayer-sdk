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

// POSTGiftCardRecipients201Response struct for POSTGiftCardRecipients201Response
type POSTGiftCardRecipients201Response struct {
	Data *POSTGiftCardRecipients201ResponseData `json:"data,omitempty"`
}

// NewPOSTGiftCardRecipients201Response instantiates a new POSTGiftCardRecipients201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTGiftCardRecipients201Response() *POSTGiftCardRecipients201Response {
	this := POSTGiftCardRecipients201Response{}
	return &this
}

// NewPOSTGiftCardRecipients201ResponseWithDefaults instantiates a new POSTGiftCardRecipients201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTGiftCardRecipients201ResponseWithDefaults() *POSTGiftCardRecipients201Response {
	this := POSTGiftCardRecipients201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTGiftCardRecipients201Response) GetData() POSTGiftCardRecipients201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTGiftCardRecipients201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCardRecipients201Response) GetDataOk() (*POSTGiftCardRecipients201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTGiftCardRecipients201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTGiftCardRecipients201ResponseData and assigns it to the Data field.
func (o *POSTGiftCardRecipients201Response) SetData(v POSTGiftCardRecipients201ResponseData) {
	o.Data = &v
}

func (o POSTGiftCardRecipients201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTGiftCardRecipients201Response struct {
	value *POSTGiftCardRecipients201Response
	isSet bool
}

func (v NullablePOSTGiftCardRecipients201Response) Get() *POSTGiftCardRecipients201Response {
	return v.value
}

func (v *NullablePOSTGiftCardRecipients201Response) Set(val *POSTGiftCardRecipients201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTGiftCardRecipients201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTGiftCardRecipients201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTGiftCardRecipients201Response(val *POSTGiftCardRecipients201Response) *NullablePOSTGiftCardRecipients201Response {
	return &NullablePOSTGiftCardRecipients201Response{value: val, isSet: true}
}

func (v NullablePOSTGiftCardRecipients201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTGiftCardRecipients201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
