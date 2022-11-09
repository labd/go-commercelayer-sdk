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

// PATCHShippingCategoriesShippingCategoryId200Response struct for PATCHShippingCategoriesShippingCategoryId200Response
type PATCHShippingCategoriesShippingCategoryId200Response struct {
	Data *PATCHShippingCategoriesShippingCategoryId200ResponseData `json:"data,omitempty"`
}

// NewPATCHShippingCategoriesShippingCategoryId200Response instantiates a new PATCHShippingCategoriesShippingCategoryId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHShippingCategoriesShippingCategoryId200Response() *PATCHShippingCategoriesShippingCategoryId200Response {
	this := PATCHShippingCategoriesShippingCategoryId200Response{}
	return &this
}

// NewPATCHShippingCategoriesShippingCategoryId200ResponseWithDefaults instantiates a new PATCHShippingCategoriesShippingCategoryId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHShippingCategoriesShippingCategoryId200ResponseWithDefaults() *PATCHShippingCategoriesShippingCategoryId200Response {
	this := PATCHShippingCategoriesShippingCategoryId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHShippingCategoriesShippingCategoryId200Response) GetData() PATCHShippingCategoriesShippingCategoryId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHShippingCategoriesShippingCategoryId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHShippingCategoriesShippingCategoryId200Response) GetDataOk() (*PATCHShippingCategoriesShippingCategoryId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHShippingCategoriesShippingCategoryId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHShippingCategoriesShippingCategoryId200ResponseData and assigns it to the Data field.
func (o *PATCHShippingCategoriesShippingCategoryId200Response) SetData(v PATCHShippingCategoriesShippingCategoryId200ResponseData) {
	o.Data = &v
}

func (o PATCHShippingCategoriesShippingCategoryId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHShippingCategoriesShippingCategoryId200Response struct {
	value *PATCHShippingCategoriesShippingCategoryId200Response
	isSet bool
}

func (v NullablePATCHShippingCategoriesShippingCategoryId200Response) Get() *PATCHShippingCategoriesShippingCategoryId200Response {
	return v.value
}

func (v *NullablePATCHShippingCategoriesShippingCategoryId200Response) Set(val *PATCHShippingCategoriesShippingCategoryId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHShippingCategoriesShippingCategoryId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHShippingCategoriesShippingCategoryId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHShippingCategoriesShippingCategoryId200Response(val *PATCHShippingCategoriesShippingCategoryId200Response) *NullablePATCHShippingCategoriesShippingCategoryId200Response {
	return &NullablePATCHShippingCategoriesShippingCategoryId200Response{value: val, isSet: true}
}

func (v NullablePATCHShippingCategoriesShippingCategoryId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHShippingCategoriesShippingCategoryId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}