/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LineItemResponseDataRelationshipsItem struct for LineItemResponseDataRelationshipsItem
type LineItemResponseDataRelationshipsItem struct {
	Links *AddressResponseDataRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *LineItemResponseDataRelationshipsItemData     `json:"data,omitempty"`
}

// NewLineItemResponseDataRelationshipsItem instantiates a new LineItemResponseDataRelationshipsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemResponseDataRelationshipsItem() *LineItemResponseDataRelationshipsItem {
	this := LineItemResponseDataRelationshipsItem{}
	return &this
}

// NewLineItemResponseDataRelationshipsItemWithDefaults instantiates a new LineItemResponseDataRelationshipsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemResponseDataRelationshipsItemWithDefaults() *LineItemResponseDataRelationshipsItem {
	this := LineItemResponseDataRelationshipsItem{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *LineItemResponseDataRelationshipsItem) GetLinks() AddressResponseDataRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemResponseDataRelationshipsItem) GetLinksOk() (*AddressResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *LineItemResponseDataRelationshipsItem) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *LineItemResponseDataRelationshipsItem) SetLinks(v AddressResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LineItemResponseDataRelationshipsItem) GetData() LineItemResponseDataRelationshipsItemData {
	if o == nil || o.Data == nil {
		var ret LineItemResponseDataRelationshipsItemData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemResponseDataRelationshipsItem) GetDataOk() (*LineItemResponseDataRelationshipsItemData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LineItemResponseDataRelationshipsItem) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given LineItemResponseDataRelationshipsItemData and assigns it to the Data field.
func (o *LineItemResponseDataRelationshipsItem) SetData(v LineItemResponseDataRelationshipsItemData) {
	o.Data = &v
}

func (o LineItemResponseDataRelationshipsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableLineItemResponseDataRelationshipsItem struct {
	value *LineItemResponseDataRelationshipsItem
	isSet bool
}

func (v NullableLineItemResponseDataRelationshipsItem) Get() *LineItemResponseDataRelationshipsItem {
	return v.value
}

func (v *NullableLineItemResponseDataRelationshipsItem) Set(val *LineItemResponseDataRelationshipsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemResponseDataRelationshipsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemResponseDataRelationshipsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemResponseDataRelationshipsItem(val *LineItemResponseDataRelationshipsItem) *NullableLineItemResponseDataRelationshipsItem {
	return &NullableLineItemResponseDataRelationshipsItem{value: val, isSet: true}
}

func (v NullableLineItemResponseDataRelationshipsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemResponseDataRelationshipsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
