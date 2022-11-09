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

// ShippingWeightTierCreateData struct for ShippingWeightTierCreateData
type ShippingWeightTierCreateData struct {
	// The resource's type
	Type          string                                           `json:"type"`
	Attributes    POSTShippingWeightTiers201ResponseDataAttributes `json:"attributes"`
	Relationships *ShippingWeightTierCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewShippingWeightTierCreateData instantiates a new ShippingWeightTierCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingWeightTierCreateData(type_ string, attributes POSTShippingWeightTiers201ResponseDataAttributes) *ShippingWeightTierCreateData {
	this := ShippingWeightTierCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewShippingWeightTierCreateDataWithDefaults instantiates a new ShippingWeightTierCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingWeightTierCreateDataWithDefaults() *ShippingWeightTierCreateData {
	this := ShippingWeightTierCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *ShippingWeightTierCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ShippingWeightTierCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShippingWeightTierCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ShippingWeightTierCreateData) GetAttributes() POSTShippingWeightTiers201ResponseDataAttributes {
	if o == nil {
		var ret POSTShippingWeightTiers201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ShippingWeightTierCreateData) GetAttributesOk() (*POSTShippingWeightTiers201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ShippingWeightTierCreateData) SetAttributes(v POSTShippingWeightTiers201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ShippingWeightTierCreateData) GetRelationships() ShippingWeightTierCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ShippingWeightTierCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingWeightTierCreateData) GetRelationshipsOk() (*ShippingWeightTierCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ShippingWeightTierCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ShippingWeightTierCreateDataRelationships and assigns it to the Relationships field.
func (o *ShippingWeightTierCreateData) SetRelationships(v ShippingWeightTierCreateDataRelationships) {
	o.Relationships = &v
}

func (o ShippingWeightTierCreateData) MarshalJSON() ([]byte, error) {
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

type NullableShippingWeightTierCreateData struct {
	value *ShippingWeightTierCreateData
	isSet bool
}

func (v NullableShippingWeightTierCreateData) Get() *ShippingWeightTierCreateData {
	return v.value
}

func (v *NullableShippingWeightTierCreateData) Set(val *ShippingWeightTierCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingWeightTierCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingWeightTierCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingWeightTierCreateData(val *ShippingWeightTierCreateData) *NullableShippingWeightTierCreateData {
	return &NullableShippingWeightTierCreateData{value: val, isSet: true}
}

func (v NullableShippingWeightTierCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingWeightTierCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
