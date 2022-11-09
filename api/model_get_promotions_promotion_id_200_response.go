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

// GETPromotionsPromotionId200Response struct for GETPromotionsPromotionId200Response
type GETPromotionsPromotionId200Response struct {
	Data *GETPromotions200ResponseDataInner `json:"data,omitempty"`
}

// NewGETPromotionsPromotionId200Response instantiates a new GETPromotionsPromotionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPromotionsPromotionId200Response() *GETPromotionsPromotionId200Response {
	this := GETPromotionsPromotionId200Response{}
	return &this
}

// NewGETPromotionsPromotionId200ResponseWithDefaults instantiates a new GETPromotionsPromotionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPromotionsPromotionId200ResponseWithDefaults() *GETPromotionsPromotionId200Response {
	this := GETPromotionsPromotionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETPromotionsPromotionId200Response) GetData() GETPromotions200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETPromotions200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPromotionsPromotionId200Response) GetDataOk() (*GETPromotions200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETPromotionsPromotionId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETPromotions200ResponseDataInner and assigns it to the Data field.
func (o *GETPromotionsPromotionId200Response) SetData(v GETPromotions200ResponseDataInner) {
	o.Data = &v
}

func (o GETPromotionsPromotionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETPromotionsPromotionId200Response struct {
	value *GETPromotionsPromotionId200Response
	isSet bool
}

func (v NullableGETPromotionsPromotionId200Response) Get() *GETPromotionsPromotionId200Response {
	return v.value
}

func (v *NullableGETPromotionsPromotionId200Response) Set(val *GETPromotionsPromotionId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPromotionsPromotionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPromotionsPromotionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPromotionsPromotionId200Response(val *GETPromotionsPromotionId200Response) *NullableGETPromotionsPromotionId200Response {
	return &NullableGETPromotionsPromotionId200Response{value: val, isSet: true}
}

func (v NullableGETPromotionsPromotionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPromotionsPromotionId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
