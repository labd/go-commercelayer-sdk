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

// LineItemData struct for LineItemData
type LineItemData struct {
	// The resource's type
	Type          string                                     `json:"type"`
	Attributes    GETLineItems200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *LineItemDataRelationships                 `json:"relationships,omitempty"`
}

// NewLineItemData instantiates a new LineItemData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemData(type_ string, attributes GETLineItems200ResponseDataInnerAttributes) *LineItemData {
	this := LineItemData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewLineItemDataWithDefaults instantiates a new LineItemData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemDataWithDefaults() *LineItemData {
	this := LineItemData{}
	return &this
}

// GetType returns the Type field value
func (o *LineItemData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LineItemData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LineItemData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *LineItemData) GetAttributes() GETLineItems200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETLineItems200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *LineItemData) GetAttributesOk() (*GETLineItems200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *LineItemData) SetAttributes(v GETLineItems200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *LineItemData) GetRelationships() LineItemDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret LineItemDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemData) GetRelationshipsOk() (*LineItemDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *LineItemData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given LineItemDataRelationships and assigns it to the Relationships field.
func (o *LineItemData) SetRelationships(v LineItemDataRelationships) {
	o.Relationships = &v
}

func (o LineItemData) MarshalJSON() ([]byte, error) {
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

type NullableLineItemData struct {
	value *LineItemData
	isSet bool
}

func (v NullableLineItemData) Get() *LineItemData {
	return v.value
}

func (v *NullableLineItemData) Set(val *LineItemData) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemData) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemData(val *LineItemData) *NullableLineItemData {
	return &NullableLineItemData{value: val, isSet: true}
}

func (v NullableLineItemData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
