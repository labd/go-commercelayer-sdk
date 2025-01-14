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

// checks if the LinkCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkCreateDataRelationships{}

// LinkCreateDataRelationships struct for LinkCreateDataRelationships
type LinkCreateDataRelationships struct {
	Item LinkCreateDataRelationshipsItem `json:"item"`
}

// NewLinkCreateDataRelationships instantiates a new LinkCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkCreateDataRelationships(item LinkCreateDataRelationshipsItem) *LinkCreateDataRelationships {
	this := LinkCreateDataRelationships{}
	this.Item = item
	return &this
}

// NewLinkCreateDataRelationshipsWithDefaults instantiates a new LinkCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkCreateDataRelationshipsWithDefaults() *LinkCreateDataRelationships {
	this := LinkCreateDataRelationships{}
	return &this
}

// GetItem returns the Item field value
func (o *LinkCreateDataRelationships) GetItem() LinkCreateDataRelationshipsItem {
	if o == nil {
		var ret LinkCreateDataRelationshipsItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *LinkCreateDataRelationships) GetItemOk() (*LinkCreateDataRelationshipsItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *LinkCreateDataRelationships) SetItem(v LinkCreateDataRelationshipsItem) {
	o.Item = v
}

func (o LinkCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["item"] = o.Item
	return toSerialize, nil
}

type NullableLinkCreateDataRelationships struct {
	value *LinkCreateDataRelationships
	isSet bool
}

func (v NullableLinkCreateDataRelationships) Get() *LinkCreateDataRelationships {
	return v.value
}

func (v *NullableLinkCreateDataRelationships) Set(val *LinkCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkCreateDataRelationships(val *LinkCreateDataRelationships) *NullableLinkCreateDataRelationships {
	return &NullableLinkCreateDataRelationships{value: val, isSet: true}
}

func (v NullableLinkCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
