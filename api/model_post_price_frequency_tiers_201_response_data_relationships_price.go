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

// checks if the POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice{}

// POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice struct for POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice
type POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks       `json:"links,omitempty"`
	Data  *POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData `json:"data,omitempty"`
}

// NewPOSTPriceFrequencyTiers201ResponseDataRelationshipsPrice instantiates a new POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPriceFrequencyTiers201ResponseDataRelationshipsPrice() *POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice {
	this := POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice{}
	return &this
}

// NewPOSTPriceFrequencyTiers201ResponseDataRelationshipsPriceWithDefaults instantiates a new POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPriceFrequencyTiers201ResponseDataRelationshipsPriceWithDefaults() *POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice {
	this := POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) GetData() POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData {
	if o == nil || IsNil(o.Data) {
		var ret POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) GetDataOk() (*POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData and assigns it to the Data field.
func (o *POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) SetData(v POSTPriceFrequencyTiers201ResponseDataRelationshipsPriceData) {
	o.Data = &v
}

func (o POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPrice struct {
	value *POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice
	isSet bool
}

func (v NullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) Get() *POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice {
	return v.value
}

func (v *NullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) Set(val *POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPrice(val *POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) *NullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPrice {
	return &NullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPrice{value: val, isSet: true}
}

func (v NullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
