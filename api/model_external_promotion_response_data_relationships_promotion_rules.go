/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ExternalPromotionResponseDataRelationshipsPromotionRules struct for ExternalPromotionResponseDataRelationshipsPromotionRules
type ExternalPromotionResponseDataRelationshipsPromotionRules struct {
	Links *AddressResponseDataRelationshipsGeocoderLinks                      `json:"links,omitempty"`
	Data  []ExternalPromotionResponseDataRelationshipsPromotionRulesDataInner `json:"data,omitempty"`
}

// NewExternalPromotionResponseDataRelationshipsPromotionRules instantiates a new ExternalPromotionResponseDataRelationshipsPromotionRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPromotionResponseDataRelationshipsPromotionRules() *ExternalPromotionResponseDataRelationshipsPromotionRules {
	this := ExternalPromotionResponseDataRelationshipsPromotionRules{}
	return &this
}

// NewExternalPromotionResponseDataRelationshipsPromotionRulesWithDefaults instantiates a new ExternalPromotionResponseDataRelationshipsPromotionRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPromotionResponseDataRelationshipsPromotionRulesWithDefaults() *ExternalPromotionResponseDataRelationshipsPromotionRules {
	this := ExternalPromotionResponseDataRelationshipsPromotionRules{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ExternalPromotionResponseDataRelationshipsPromotionRules) GetLinks() AddressResponseDataRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionResponseDataRelationshipsPromotionRules) GetLinksOk() (*AddressResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ExternalPromotionResponseDataRelationshipsPromotionRules) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *ExternalPromotionResponseDataRelationshipsPromotionRules) SetLinks(v AddressResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ExternalPromotionResponseDataRelationshipsPromotionRules) GetData() []ExternalPromotionResponseDataRelationshipsPromotionRulesDataInner {
	if o == nil || o.Data == nil {
		var ret []ExternalPromotionResponseDataRelationshipsPromotionRulesDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionResponseDataRelationshipsPromotionRules) GetDataOk() ([]ExternalPromotionResponseDataRelationshipsPromotionRulesDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ExternalPromotionResponseDataRelationshipsPromotionRules) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ExternalPromotionResponseDataRelationshipsPromotionRulesDataInner and assigns it to the Data field.
func (o *ExternalPromotionResponseDataRelationshipsPromotionRules) SetData(v []ExternalPromotionResponseDataRelationshipsPromotionRulesDataInner) {
	o.Data = v
}

func (o ExternalPromotionResponseDataRelationshipsPromotionRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableExternalPromotionResponseDataRelationshipsPromotionRules struct {
	value *ExternalPromotionResponseDataRelationshipsPromotionRules
	isSet bool
}

func (v NullableExternalPromotionResponseDataRelationshipsPromotionRules) Get() *ExternalPromotionResponseDataRelationshipsPromotionRules {
	return v.value
}

func (v *NullableExternalPromotionResponseDataRelationshipsPromotionRules) Set(val *ExternalPromotionResponseDataRelationshipsPromotionRules) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPromotionResponseDataRelationshipsPromotionRules) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPromotionResponseDataRelationshipsPromotionRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPromotionResponseDataRelationshipsPromotionRules(val *ExternalPromotionResponseDataRelationshipsPromotionRules) *NullableExternalPromotionResponseDataRelationshipsPromotionRules {
	return &NullableExternalPromotionResponseDataRelationshipsPromotionRules{value: val, isSet: true}
}

func (v NullableExternalPromotionResponseDataRelationshipsPromotionRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPromotionResponseDataRelationshipsPromotionRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
