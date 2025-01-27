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

// checks if the ReturnLineItemData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnLineItemData{}

// ReturnLineItemData struct for ReturnLineItemData
type ReturnLineItemData struct {
	// The resource's type
	Type          interface{}                                                 `json:"type"`
	Attributes    GETReturnLineItemsReturnLineItemId200ResponseDataAttributes `json:"attributes"`
	Relationships *ReturnLineItemDataRelationships                            `json:"relationships,omitempty"`
}

// NewReturnLineItemData instantiates a new ReturnLineItemData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnLineItemData(type_ interface{}, attributes GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) *ReturnLineItemData {
	this := ReturnLineItemData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewReturnLineItemDataWithDefaults instantiates a new ReturnLineItemData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnLineItemDataWithDefaults() *ReturnLineItemData {
	this := ReturnLineItemData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ReturnLineItemData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReturnLineItemData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReturnLineItemData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ReturnLineItemData) GetAttributes() GETReturnLineItemsReturnLineItemId200ResponseDataAttributes {
	if o == nil {
		var ret GETReturnLineItemsReturnLineItemId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ReturnLineItemData) GetAttributesOk() (*GETReturnLineItemsReturnLineItemId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ReturnLineItemData) SetAttributes(v GETReturnLineItemsReturnLineItemId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ReturnLineItemData) GetRelationships() ReturnLineItemDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ReturnLineItemDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLineItemData) GetRelationshipsOk() (*ReturnLineItemDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ReturnLineItemData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ReturnLineItemDataRelationships and assigns it to the Relationships field.
func (o *ReturnLineItemData) SetRelationships(v ReturnLineItemDataRelationships) {
	o.Relationships = &v
}

func (o ReturnLineItemData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnLineItemData) ToMap() (map[string]interface{}, error) {
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

type NullableReturnLineItemData struct {
	value *ReturnLineItemData
	isSet bool
}

func (v NullableReturnLineItemData) Get() *ReturnLineItemData {
	return v.value
}

func (v *NullableReturnLineItemData) Set(val *ReturnLineItemData) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnLineItemData) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnLineItemData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnLineItemData(val *ReturnLineItemData) *NullableReturnLineItemData {
	return &NullableReturnLineItemData{value: val, isSet: true}
}

func (v NullableReturnLineItemData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnLineItemData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
