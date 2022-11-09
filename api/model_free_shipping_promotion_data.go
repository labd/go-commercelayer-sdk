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

// FreeShippingPromotionData struct for FreeShippingPromotionData
type FreeShippingPromotionData struct {
	// The resource's type
	Type          string                                                  `json:"type"`
	Attributes    GETFreeShippingPromotions200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *ExternalPromotionDataRelationships                     `json:"relationships,omitempty"`
}

// NewFreeShippingPromotionData instantiates a new FreeShippingPromotionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeShippingPromotionData(type_ string, attributes GETFreeShippingPromotions200ResponseDataInnerAttributes) *FreeShippingPromotionData {
	this := FreeShippingPromotionData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewFreeShippingPromotionDataWithDefaults instantiates a new FreeShippingPromotionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreeShippingPromotionDataWithDefaults() *FreeShippingPromotionData {
	this := FreeShippingPromotionData{}
	return &this
}

// GetType returns the Type field value
func (o *FreeShippingPromotionData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FreeShippingPromotionData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *FreeShippingPromotionData) GetAttributes() GETFreeShippingPromotions200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETFreeShippingPromotions200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionData) GetAttributesOk() (*GETFreeShippingPromotions200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *FreeShippingPromotionData) SetAttributes(v GETFreeShippingPromotions200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *FreeShippingPromotionData) GetRelationships() ExternalPromotionDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ExternalPromotionDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionData) GetRelationshipsOk() (*ExternalPromotionDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *FreeShippingPromotionData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ExternalPromotionDataRelationships and assigns it to the Relationships field.
func (o *FreeShippingPromotionData) SetRelationships(v ExternalPromotionDataRelationships) {
	o.Relationships = &v
}

func (o FreeShippingPromotionData) MarshalJSON() ([]byte, error) {
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

type NullableFreeShippingPromotionData struct {
	value *FreeShippingPromotionData
	isSet bool
}

func (v NullableFreeShippingPromotionData) Get() *FreeShippingPromotionData {
	return v.value
}

func (v *NullableFreeShippingPromotionData) Set(val *FreeShippingPromotionData) {
	v.value = val
	v.isSet = true
}

func (v NullableFreeShippingPromotionData) IsSet() bool {
	return v.isSet
}

func (v *NullableFreeShippingPromotionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreeShippingPromotionData(val *FreeShippingPromotionData) *NullableFreeShippingPromotionData {
	return &NullableFreeShippingPromotionData{value: val, isSet: true}
}

func (v NullableFreeShippingPromotionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreeShippingPromotionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
