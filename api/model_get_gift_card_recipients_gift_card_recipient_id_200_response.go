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

// checks if the GETGiftCardRecipientsGiftCardRecipientId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETGiftCardRecipientsGiftCardRecipientId200Response{}

// GETGiftCardRecipientsGiftCardRecipientId200Response struct for GETGiftCardRecipientsGiftCardRecipientId200Response
type GETGiftCardRecipientsGiftCardRecipientId200Response struct {
	Data *GETGiftCardRecipientsGiftCardRecipientId200ResponseData `json:"data,omitempty"`
}

// NewGETGiftCardRecipientsGiftCardRecipientId200Response instantiates a new GETGiftCardRecipientsGiftCardRecipientId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETGiftCardRecipientsGiftCardRecipientId200Response() *GETGiftCardRecipientsGiftCardRecipientId200Response {
	this := GETGiftCardRecipientsGiftCardRecipientId200Response{}
	return &this
}

// NewGETGiftCardRecipientsGiftCardRecipientId200ResponseWithDefaults instantiates a new GETGiftCardRecipientsGiftCardRecipientId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETGiftCardRecipientsGiftCardRecipientId200ResponseWithDefaults() *GETGiftCardRecipientsGiftCardRecipientId200Response {
	this := GETGiftCardRecipientsGiftCardRecipientId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETGiftCardRecipientsGiftCardRecipientId200Response) GetData() GETGiftCardRecipientsGiftCardRecipientId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETGiftCardRecipientsGiftCardRecipientId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETGiftCardRecipientsGiftCardRecipientId200Response) GetDataOk() (*GETGiftCardRecipientsGiftCardRecipientId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETGiftCardRecipientsGiftCardRecipientId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETGiftCardRecipientsGiftCardRecipientId200ResponseData and assigns it to the Data field.
func (o *GETGiftCardRecipientsGiftCardRecipientId200Response) SetData(v GETGiftCardRecipientsGiftCardRecipientId200ResponseData) {
	o.Data = &v
}

func (o GETGiftCardRecipientsGiftCardRecipientId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETGiftCardRecipientsGiftCardRecipientId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETGiftCardRecipientsGiftCardRecipientId200Response struct {
	value *GETGiftCardRecipientsGiftCardRecipientId200Response
	isSet bool
}

func (v NullableGETGiftCardRecipientsGiftCardRecipientId200Response) Get() *GETGiftCardRecipientsGiftCardRecipientId200Response {
	return v.value
}

func (v *NullableGETGiftCardRecipientsGiftCardRecipientId200Response) Set(val *GETGiftCardRecipientsGiftCardRecipientId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETGiftCardRecipientsGiftCardRecipientId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETGiftCardRecipientsGiftCardRecipientId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETGiftCardRecipientsGiftCardRecipientId200Response(val *GETGiftCardRecipientsGiftCardRecipientId200Response) *NullableGETGiftCardRecipientsGiftCardRecipientId200Response {
	return &NullableGETGiftCardRecipientsGiftCardRecipientId200Response{value: val, isSet: true}
}

func (v NullableGETGiftCardRecipientsGiftCardRecipientId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETGiftCardRecipientsGiftCardRecipientId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
