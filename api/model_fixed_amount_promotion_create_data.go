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

// FixedAmountPromotionCreateData struct for FixedAmountPromotionCreateData
type FixedAmountPromotionCreateData struct {
	// The resource's type
	Type          string                                             `json:"type"`
	Attributes    POSTFixedAmountPromotions201ResponseDataAttributes `json:"attributes"`
	Relationships *ExternalPromotionCreateDataRelationships          `json:"relationships,omitempty"`
}

// NewFixedAmountPromotionCreateData instantiates a new FixedAmountPromotionCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixedAmountPromotionCreateData(type_ string, attributes POSTFixedAmountPromotions201ResponseDataAttributes) *FixedAmountPromotionCreateData {
	this := FixedAmountPromotionCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewFixedAmountPromotionCreateDataWithDefaults instantiates a new FixedAmountPromotionCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixedAmountPromotionCreateDataWithDefaults() *FixedAmountPromotionCreateData {
	this := FixedAmountPromotionCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *FixedAmountPromotionCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FixedAmountPromotionCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FixedAmountPromotionCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *FixedAmountPromotionCreateData) GetAttributes() POSTFixedAmountPromotions201ResponseDataAttributes {
	if o == nil {
		var ret POSTFixedAmountPromotions201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *FixedAmountPromotionCreateData) GetAttributesOk() (*POSTFixedAmountPromotions201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *FixedAmountPromotionCreateData) SetAttributes(v POSTFixedAmountPromotions201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *FixedAmountPromotionCreateData) GetRelationships() ExternalPromotionCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ExternalPromotionCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedAmountPromotionCreateData) GetRelationshipsOk() (*ExternalPromotionCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *FixedAmountPromotionCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ExternalPromotionCreateDataRelationships and assigns it to the Relationships field.
func (o *FixedAmountPromotionCreateData) SetRelationships(v ExternalPromotionCreateDataRelationships) {
	o.Relationships = &v
}

func (o FixedAmountPromotionCreateData) MarshalJSON() ([]byte, error) {
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

type NullableFixedAmountPromotionCreateData struct {
	value *FixedAmountPromotionCreateData
	isSet bool
}

func (v NullableFixedAmountPromotionCreateData) Get() *FixedAmountPromotionCreateData {
	return v.value
}

func (v *NullableFixedAmountPromotionCreateData) Set(val *FixedAmountPromotionCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableFixedAmountPromotionCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableFixedAmountPromotionCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixedAmountPromotionCreateData(val *FixedAmountPromotionCreateData) *NullableFixedAmountPromotionCreateData {
	return &NullableFixedAmountPromotionCreateData{value: val, isSet: true}
}

func (v NullableFixedAmountPromotionCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFixedAmountPromotionCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
