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

// checks if the POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems{}

// POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems struct for POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems
type POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks                       `json:"links,omitempty"`
	Data  *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData `json:"data,omitempty"`
}

// NewPOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems instantiates a new POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems() *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems {
	this := POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems{}
	return &this
}

// NewPOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsWithDefaults instantiates a new POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsWithDefaults() *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems {
	this := POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) GetData() POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) GetDataOk() (*POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData and assigns it to the Data field.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) SetData(v POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) {
	o.Data = &v
}

func (o POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems struct {
	value *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems
	isSet bool
}

func (v NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) Get() *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems {
	return v.value
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) Set(val *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems(val *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems {
	return &NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems{value: val, isSet: true}
}

func (v NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
