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

// checks if the GiftCardCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftCardCreateDataRelationships{}

// GiftCardCreateDataRelationships struct for GiftCardCreateDataRelationships
type GiftCardCreateDataRelationships struct {
	Market            *BundleCreateDataRelationshipsMarket              `json:"market,omitempty"`
	GiftCardRecipient *GiftCardCreateDataRelationshipsGiftCardRecipient `json:"gift_card_recipient,omitempty"`
	Tags              *AddressCreateDataRelationshipsTags               `json:"tags,omitempty"`
}

// NewGiftCardCreateDataRelationships instantiates a new GiftCardCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardCreateDataRelationships() *GiftCardCreateDataRelationships {
	this := GiftCardCreateDataRelationships{}
	return &this
}

// NewGiftCardCreateDataRelationshipsWithDefaults instantiates a new GiftCardCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardCreateDataRelationshipsWithDefaults() *GiftCardCreateDataRelationships {
	this := GiftCardCreateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *GiftCardCreateDataRelationships) GetMarket() BundleCreateDataRelationshipsMarket {
	if o == nil || IsNil(o.Market) {
		var ret BundleCreateDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardCreateDataRelationships) GetMarketOk() (*BundleCreateDataRelationshipsMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *GiftCardCreateDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given BundleCreateDataRelationshipsMarket and assigns it to the Market field.
func (o *GiftCardCreateDataRelationships) SetMarket(v BundleCreateDataRelationshipsMarket) {
	o.Market = &v
}

// GetGiftCardRecipient returns the GiftCardRecipient field value if set, zero value otherwise.
func (o *GiftCardCreateDataRelationships) GetGiftCardRecipient() GiftCardCreateDataRelationshipsGiftCardRecipient {
	if o == nil || IsNil(o.GiftCardRecipient) {
		var ret GiftCardCreateDataRelationshipsGiftCardRecipient
		return ret
	}
	return *o.GiftCardRecipient
}

// GetGiftCardRecipientOk returns a tuple with the GiftCardRecipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardCreateDataRelationships) GetGiftCardRecipientOk() (*GiftCardCreateDataRelationshipsGiftCardRecipient, bool) {
	if o == nil || IsNil(o.GiftCardRecipient) {
		return nil, false
	}
	return o.GiftCardRecipient, true
}

// HasGiftCardRecipient returns a boolean if a field has been set.
func (o *GiftCardCreateDataRelationships) HasGiftCardRecipient() bool {
	if o != nil && !IsNil(o.GiftCardRecipient) {
		return true
	}

	return false
}

// SetGiftCardRecipient gets a reference to the given GiftCardCreateDataRelationshipsGiftCardRecipient and assigns it to the GiftCardRecipient field.
func (o *GiftCardCreateDataRelationships) SetGiftCardRecipient(v GiftCardCreateDataRelationshipsGiftCardRecipient) {
	o.GiftCardRecipient = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GiftCardCreateDataRelationships) GetTags() AddressCreateDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret AddressCreateDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardCreateDataRelationships) GetTagsOk() (*AddressCreateDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GiftCardCreateDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given AddressCreateDataRelationshipsTags and assigns it to the Tags field.
func (o *GiftCardCreateDataRelationships) SetTags(v AddressCreateDataRelationshipsTags) {
	o.Tags = &v
}

func (o GiftCardCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftCardCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.GiftCardRecipient) {
		toSerialize["gift_card_recipient"] = o.GiftCardRecipient
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableGiftCardCreateDataRelationships struct {
	value *GiftCardCreateDataRelationships
	isSet bool
}

func (v NullableGiftCardCreateDataRelationships) Get() *GiftCardCreateDataRelationships {
	return v.value
}

func (v *NullableGiftCardCreateDataRelationships) Set(val *GiftCardCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardCreateDataRelationships(val *GiftCardCreateDataRelationships) *NullableGiftCardCreateDataRelationships {
	return &NullableGiftCardCreateDataRelationships{value: val, isSet: true}
}

func (v NullableGiftCardCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
