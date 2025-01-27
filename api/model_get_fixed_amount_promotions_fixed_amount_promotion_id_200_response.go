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

// checks if the GETFixedAmountPromotionsFixedAmountPromotionId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETFixedAmountPromotionsFixedAmountPromotionId200Response{}

// GETFixedAmountPromotionsFixedAmountPromotionId200Response struct for GETFixedAmountPromotionsFixedAmountPromotionId200Response
type GETFixedAmountPromotionsFixedAmountPromotionId200Response struct {
	Data *GETFixedAmountPromotionsFixedAmountPromotionId200ResponseData `json:"data,omitempty"`
}

// NewGETFixedAmountPromotionsFixedAmountPromotionId200Response instantiates a new GETFixedAmountPromotionsFixedAmountPromotionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETFixedAmountPromotionsFixedAmountPromotionId200Response() *GETFixedAmountPromotionsFixedAmountPromotionId200Response {
	this := GETFixedAmountPromotionsFixedAmountPromotionId200Response{}
	return &this
}

// NewGETFixedAmountPromotionsFixedAmountPromotionId200ResponseWithDefaults instantiates a new GETFixedAmountPromotionsFixedAmountPromotionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETFixedAmountPromotionsFixedAmountPromotionId200ResponseWithDefaults() *GETFixedAmountPromotionsFixedAmountPromotionId200Response {
	this := GETFixedAmountPromotionsFixedAmountPromotionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200Response) GetData() GETFixedAmountPromotionsFixedAmountPromotionId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETFixedAmountPromotionsFixedAmountPromotionId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200Response) GetDataOk() (*GETFixedAmountPromotionsFixedAmountPromotionId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETFixedAmountPromotionsFixedAmountPromotionId200ResponseData and assigns it to the Data field.
func (o *GETFixedAmountPromotionsFixedAmountPromotionId200Response) SetData(v GETFixedAmountPromotionsFixedAmountPromotionId200ResponseData) {
	o.Data = &v
}

func (o GETFixedAmountPromotionsFixedAmountPromotionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETFixedAmountPromotionsFixedAmountPromotionId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETFixedAmountPromotionsFixedAmountPromotionId200Response struct {
	value *GETFixedAmountPromotionsFixedAmountPromotionId200Response
	isSet bool
}

func (v NullableGETFixedAmountPromotionsFixedAmountPromotionId200Response) Get() *GETFixedAmountPromotionsFixedAmountPromotionId200Response {
	return v.value
}

func (v *NullableGETFixedAmountPromotionsFixedAmountPromotionId200Response) Set(val *GETFixedAmountPromotionsFixedAmountPromotionId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETFixedAmountPromotionsFixedAmountPromotionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETFixedAmountPromotionsFixedAmountPromotionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETFixedAmountPromotionsFixedAmountPromotionId200Response(val *GETFixedAmountPromotionsFixedAmountPromotionId200Response) *NullableGETFixedAmountPromotionsFixedAmountPromotionId200Response {
	return &NullableGETFixedAmountPromotionsFixedAmountPromotionId200Response{value: val, isSet: true}
}

func (v NullableGETFixedAmountPromotionsFixedAmountPromotionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETFixedAmountPromotionsFixedAmountPromotionId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
