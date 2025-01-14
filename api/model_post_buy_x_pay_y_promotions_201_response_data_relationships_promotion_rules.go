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

// checks if the POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules{}

// POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules struct for POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules
type POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks               `json:"links,omitempty"`
	Data  *POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRulesData `json:"data,omitempty"`
}

// NewPOSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules instantiates a new POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules() *POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules {
	this := POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules{}
	return &this
}

// NewPOSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRulesWithDefaults instantiates a new POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRulesWithDefaults() *POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules {
	this := POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules) GetData() POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRulesData {
	if o == nil || IsNil(o.Data) {
		var ret POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRulesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules) GetDataOk() (*POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRulesData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRulesData and assigns it to the Data field.
func (o *POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules) SetData(v POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRulesData) {
	o.Data = &v
}

func (o POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules struct {
	value *POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules
	isSet bool
}

func (v NullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules) Get() *POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules {
	return v.value
}

func (v *NullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules) Set(val *POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules(val *POSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules) *NullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules {
	return &NullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules{value: val, isSet: true}
}

func (v NullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsPromotionRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
