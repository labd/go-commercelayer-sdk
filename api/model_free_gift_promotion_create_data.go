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

// checks if the FreeGiftPromotionCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FreeGiftPromotionCreateData{}

// FreeGiftPromotionCreateData struct for FreeGiftPromotionCreateData
type FreeGiftPromotionCreateData struct {
	// The resource's type
	Type          interface{}                                     `json:"type"`
	Attributes    POSTFreeGiftPromotions201ResponseDataAttributes `json:"attributes"`
	Relationships *BuyXPayYPromotionCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewFreeGiftPromotionCreateData instantiates a new FreeGiftPromotionCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeGiftPromotionCreateData(type_ interface{}, attributes POSTFreeGiftPromotions201ResponseDataAttributes) *FreeGiftPromotionCreateData {
	this := FreeGiftPromotionCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewFreeGiftPromotionCreateDataWithDefaults instantiates a new FreeGiftPromotionCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreeGiftPromotionCreateDataWithDefaults() *FreeGiftPromotionCreateData {
	this := FreeGiftPromotionCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FreeGiftPromotionCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FreeGiftPromotionCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FreeGiftPromotionCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *FreeGiftPromotionCreateData) GetAttributes() POSTFreeGiftPromotions201ResponseDataAttributes {
	if o == nil {
		var ret POSTFreeGiftPromotions201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *FreeGiftPromotionCreateData) GetAttributesOk() (*POSTFreeGiftPromotions201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *FreeGiftPromotionCreateData) SetAttributes(v POSTFreeGiftPromotions201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *FreeGiftPromotionCreateData) GetRelationships() BuyXPayYPromotionCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret BuyXPayYPromotionCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeGiftPromotionCreateData) GetRelationshipsOk() (*BuyXPayYPromotionCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *FreeGiftPromotionCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BuyXPayYPromotionCreateDataRelationships and assigns it to the Relationships field.
func (o *FreeGiftPromotionCreateData) SetRelationships(v BuyXPayYPromotionCreateDataRelationships) {
	o.Relationships = &v
}

func (o FreeGiftPromotionCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FreeGiftPromotionCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableFreeGiftPromotionCreateData struct {
	value *FreeGiftPromotionCreateData
	isSet bool
}

func (v NullableFreeGiftPromotionCreateData) Get() *FreeGiftPromotionCreateData {
	return v.value
}

func (v *NullableFreeGiftPromotionCreateData) Set(val *FreeGiftPromotionCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableFreeGiftPromotionCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableFreeGiftPromotionCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreeGiftPromotionCreateData(val *FreeGiftPromotionCreateData) *NullableFreeGiftPromotionCreateData {
	return &NullableFreeGiftPromotionCreateData{value: val, isSet: true}
}

func (v NullableFreeGiftPromotionCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreeGiftPromotionCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
