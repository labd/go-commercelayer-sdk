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

// checks if the SkuOptionCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuOptionCreateDataRelationships{}

// SkuOptionCreateDataRelationships struct for SkuOptionCreateDataRelationships
type SkuOptionCreateDataRelationships struct {
	Market *BundleCreateDataRelationshipsMarket `json:"market,omitempty"`
	Tags   *AddressCreateDataRelationshipsTags  `json:"tags,omitempty"`
}

// NewSkuOptionCreateDataRelationships instantiates a new SkuOptionCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuOptionCreateDataRelationships() *SkuOptionCreateDataRelationships {
	this := SkuOptionCreateDataRelationships{}
	return &this
}

// NewSkuOptionCreateDataRelationshipsWithDefaults instantiates a new SkuOptionCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuOptionCreateDataRelationshipsWithDefaults() *SkuOptionCreateDataRelationships {
	this := SkuOptionCreateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *SkuOptionCreateDataRelationships) GetMarket() BundleCreateDataRelationshipsMarket {
	if o == nil || IsNil(o.Market) {
		var ret BundleCreateDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuOptionCreateDataRelationships) GetMarketOk() (*BundleCreateDataRelationshipsMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *SkuOptionCreateDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given BundleCreateDataRelationshipsMarket and assigns it to the Market field.
func (o *SkuOptionCreateDataRelationships) SetMarket(v BundleCreateDataRelationshipsMarket) {
	o.Market = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SkuOptionCreateDataRelationships) GetTags() AddressCreateDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret AddressCreateDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuOptionCreateDataRelationships) GetTagsOk() (*AddressCreateDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SkuOptionCreateDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given AddressCreateDataRelationshipsTags and assigns it to the Tags field.
func (o *SkuOptionCreateDataRelationships) SetTags(v AddressCreateDataRelationshipsTags) {
	o.Tags = &v
}

func (o SkuOptionCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuOptionCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableSkuOptionCreateDataRelationships struct {
	value *SkuOptionCreateDataRelationships
	isSet bool
}

func (v NullableSkuOptionCreateDataRelationships) Get() *SkuOptionCreateDataRelationships {
	return v.value
}

func (v *NullableSkuOptionCreateDataRelationships) Set(val *SkuOptionCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuOptionCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuOptionCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuOptionCreateDataRelationships(val *SkuOptionCreateDataRelationships) *NullableSkuOptionCreateDataRelationships {
	return &NullableSkuOptionCreateDataRelationships{value: val, isSet: true}
}

func (v NullableSkuOptionCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuOptionCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
