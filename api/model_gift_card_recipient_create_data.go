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

// GiftCardRecipientCreateData struct for GiftCardRecipientCreateData
type GiftCardRecipientCreateData struct {
	// The resource's type
	Type          string                                  `json:"type"`
	Attributes    CouponRecipientCreateDataAttributes     `json:"attributes"`
	Relationships *CouponRecipientCreateDataRelationships `json:"relationships,omitempty"`
}

// NewGiftCardRecipientCreateData instantiates a new GiftCardRecipientCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardRecipientCreateData(type_ string, attributes CouponRecipientCreateDataAttributes) *GiftCardRecipientCreateData {
	this := GiftCardRecipientCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewGiftCardRecipientCreateDataWithDefaults instantiates a new GiftCardRecipientCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardRecipientCreateDataWithDefaults() *GiftCardRecipientCreateData {
	this := GiftCardRecipientCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *GiftCardRecipientCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GiftCardRecipientCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GiftCardRecipientCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *GiftCardRecipientCreateData) GetAttributes() CouponRecipientCreateDataAttributes {
	if o == nil {
		var ret CouponRecipientCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GiftCardRecipientCreateData) GetAttributesOk() (*CouponRecipientCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GiftCardRecipientCreateData) SetAttributes(v CouponRecipientCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GiftCardRecipientCreateData) GetRelationships() CouponRecipientCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CouponRecipientCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardRecipientCreateData) GetRelationshipsOk() (*CouponRecipientCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GiftCardRecipientCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CouponRecipientCreateDataRelationships and assigns it to the Relationships field.
func (o *GiftCardRecipientCreateData) SetRelationships(v CouponRecipientCreateDataRelationships) {
	o.Relationships = &v
}

func (o GiftCardRecipientCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableGiftCardRecipientCreateData struct {
	value *GiftCardRecipientCreateData
	isSet bool
}

func (v NullableGiftCardRecipientCreateData) Get() *GiftCardRecipientCreateData {
	return v.value
}

func (v *NullableGiftCardRecipientCreateData) Set(val *GiftCardRecipientCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardRecipientCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardRecipientCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardRecipientCreateData(val *GiftCardRecipientCreateData) *NullableGiftCardRecipientCreateData {
	return &NullableGiftCardRecipientCreateData{value: val, isSet: true}
}

func (v NullableGiftCardRecipientCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardRecipientCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
