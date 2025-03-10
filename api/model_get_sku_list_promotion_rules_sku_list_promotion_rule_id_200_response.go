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

// checks if the GETSkuListPromotionRulesSkuListPromotionRuleId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETSkuListPromotionRulesSkuListPromotionRuleId200Response{}

// GETSkuListPromotionRulesSkuListPromotionRuleId200Response struct for GETSkuListPromotionRulesSkuListPromotionRuleId200Response
type GETSkuListPromotionRulesSkuListPromotionRuleId200Response struct {
	Data *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseData `json:"data,omitempty"`
}

// NewGETSkuListPromotionRulesSkuListPromotionRuleId200Response instantiates a new GETSkuListPromotionRulesSkuListPromotionRuleId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSkuListPromotionRulesSkuListPromotionRuleId200Response() *GETSkuListPromotionRulesSkuListPromotionRuleId200Response {
	this := GETSkuListPromotionRulesSkuListPromotionRuleId200Response{}
	return &this
}

// NewGETSkuListPromotionRulesSkuListPromotionRuleId200ResponseWithDefaults instantiates a new GETSkuListPromotionRulesSkuListPromotionRuleId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSkuListPromotionRulesSkuListPromotionRuleId200ResponseWithDefaults() *GETSkuListPromotionRulesSkuListPromotionRuleId200Response {
	this := GETSkuListPromotionRulesSkuListPromotionRuleId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200Response) GetData() GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200Response) GetDataOk() (*GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseData and assigns it to the Data field.
func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200Response) SetData(v GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) {
	o.Data = &v
}

func (o GETSkuListPromotionRulesSkuListPromotionRuleId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETSkuListPromotionRulesSkuListPromotionRuleId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETSkuListPromotionRulesSkuListPromotionRuleId200Response struct {
	value *GETSkuListPromotionRulesSkuListPromotionRuleId200Response
	isSet bool
}

func (v NullableGETSkuListPromotionRulesSkuListPromotionRuleId200Response) Get() *GETSkuListPromotionRulesSkuListPromotionRuleId200Response {
	return v.value
}

func (v *NullableGETSkuListPromotionRulesSkuListPromotionRuleId200Response) Set(val *GETSkuListPromotionRulesSkuListPromotionRuleId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSkuListPromotionRulesSkuListPromotionRuleId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSkuListPromotionRulesSkuListPromotionRuleId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSkuListPromotionRulesSkuListPromotionRuleId200Response(val *GETSkuListPromotionRulesSkuListPromotionRuleId200Response) *NullableGETSkuListPromotionRulesSkuListPromotionRuleId200Response {
	return &NullableGETSkuListPromotionRulesSkuListPromotionRuleId200Response{value: val, isSet: true}
}

func (v NullableGETSkuListPromotionRulesSkuListPromotionRuleId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSkuListPromotionRulesSkuListPromotionRuleId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
