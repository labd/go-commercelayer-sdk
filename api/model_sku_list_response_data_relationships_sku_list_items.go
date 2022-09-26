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

// SkuListResponseDataRelationshipsSkuListItems struct for SkuListResponseDataRelationshipsSkuListItems
type SkuListResponseDataRelationshipsSkuListItems struct {
	Links *AddressResponseDataRelationshipsGeocoderLinks          `json:"links,omitempty"`
	Data  []SkuListResponseDataRelationshipsSkuListItemsDataInner `json:"data,omitempty"`
}

// NewSkuListResponseDataRelationshipsSkuListItems instantiates a new SkuListResponseDataRelationshipsSkuListItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListResponseDataRelationshipsSkuListItems() *SkuListResponseDataRelationshipsSkuListItems {
	this := SkuListResponseDataRelationshipsSkuListItems{}
	return &this
}

// NewSkuListResponseDataRelationshipsSkuListItemsWithDefaults instantiates a new SkuListResponseDataRelationshipsSkuListItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListResponseDataRelationshipsSkuListItemsWithDefaults() *SkuListResponseDataRelationshipsSkuListItems {
	this := SkuListResponseDataRelationshipsSkuListItems{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SkuListResponseDataRelationshipsSkuListItems) GetLinks() AddressResponseDataRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListResponseDataRelationshipsSkuListItems) GetLinksOk() (*AddressResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SkuListResponseDataRelationshipsSkuListItems) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *SkuListResponseDataRelationshipsSkuListItems) SetLinks(v AddressResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SkuListResponseDataRelationshipsSkuListItems) GetData() []SkuListResponseDataRelationshipsSkuListItemsDataInner {
	if o == nil || o.Data == nil {
		var ret []SkuListResponseDataRelationshipsSkuListItemsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListResponseDataRelationshipsSkuListItems) GetDataOk() ([]SkuListResponseDataRelationshipsSkuListItemsDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SkuListResponseDataRelationshipsSkuListItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []SkuListResponseDataRelationshipsSkuListItemsDataInner and assigns it to the Data field.
func (o *SkuListResponseDataRelationshipsSkuListItems) SetData(v []SkuListResponseDataRelationshipsSkuListItemsDataInner) {
	o.Data = v
}

func (o SkuListResponseDataRelationshipsSkuListItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSkuListResponseDataRelationshipsSkuListItems struct {
	value *SkuListResponseDataRelationshipsSkuListItems
	isSet bool
}

func (v NullableSkuListResponseDataRelationshipsSkuListItems) Get() *SkuListResponseDataRelationshipsSkuListItems {
	return v.value
}

func (v *NullableSkuListResponseDataRelationshipsSkuListItems) Set(val *SkuListResponseDataRelationshipsSkuListItems) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListResponseDataRelationshipsSkuListItems) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListResponseDataRelationshipsSkuListItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListResponseDataRelationshipsSkuListItems(val *SkuListResponseDataRelationshipsSkuListItems) *NullableSkuListResponseDataRelationshipsSkuListItems {
	return &NullableSkuListResponseDataRelationshipsSkuListItems{value: val, isSet: true}
}

func (v NullableSkuListResponseDataRelationshipsSkuListItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListResponseDataRelationshipsSkuListItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
