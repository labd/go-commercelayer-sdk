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

// FreeShippingPromotionCreateData struct for FreeShippingPromotionCreateData
type FreeShippingPromotionCreateData struct {
	// The resource's type
	Type          string                                              `json:"type"`
	Attributes    POSTFreeShippingPromotions201ResponseDataAttributes `json:"attributes"`
	Relationships *ExternalPromotionCreateDataRelationships           `json:"relationships,omitempty"`
}

// NewFreeShippingPromotionCreateData instantiates a new FreeShippingPromotionCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeShippingPromotionCreateData(type_ string, attributes POSTFreeShippingPromotions201ResponseDataAttributes) *FreeShippingPromotionCreateData {
	this := FreeShippingPromotionCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewFreeShippingPromotionCreateDataWithDefaults instantiates a new FreeShippingPromotionCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreeShippingPromotionCreateDataWithDefaults() *FreeShippingPromotionCreateData {
	this := FreeShippingPromotionCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *FreeShippingPromotionCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FreeShippingPromotionCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *FreeShippingPromotionCreateData) GetAttributes() POSTFreeShippingPromotions201ResponseDataAttributes {
	if o == nil {
		var ret POSTFreeShippingPromotions201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionCreateData) GetAttributesOk() (*POSTFreeShippingPromotions201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *FreeShippingPromotionCreateData) SetAttributes(v POSTFreeShippingPromotions201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *FreeShippingPromotionCreateData) GetRelationships() ExternalPromotionCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ExternalPromotionCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionCreateData) GetRelationshipsOk() (*ExternalPromotionCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *FreeShippingPromotionCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ExternalPromotionCreateDataRelationships and assigns it to the Relationships field.
func (o *FreeShippingPromotionCreateData) SetRelationships(v ExternalPromotionCreateDataRelationships) {
	o.Relationships = &v
}

func (o FreeShippingPromotionCreateData) MarshalJSON() ([]byte, error) {
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

type NullableFreeShippingPromotionCreateData struct {
	value *FreeShippingPromotionCreateData
	isSet bool
}

func (v NullableFreeShippingPromotionCreateData) Get() *FreeShippingPromotionCreateData {
	return v.value
}

func (v *NullableFreeShippingPromotionCreateData) Set(val *FreeShippingPromotionCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableFreeShippingPromotionCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableFreeShippingPromotionCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreeShippingPromotionCreateData(val *FreeShippingPromotionCreateData) *NullableFreeShippingPromotionCreateData {
	return &NullableFreeShippingPromotionCreateData{value: val, isSet: true}
}

func (v NullableFreeShippingPromotionCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreeShippingPromotionCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
