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

// ExternalPromotionCreateDataRelationshipsPromotionRules struct for ExternalPromotionCreateDataRelationshipsPromotionRules
type ExternalPromotionCreateDataRelationshipsPromotionRules struct {
	Data ExternalPromotionDataRelationshipsPromotionRulesData `json:"data"`
}

// NewExternalPromotionCreateDataRelationshipsPromotionRules instantiates a new ExternalPromotionCreateDataRelationshipsPromotionRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPromotionCreateDataRelationshipsPromotionRules(data ExternalPromotionDataRelationshipsPromotionRulesData) *ExternalPromotionCreateDataRelationshipsPromotionRules {
	this := ExternalPromotionCreateDataRelationshipsPromotionRules{}
	this.Data = data
	return &this
}

// NewExternalPromotionCreateDataRelationshipsPromotionRulesWithDefaults instantiates a new ExternalPromotionCreateDataRelationshipsPromotionRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPromotionCreateDataRelationshipsPromotionRulesWithDefaults() *ExternalPromotionCreateDataRelationshipsPromotionRules {
	this := ExternalPromotionCreateDataRelationshipsPromotionRules{}
	return &this
}

// GetData returns the Data field value
func (o *ExternalPromotionCreateDataRelationshipsPromotionRules) GetData() ExternalPromotionDataRelationshipsPromotionRulesData {
	if o == nil {
		var ret ExternalPromotionDataRelationshipsPromotionRulesData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExternalPromotionCreateDataRelationshipsPromotionRules) GetDataOk() (*ExternalPromotionDataRelationshipsPromotionRulesData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ExternalPromotionCreateDataRelationshipsPromotionRules) SetData(v ExternalPromotionDataRelationshipsPromotionRulesData) {
	o.Data = v
}

func (o ExternalPromotionCreateDataRelationshipsPromotionRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableExternalPromotionCreateDataRelationshipsPromotionRules struct {
	value *ExternalPromotionCreateDataRelationshipsPromotionRules
	isSet bool
}

func (v NullableExternalPromotionCreateDataRelationshipsPromotionRules) Get() *ExternalPromotionCreateDataRelationshipsPromotionRules {
	return v.value
}

func (v *NullableExternalPromotionCreateDataRelationshipsPromotionRules) Set(val *ExternalPromotionCreateDataRelationshipsPromotionRules) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPromotionCreateDataRelationshipsPromotionRules) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPromotionCreateDataRelationshipsPromotionRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPromotionCreateDataRelationshipsPromotionRules(val *ExternalPromotionCreateDataRelationshipsPromotionRules) *NullableExternalPromotionCreateDataRelationshipsPromotionRules {
	return &NullableExternalPromotionCreateDataRelationshipsPromotionRules{value: val, isSet: true}
}

func (v NullableExternalPromotionCreateDataRelationshipsPromotionRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPromotionCreateDataRelationshipsPromotionRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
