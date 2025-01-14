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

// checks if the LineItemCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItemCreateDataRelationships{}

// LineItemCreateDataRelationships struct for LineItemCreateDataRelationships
type LineItemCreateDataRelationships struct {
	Order AdyenPaymentCreateDataRelationshipsOrder `json:"order"`
	Item  *LineItemCreateDataRelationshipsItem     `json:"item,omitempty"`
	Tags  *AddressCreateDataRelationshipsTags      `json:"tags,omitempty"`
}

// NewLineItemCreateDataRelationships instantiates a new LineItemCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemCreateDataRelationships(order AdyenPaymentCreateDataRelationshipsOrder) *LineItemCreateDataRelationships {
	this := LineItemCreateDataRelationships{}
	this.Order = order
	return &this
}

// NewLineItemCreateDataRelationshipsWithDefaults instantiates a new LineItemCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemCreateDataRelationshipsWithDefaults() *LineItemCreateDataRelationships {
	this := LineItemCreateDataRelationships{}
	return &this
}

// GetOrder returns the Order field value
func (o *LineItemCreateDataRelationships) GetOrder() AdyenPaymentCreateDataRelationshipsOrder {
	if o == nil {
		var ret AdyenPaymentCreateDataRelationshipsOrder
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *LineItemCreateDataRelationships) GetOrderOk() (*AdyenPaymentCreateDataRelationshipsOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *LineItemCreateDataRelationships) SetOrder(v AdyenPaymentCreateDataRelationshipsOrder) {
	o.Order = v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *LineItemCreateDataRelationships) GetItem() LineItemCreateDataRelationshipsItem {
	if o == nil || IsNil(o.Item) {
		var ret LineItemCreateDataRelationshipsItem
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemCreateDataRelationships) GetItemOk() (*LineItemCreateDataRelationshipsItem, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *LineItemCreateDataRelationships) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given LineItemCreateDataRelationshipsItem and assigns it to the Item field.
func (o *LineItemCreateDataRelationships) SetItem(v LineItemCreateDataRelationshipsItem) {
	o.Item = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LineItemCreateDataRelationships) GetTags() AddressCreateDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret AddressCreateDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemCreateDataRelationships) GetTagsOk() (*AddressCreateDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LineItemCreateDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given AddressCreateDataRelationshipsTags and assigns it to the Tags field.
func (o *LineItemCreateDataRelationships) SetTags(v AddressCreateDataRelationshipsTags) {
	o.Tags = &v
}

func (o LineItemCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItemCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["order"] = o.Order
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableLineItemCreateDataRelationships struct {
	value *LineItemCreateDataRelationships
	isSet bool
}

func (v NullableLineItemCreateDataRelationships) Get() *LineItemCreateDataRelationships {
	return v.value
}

func (v *NullableLineItemCreateDataRelationships) Set(val *LineItemCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemCreateDataRelationships(val *LineItemCreateDataRelationships) *NullableLineItemCreateDataRelationships {
	return &NullableLineItemCreateDataRelationships{value: val, isSet: true}
}

func (v NullableLineItemCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
