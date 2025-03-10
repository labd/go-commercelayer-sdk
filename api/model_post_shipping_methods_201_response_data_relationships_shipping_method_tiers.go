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

// checks if the POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers{}

// POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers struct for POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers
type POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks                 `json:"links,omitempty"`
	Data  *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData `json:"data,omitempty"`
}

// NewPOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers instantiates a new POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers() *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers {
	this := POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers{}
	return &this
}

// NewPOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersWithDefaults instantiates a new POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersWithDefaults() *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers {
	this := POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) GetData() POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData {
	if o == nil || IsNil(o.Data) {
		var ret POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) GetDataOk() (*POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData and assigns it to the Data field.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) SetData(v POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) {
	o.Data = &v
}

func (o POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers struct {
	value *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers
	isSet bool
}

func (v NullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) Get() *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers {
	return v.value
}

func (v *NullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) Set(val *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers(val *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) *NullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers {
	return &NullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers{value: val, isSet: true}
}

func (v NullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
