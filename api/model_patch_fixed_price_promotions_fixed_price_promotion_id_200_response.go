/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PATCHFixedPricePromotionsFixedPricePromotionId200Response struct for PATCHFixedPricePromotionsFixedPricePromotionId200Response
type PATCHFixedPricePromotionsFixedPricePromotionId200Response struct {
	Data *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData `json:"data,omitempty"`
}

// NewPATCHFixedPricePromotionsFixedPricePromotionId200Response instantiates a new PATCHFixedPricePromotionsFixedPricePromotionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHFixedPricePromotionsFixedPricePromotionId200Response() *PATCHFixedPricePromotionsFixedPricePromotionId200Response {
	this := PATCHFixedPricePromotionsFixedPricePromotionId200Response{}
	return &this
}

// NewPATCHFixedPricePromotionsFixedPricePromotionId200ResponseWithDefaults instantiates a new PATCHFixedPricePromotionsFixedPricePromotionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHFixedPricePromotionsFixedPricePromotionId200ResponseWithDefaults() *PATCHFixedPricePromotionsFixedPricePromotionId200Response {
	this := PATCHFixedPricePromotionsFixedPricePromotionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200Response) GetData() PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200Response) GetDataOk() (*PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData and assigns it to the Data field.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200Response) SetData(v PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) {
	o.Data = &v
}

func (o PATCHFixedPricePromotionsFixedPricePromotionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHFixedPricePromotionsFixedPricePromotionId200Response struct {
	value *PATCHFixedPricePromotionsFixedPricePromotionId200Response
	isSet bool
}

func (v NullablePATCHFixedPricePromotionsFixedPricePromotionId200Response) Get() *PATCHFixedPricePromotionsFixedPricePromotionId200Response {
	return v.value
}

func (v *NullablePATCHFixedPricePromotionsFixedPricePromotionId200Response) Set(val *PATCHFixedPricePromotionsFixedPricePromotionId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHFixedPricePromotionsFixedPricePromotionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHFixedPricePromotionsFixedPricePromotionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHFixedPricePromotionsFixedPricePromotionId200Response(val *PATCHFixedPricePromotionsFixedPricePromotionId200Response) *NullablePATCHFixedPricePromotionsFixedPricePromotionId200Response {
	return &NullablePATCHFixedPricePromotionsFixedPricePromotionId200Response{value: val, isSet: true}
}

func (v NullablePATCHFixedPricePromotionsFixedPricePromotionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHFixedPricePromotionsFixedPricePromotionId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
