/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PATCHTaxCategoriesTaxCategoryId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHTaxCategoriesTaxCategoryId200Response{}

// PATCHTaxCategoriesTaxCategoryId200Response struct for PATCHTaxCategoriesTaxCategoryId200Response
type PATCHTaxCategoriesTaxCategoryId200Response struct {
	Data *PATCHTaxCategoriesTaxCategoryId200ResponseData `json:"data,omitempty"`
}

// NewPATCHTaxCategoriesTaxCategoryId200Response instantiates a new PATCHTaxCategoriesTaxCategoryId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHTaxCategoriesTaxCategoryId200Response() *PATCHTaxCategoriesTaxCategoryId200Response {
	this := PATCHTaxCategoriesTaxCategoryId200Response{}
	return &this
}

// NewPATCHTaxCategoriesTaxCategoryId200ResponseWithDefaults instantiates a new PATCHTaxCategoriesTaxCategoryId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHTaxCategoriesTaxCategoryId200ResponseWithDefaults() *PATCHTaxCategoriesTaxCategoryId200Response {
	this := PATCHTaxCategoriesTaxCategoryId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHTaxCategoriesTaxCategoryId200Response) GetData() PATCHTaxCategoriesTaxCategoryId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHTaxCategoriesTaxCategoryId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHTaxCategoriesTaxCategoryId200Response) GetDataOk() (*PATCHTaxCategoriesTaxCategoryId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHTaxCategoriesTaxCategoryId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHTaxCategoriesTaxCategoryId200ResponseData and assigns it to the Data field.
func (o *PATCHTaxCategoriesTaxCategoryId200Response) SetData(v PATCHTaxCategoriesTaxCategoryId200ResponseData) {
	o.Data = &v
}

func (o PATCHTaxCategoriesTaxCategoryId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHTaxCategoriesTaxCategoryId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHTaxCategoriesTaxCategoryId200Response struct {
	value *PATCHTaxCategoriesTaxCategoryId200Response
	isSet bool
}

func (v NullablePATCHTaxCategoriesTaxCategoryId200Response) Get() *PATCHTaxCategoriesTaxCategoryId200Response {
	return v.value
}

func (v *NullablePATCHTaxCategoriesTaxCategoryId200Response) Set(val *PATCHTaxCategoriesTaxCategoryId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHTaxCategoriesTaxCategoryId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHTaxCategoriesTaxCategoryId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHTaxCategoriesTaxCategoryId200Response(val *PATCHTaxCategoriesTaxCategoryId200Response) *NullablePATCHTaxCategoriesTaxCategoryId200Response {
	return &NullablePATCHTaxCategoriesTaxCategoryId200Response{value: val, isSet: true}
}

func (v NullablePATCHTaxCategoriesTaxCategoryId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHTaxCategoriesTaxCategoryId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
