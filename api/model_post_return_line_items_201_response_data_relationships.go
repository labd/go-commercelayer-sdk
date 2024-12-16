/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTReturnLineItems201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTReturnLineItems201ResponseDataRelationships{}

// POSTReturnLineItems201ResponseDataRelationships struct for POSTReturnLineItems201ResponseDataRelationships
type POSTReturnLineItems201ResponseDataRelationships struct {
	Return   *GETCapturesCaptureId200ResponseDataRelationshipsReturn  `json:"return,omitempty"`
	LineItem *POSTLineItemOptions201ResponseDataRelationshipsLineItem `json:"line_item,omitempty"`
}

// NewPOSTReturnLineItems201ResponseDataRelationships instantiates a new POSTReturnLineItems201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTReturnLineItems201ResponseDataRelationships() *POSTReturnLineItems201ResponseDataRelationships {
	this := POSTReturnLineItems201ResponseDataRelationships{}
	return &this
}

// NewPOSTReturnLineItems201ResponseDataRelationshipsWithDefaults instantiates a new POSTReturnLineItems201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTReturnLineItems201ResponseDataRelationshipsWithDefaults() *POSTReturnLineItems201ResponseDataRelationships {
	this := POSTReturnLineItems201ResponseDataRelationships{}
	return &this
}

// GetReturn returns the Return field value if set, zero value otherwise.
func (o *POSTReturnLineItems201ResponseDataRelationships) GetReturn() GETCapturesCaptureId200ResponseDataRelationshipsReturn {
	if o == nil || IsNil(o.Return) {
		var ret GETCapturesCaptureId200ResponseDataRelationshipsReturn
		return ret
	}
	return *o.Return
}

// GetReturnOk returns a tuple with the Return field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturnLineItems201ResponseDataRelationships) GetReturnOk() (*GETCapturesCaptureId200ResponseDataRelationshipsReturn, bool) {
	if o == nil || IsNil(o.Return) {
		return nil, false
	}
	return o.Return, true
}

// HasReturn returns a boolean if a field has been set.
func (o *POSTReturnLineItems201ResponseDataRelationships) HasReturn() bool {
	if o != nil && !IsNil(o.Return) {
		return true
	}

	return false
}

// SetReturn gets a reference to the given GETCapturesCaptureId200ResponseDataRelationshipsReturn and assigns it to the Return field.
func (o *POSTReturnLineItems201ResponseDataRelationships) SetReturn(v GETCapturesCaptureId200ResponseDataRelationshipsReturn) {
	o.Return = &v
}

// GetLineItem returns the LineItem field value if set, zero value otherwise.
func (o *POSTReturnLineItems201ResponseDataRelationships) GetLineItem() POSTLineItemOptions201ResponseDataRelationshipsLineItem {
	if o == nil || IsNil(o.LineItem) {
		var ret POSTLineItemOptions201ResponseDataRelationshipsLineItem
		return ret
	}
	return *o.LineItem
}

// GetLineItemOk returns a tuple with the LineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturnLineItems201ResponseDataRelationships) GetLineItemOk() (*POSTLineItemOptions201ResponseDataRelationshipsLineItem, bool) {
	if o == nil || IsNil(o.LineItem) {
		return nil, false
	}
	return o.LineItem, true
}

// HasLineItem returns a boolean if a field has been set.
func (o *POSTReturnLineItems201ResponseDataRelationships) HasLineItem() bool {
	if o != nil && !IsNil(o.LineItem) {
		return true
	}

	return false
}

// SetLineItem gets a reference to the given POSTLineItemOptions201ResponseDataRelationshipsLineItem and assigns it to the LineItem field.
func (o *POSTReturnLineItems201ResponseDataRelationships) SetLineItem(v POSTLineItemOptions201ResponseDataRelationshipsLineItem) {
	o.LineItem = &v
}

func (o POSTReturnLineItems201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTReturnLineItems201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Return) {
		toSerialize["return"] = o.Return
	}
	if !IsNil(o.LineItem) {
		toSerialize["line_item"] = o.LineItem
	}
	return toSerialize, nil
}

type NullablePOSTReturnLineItems201ResponseDataRelationships struct {
	value *POSTReturnLineItems201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTReturnLineItems201ResponseDataRelationships) Get() *POSTReturnLineItems201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTReturnLineItems201ResponseDataRelationships) Set(val *POSTReturnLineItems201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTReturnLineItems201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTReturnLineItems201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTReturnLineItems201ResponseDataRelationships(val *POSTReturnLineItems201ResponseDataRelationships) *NullablePOSTReturnLineItems201ResponseDataRelationships {
	return &NullablePOSTReturnLineItems201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTReturnLineItems201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTReturnLineItems201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
