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

// checks if the POSTAddresses201ResponseDataRelationshipsVersions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAddresses201ResponseDataRelationshipsVersions{}

// POSTAddresses201ResponseDataRelationshipsVersions struct for POSTAddresses201ResponseDataRelationshipsVersions
type POSTAddresses201ResponseDataRelationshipsVersions struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *POSTAddresses201ResponseDataRelationshipsVersionsData  `json:"data,omitempty"`
}

// NewPOSTAddresses201ResponseDataRelationshipsVersions instantiates a new POSTAddresses201ResponseDataRelationshipsVersions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAddresses201ResponseDataRelationshipsVersions() *POSTAddresses201ResponseDataRelationshipsVersions {
	this := POSTAddresses201ResponseDataRelationshipsVersions{}
	return &this
}

// NewPOSTAddresses201ResponseDataRelationshipsVersionsWithDefaults instantiates a new POSTAddresses201ResponseDataRelationshipsVersions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAddresses201ResponseDataRelationshipsVersionsWithDefaults() *POSTAddresses201ResponseDataRelationshipsVersions {
	this := POSTAddresses201ResponseDataRelationshipsVersions{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTAddresses201ResponseDataRelationshipsVersions) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAddresses201ResponseDataRelationshipsVersions) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataRelationshipsVersions) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTAddresses201ResponseDataRelationshipsVersions) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTAddresses201ResponseDataRelationshipsVersions) GetData() POSTAddresses201ResponseDataRelationshipsVersionsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTAddresses201ResponseDataRelationshipsVersionsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAddresses201ResponseDataRelationshipsVersions) GetDataOk() (*POSTAddresses201ResponseDataRelationshipsVersionsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataRelationshipsVersions) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersionsData and assigns it to the Data field.
func (o *POSTAddresses201ResponseDataRelationshipsVersions) SetData(v POSTAddresses201ResponseDataRelationshipsVersionsData) {
	o.Data = &v
}

func (o POSTAddresses201ResponseDataRelationshipsVersions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAddresses201ResponseDataRelationshipsVersions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTAddresses201ResponseDataRelationshipsVersions struct {
	value *POSTAddresses201ResponseDataRelationshipsVersions
	isSet bool
}

func (v NullablePOSTAddresses201ResponseDataRelationshipsVersions) Get() *POSTAddresses201ResponseDataRelationshipsVersions {
	return v.value
}

func (v *NullablePOSTAddresses201ResponseDataRelationshipsVersions) Set(val *POSTAddresses201ResponseDataRelationshipsVersions) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAddresses201ResponseDataRelationshipsVersions) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAddresses201ResponseDataRelationshipsVersions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAddresses201ResponseDataRelationshipsVersions(val *POSTAddresses201ResponseDataRelationshipsVersions) *NullablePOSTAddresses201ResponseDataRelationshipsVersions {
	return &NullablePOSTAddresses201ResponseDataRelationshipsVersions{value: val, isSet: true}
}

func (v NullablePOSTAddresses201ResponseDataRelationshipsVersions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAddresses201ResponseDataRelationshipsVersions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
