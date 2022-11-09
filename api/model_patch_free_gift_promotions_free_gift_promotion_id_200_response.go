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

// PATCHFreeGiftPromotionsFreeGiftPromotionId200Response struct for PATCHFreeGiftPromotionsFreeGiftPromotionId200Response
type PATCHFreeGiftPromotionsFreeGiftPromotionId200Response struct {
	Data *PATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseData `json:"data,omitempty"`
}

// NewPATCHFreeGiftPromotionsFreeGiftPromotionId200Response instantiates a new PATCHFreeGiftPromotionsFreeGiftPromotionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHFreeGiftPromotionsFreeGiftPromotionId200Response() *PATCHFreeGiftPromotionsFreeGiftPromotionId200Response {
	this := PATCHFreeGiftPromotionsFreeGiftPromotionId200Response{}
	return &this
}

// NewPATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseWithDefaults instantiates a new PATCHFreeGiftPromotionsFreeGiftPromotionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseWithDefaults() *PATCHFreeGiftPromotionsFreeGiftPromotionId200Response {
	this := PATCHFreeGiftPromotionsFreeGiftPromotionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHFreeGiftPromotionsFreeGiftPromotionId200Response) GetData() PATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHFreeGiftPromotionsFreeGiftPromotionId200Response) GetDataOk() (*PATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHFreeGiftPromotionsFreeGiftPromotionId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseData and assigns it to the Data field.
func (o *PATCHFreeGiftPromotionsFreeGiftPromotionId200Response) SetData(v PATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseData) {
	o.Data = &v
}

func (o PATCHFreeGiftPromotionsFreeGiftPromotionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHFreeGiftPromotionsFreeGiftPromotionId200Response struct {
	value *PATCHFreeGiftPromotionsFreeGiftPromotionId200Response
	isSet bool
}

func (v NullablePATCHFreeGiftPromotionsFreeGiftPromotionId200Response) Get() *PATCHFreeGiftPromotionsFreeGiftPromotionId200Response {
	return v.value
}

func (v *NullablePATCHFreeGiftPromotionsFreeGiftPromotionId200Response) Set(val *PATCHFreeGiftPromotionsFreeGiftPromotionId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHFreeGiftPromotionsFreeGiftPromotionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHFreeGiftPromotionsFreeGiftPromotionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHFreeGiftPromotionsFreeGiftPromotionId200Response(val *PATCHFreeGiftPromotionsFreeGiftPromotionId200Response) *NullablePATCHFreeGiftPromotionsFreeGiftPromotionId200Response {
	return &NullablePATCHFreeGiftPromotionsFreeGiftPromotionId200Response{value: val, isSet: true}
}

func (v NullablePATCHFreeGiftPromotionsFreeGiftPromotionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHFreeGiftPromotionsFreeGiftPromotionId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
