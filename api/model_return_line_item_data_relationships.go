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

// ReturnLineItemDataRelationships struct for ReturnLineItemDataRelationships
type ReturnLineItemDataRelationships struct {
	Return   *CustomerDataRelationshipsReturns        `json:"return,omitempty"`
	LineItem *LineItemOptionDataRelationshipsLineItem `json:"line_item,omitempty"`
}

// NewReturnLineItemDataRelationships instantiates a new ReturnLineItemDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnLineItemDataRelationships() *ReturnLineItemDataRelationships {
	this := ReturnLineItemDataRelationships{}
	return &this
}

// NewReturnLineItemDataRelationshipsWithDefaults instantiates a new ReturnLineItemDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnLineItemDataRelationshipsWithDefaults() *ReturnLineItemDataRelationships {
	this := ReturnLineItemDataRelationships{}
	return &this
}

// GetReturn returns the Return field value if set, zero value otherwise.
func (o *ReturnLineItemDataRelationships) GetReturn() CustomerDataRelationshipsReturns {
	if o == nil || o.Return == nil {
		var ret CustomerDataRelationshipsReturns
		return ret
	}
	return *o.Return
}

// GetReturnOk returns a tuple with the Return field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLineItemDataRelationships) GetReturnOk() (*CustomerDataRelationshipsReturns, bool) {
	if o == nil || o.Return == nil {
		return nil, false
	}
	return o.Return, true
}

// HasReturn returns a boolean if a field has been set.
func (o *ReturnLineItemDataRelationships) HasReturn() bool {
	if o != nil && o.Return != nil {
		return true
	}

	return false
}

// SetReturn gets a reference to the given CustomerDataRelationshipsReturns and assigns it to the Return field.
func (o *ReturnLineItemDataRelationships) SetReturn(v CustomerDataRelationshipsReturns) {
	o.Return = &v
}

// GetLineItem returns the LineItem field value if set, zero value otherwise.
func (o *ReturnLineItemDataRelationships) GetLineItem() LineItemOptionDataRelationshipsLineItem {
	if o == nil || o.LineItem == nil {
		var ret LineItemOptionDataRelationshipsLineItem
		return ret
	}
	return *o.LineItem
}

// GetLineItemOk returns a tuple with the LineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLineItemDataRelationships) GetLineItemOk() (*LineItemOptionDataRelationshipsLineItem, bool) {
	if o == nil || o.LineItem == nil {
		return nil, false
	}
	return o.LineItem, true
}

// HasLineItem returns a boolean if a field has been set.
func (o *ReturnLineItemDataRelationships) HasLineItem() bool {
	if o != nil && o.LineItem != nil {
		return true
	}

	return false
}

// SetLineItem gets a reference to the given LineItemOptionDataRelationshipsLineItem and assigns it to the LineItem field.
func (o *ReturnLineItemDataRelationships) SetLineItem(v LineItemOptionDataRelationshipsLineItem) {
	o.LineItem = &v
}

func (o ReturnLineItemDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Return != nil {
		toSerialize["return"] = o.Return
	}
	if o.LineItem != nil {
		toSerialize["line_item"] = o.LineItem
	}
	return json.Marshal(toSerialize)
}

type NullableReturnLineItemDataRelationships struct {
	value *ReturnLineItemDataRelationships
	isSet bool
}

func (v NullableReturnLineItemDataRelationships) Get() *ReturnLineItemDataRelationships {
	return v.value
}

func (v *NullableReturnLineItemDataRelationships) Set(val *ReturnLineItemDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnLineItemDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnLineItemDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnLineItemDataRelationships(val *ReturnLineItemDataRelationships) *NullableReturnLineItemDataRelationships {
	return &NullableReturnLineItemDataRelationships{value: val, isSet: true}
}

func (v NullableReturnLineItemDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnLineItemDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
