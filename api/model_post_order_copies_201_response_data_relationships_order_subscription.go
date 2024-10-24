/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTOrderCopies201ResponseDataRelationshipsOrderSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrderCopies201ResponseDataRelationshipsOrderSubscription{}

// POSTOrderCopies201ResponseDataRelationshipsOrderSubscription struct for POSTOrderCopies201ResponseDataRelationshipsOrderSubscription
type POSTOrderCopies201ResponseDataRelationshipsOrderSubscription struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks           `json:"links,omitempty"`
	Data  *POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData `json:"data,omitempty"`
}

// NewPOSTOrderCopies201ResponseDataRelationshipsOrderSubscription instantiates a new POSTOrderCopies201ResponseDataRelationshipsOrderSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderCopies201ResponseDataRelationshipsOrderSubscription() *POSTOrderCopies201ResponseDataRelationshipsOrderSubscription {
	this := POSTOrderCopies201ResponseDataRelationshipsOrderSubscription{}
	return &this
}

// NewPOSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionWithDefaults instantiates a new POSTOrderCopies201ResponseDataRelationshipsOrderSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionWithDefaults() *POSTOrderCopies201ResponseDataRelationshipsOrderSubscription {
	this := POSTOrderCopies201ResponseDataRelationshipsOrderSubscription{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTOrderCopies201ResponseDataRelationshipsOrderSubscription) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderCopies201ResponseDataRelationshipsOrderSubscription) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataRelationshipsOrderSubscription) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTOrderCopies201ResponseDataRelationshipsOrderSubscription) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTOrderCopies201ResponseDataRelationshipsOrderSubscription) GetData() POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData {
	if o == nil || IsNil(o.Data) {
		var ret POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderCopies201ResponseDataRelationshipsOrderSubscription) GetDataOk() (*POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataRelationshipsOrderSubscription) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData and assigns it to the Data field.
func (o *POSTOrderCopies201ResponseDataRelationshipsOrderSubscription) SetData(v POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) {
	o.Data = &v
}

func (o POSTOrderCopies201ResponseDataRelationshipsOrderSubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrderCopies201ResponseDataRelationshipsOrderSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscription struct {
	value *POSTOrderCopies201ResponseDataRelationshipsOrderSubscription
	isSet bool
}

func (v NullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscription) Get() *POSTOrderCopies201ResponseDataRelationshipsOrderSubscription {
	return v.value
}

func (v *NullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscription) Set(val *POSTOrderCopies201ResponseDataRelationshipsOrderSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscription(val *POSTOrderCopies201ResponseDataRelationshipsOrderSubscription) *NullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscription {
	return &NullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscription{value: val, isSet: true}
}

func (v NullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}