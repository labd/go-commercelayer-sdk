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

// checks if the POSTLineItems201ResponseDataRelationshipsGiftCard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTLineItems201ResponseDataRelationshipsGiftCard{}

// POSTLineItems201ResponseDataRelationshipsGiftCard struct for POSTLineItems201ResponseDataRelationshipsGiftCard
type POSTLineItems201ResponseDataRelationshipsGiftCard struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *POSTLineItems201ResponseDataRelationshipsGiftCardData  `json:"data,omitempty"`
}

// NewPOSTLineItems201ResponseDataRelationshipsGiftCard instantiates a new POSTLineItems201ResponseDataRelationshipsGiftCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTLineItems201ResponseDataRelationshipsGiftCard() *POSTLineItems201ResponseDataRelationshipsGiftCard {
	this := POSTLineItems201ResponseDataRelationshipsGiftCard{}
	return &this
}

// NewPOSTLineItems201ResponseDataRelationshipsGiftCardWithDefaults instantiates a new POSTLineItems201ResponseDataRelationshipsGiftCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTLineItems201ResponseDataRelationshipsGiftCardWithDefaults() *POSTLineItems201ResponseDataRelationshipsGiftCard {
	this := POSTLineItems201ResponseDataRelationshipsGiftCard{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationshipsGiftCard) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationshipsGiftCard) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationshipsGiftCard) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTLineItems201ResponseDataRelationshipsGiftCard) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationshipsGiftCard) GetData() POSTLineItems201ResponseDataRelationshipsGiftCardData {
	if o == nil || IsNil(o.Data) {
		var ret POSTLineItems201ResponseDataRelationshipsGiftCardData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationshipsGiftCard) GetDataOk() (*POSTLineItems201ResponseDataRelationshipsGiftCardData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationshipsGiftCard) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTLineItems201ResponseDataRelationshipsGiftCardData and assigns it to the Data field.
func (o *POSTLineItems201ResponseDataRelationshipsGiftCard) SetData(v POSTLineItems201ResponseDataRelationshipsGiftCardData) {
	o.Data = &v
}

func (o POSTLineItems201ResponseDataRelationshipsGiftCard) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTLineItems201ResponseDataRelationshipsGiftCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTLineItems201ResponseDataRelationshipsGiftCard struct {
	value *POSTLineItems201ResponseDataRelationshipsGiftCard
	isSet bool
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsGiftCard) Get() *POSTLineItems201ResponseDataRelationshipsGiftCard {
	return v.value
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsGiftCard) Set(val *POSTLineItems201ResponseDataRelationshipsGiftCard) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsGiftCard) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsGiftCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTLineItems201ResponseDataRelationshipsGiftCard(val *POSTLineItems201ResponseDataRelationshipsGiftCard) *NullablePOSTLineItems201ResponseDataRelationshipsGiftCard {
	return &NullablePOSTLineItems201ResponseDataRelationshipsGiftCard{value: val, isSet: true}
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsGiftCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsGiftCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
