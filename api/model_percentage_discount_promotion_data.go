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

// PercentageDiscountPromotionData struct for PercentageDiscountPromotionData
type PercentageDiscountPromotionData struct {
	// The resource's type
	Type          string                                                        `json:"type"`
	Attributes    GETPercentageDiscountPromotions200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *FixedPricePromotionDataRelationships                         `json:"relationships,omitempty"`
}

// NewPercentageDiscountPromotionData instantiates a new PercentageDiscountPromotionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPercentageDiscountPromotionData(type_ string, attributes GETPercentageDiscountPromotions200ResponseDataInnerAttributes) *PercentageDiscountPromotionData {
	this := PercentageDiscountPromotionData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPercentageDiscountPromotionDataWithDefaults instantiates a new PercentageDiscountPromotionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPercentageDiscountPromotionDataWithDefaults() *PercentageDiscountPromotionData {
	this := PercentageDiscountPromotionData{}
	return &this
}

// GetType returns the Type field value
func (o *PercentageDiscountPromotionData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PercentageDiscountPromotionData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PercentageDiscountPromotionData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PercentageDiscountPromotionData) GetAttributes() GETPercentageDiscountPromotions200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETPercentageDiscountPromotions200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PercentageDiscountPromotionData) GetAttributesOk() (*GETPercentageDiscountPromotions200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PercentageDiscountPromotionData) SetAttributes(v GETPercentageDiscountPromotions200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PercentageDiscountPromotionData) GetRelationships() FixedPricePromotionDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret FixedPricePromotionDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PercentageDiscountPromotionData) GetRelationshipsOk() (*FixedPricePromotionDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PercentageDiscountPromotionData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given FixedPricePromotionDataRelationships and assigns it to the Relationships field.
func (o *PercentageDiscountPromotionData) SetRelationships(v FixedPricePromotionDataRelationships) {
	o.Relationships = &v
}

func (o PercentageDiscountPromotionData) MarshalJSON() ([]byte, error) {
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

type NullablePercentageDiscountPromotionData struct {
	value *PercentageDiscountPromotionData
	isSet bool
}

func (v NullablePercentageDiscountPromotionData) Get() *PercentageDiscountPromotionData {
	return v.value
}

func (v *NullablePercentageDiscountPromotionData) Set(val *PercentageDiscountPromotionData) {
	v.value = val
	v.isSet = true
}

func (v NullablePercentageDiscountPromotionData) IsSet() bool {
	return v.isSet
}

func (v *NullablePercentageDiscountPromotionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePercentageDiscountPromotionData(val *PercentageDiscountPromotionData) *NullablePercentageDiscountPromotionData {
	return &NullablePercentageDiscountPromotionData{value: val, isSet: true}
}

func (v NullablePercentageDiscountPromotionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePercentageDiscountPromotionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
