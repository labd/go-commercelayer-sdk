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

// checks if the FreeShippingPromotionCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FreeShippingPromotionCreateData{}

// FreeShippingPromotionCreateData struct for FreeShippingPromotionCreateData
type FreeShippingPromotionCreateData struct {
	// The resource's type
	Type          interface{}                                         `json:"type"`
	Attributes    POSTFreeShippingPromotions201ResponseDataAttributes `json:"attributes"`
	Relationships *BuyXPayYPromotionUpdateDataRelationships           `json:"relationships,omitempty"`
}

// NewFreeShippingPromotionCreateData instantiates a new FreeShippingPromotionCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeShippingPromotionCreateData(type_ interface{}, attributes POSTFreeShippingPromotions201ResponseDataAttributes) *FreeShippingPromotionCreateData {
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
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FreeShippingPromotionCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FreeShippingPromotionCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FreeShippingPromotionCreateData) SetType(v interface{}) {
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
func (o *FreeShippingPromotionCreateData) GetRelationships() BuyXPayYPromotionUpdateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret BuyXPayYPromotionUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionCreateData) GetRelationshipsOk() (*BuyXPayYPromotionUpdateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *FreeShippingPromotionCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BuyXPayYPromotionUpdateDataRelationships and assigns it to the Relationships field.
func (o *FreeShippingPromotionCreateData) SetRelationships(v BuyXPayYPromotionUpdateDataRelationships) {
	o.Relationships = &v
}

func (o FreeShippingPromotionCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FreeShippingPromotionCreateData) ToMap() (map[string]interface{}, error) {
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
