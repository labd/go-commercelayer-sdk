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

// checks if the LinkDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkDataRelationships{}

// LinkDataRelationships struct for LinkDataRelationships
type LinkDataRelationships struct {
	Item   *LinkDataRelationshipsItem      `json:"item,omitempty"`
	Events *AddressDataRelationshipsEvents `json:"events,omitempty"`
}

// NewLinkDataRelationships instantiates a new LinkDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkDataRelationships() *LinkDataRelationships {
	this := LinkDataRelationships{}
	return &this
}

// NewLinkDataRelationshipsWithDefaults instantiates a new LinkDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkDataRelationshipsWithDefaults() *LinkDataRelationships {
	this := LinkDataRelationships{}
	return &this
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *LinkDataRelationships) GetItem() LinkDataRelationshipsItem {
	if o == nil || IsNil(o.Item) {
		var ret LinkDataRelationshipsItem
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDataRelationships) GetItemOk() (*LinkDataRelationshipsItem, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *LinkDataRelationships) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given LinkDataRelationshipsItem and assigns it to the Item field.
func (o *LinkDataRelationships) SetItem(v LinkDataRelationshipsItem) {
	o.Item = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *LinkDataRelationships) GetEvents() AddressDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret AddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *LinkDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given AddressDataRelationshipsEvents and assigns it to the Events field.
func (o *LinkDataRelationships) SetEvents(v AddressDataRelationshipsEvents) {
	o.Events = &v
}

func (o LinkDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return toSerialize, nil
}

type NullableLinkDataRelationships struct {
	value *LinkDataRelationships
	isSet bool
}

func (v NullableLinkDataRelationships) Get() *LinkDataRelationships {
	return v.value
}

func (v *NullableLinkDataRelationships) Set(val *LinkDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDataRelationships(val *LinkDataRelationships) *NullableLinkDataRelationships {
	return &NullableLinkDataRelationships{value: val, isSet: true}
}

func (v NullableLinkDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
