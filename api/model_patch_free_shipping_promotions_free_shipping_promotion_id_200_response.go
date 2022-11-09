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

// PATCHFreeShippingPromotionsFreeShippingPromotionId200Response struct for PATCHFreeShippingPromotionsFreeShippingPromotionId200Response
type PATCHFreeShippingPromotionsFreeShippingPromotionId200Response struct {
	Data *PATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseData `json:"data,omitempty"`
}

// NewPATCHFreeShippingPromotionsFreeShippingPromotionId200Response instantiates a new PATCHFreeShippingPromotionsFreeShippingPromotionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHFreeShippingPromotionsFreeShippingPromotionId200Response() *PATCHFreeShippingPromotionsFreeShippingPromotionId200Response {
	this := PATCHFreeShippingPromotionsFreeShippingPromotionId200Response{}
	return &this
}

// NewPATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseWithDefaults instantiates a new PATCHFreeShippingPromotionsFreeShippingPromotionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseWithDefaults() *PATCHFreeShippingPromotionsFreeShippingPromotionId200Response {
	this := PATCHFreeShippingPromotionsFreeShippingPromotionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHFreeShippingPromotionsFreeShippingPromotionId200Response) GetData() PATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHFreeShippingPromotionsFreeShippingPromotionId200Response) GetDataOk() (*PATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHFreeShippingPromotionsFreeShippingPromotionId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseData and assigns it to the Data field.
func (o *PATCHFreeShippingPromotionsFreeShippingPromotionId200Response) SetData(v PATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseData) {
	o.Data = &v
}

func (o PATCHFreeShippingPromotionsFreeShippingPromotionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHFreeShippingPromotionsFreeShippingPromotionId200Response struct {
	value *PATCHFreeShippingPromotionsFreeShippingPromotionId200Response
	isSet bool
}

func (v NullablePATCHFreeShippingPromotionsFreeShippingPromotionId200Response) Get() *PATCHFreeShippingPromotionsFreeShippingPromotionId200Response {
	return v.value
}

func (v *NullablePATCHFreeShippingPromotionsFreeShippingPromotionId200Response) Set(val *PATCHFreeShippingPromotionsFreeShippingPromotionId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHFreeShippingPromotionsFreeShippingPromotionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHFreeShippingPromotionsFreeShippingPromotionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHFreeShippingPromotionsFreeShippingPromotionId200Response(val *PATCHFreeShippingPromotionsFreeShippingPromotionId200Response) *NullablePATCHFreeShippingPromotionsFreeShippingPromotionId200Response {
	return &NullablePATCHFreeShippingPromotionsFreeShippingPromotionId200Response{value: val, isSet: true}
}

func (v NullablePATCHFreeShippingPromotionsFreeShippingPromotionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHFreeShippingPromotionsFreeShippingPromotionId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
