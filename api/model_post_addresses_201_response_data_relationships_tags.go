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

// checks if the POSTAddresses201ResponseDataRelationshipsTags type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAddresses201ResponseDataRelationshipsTags{}

// POSTAddresses201ResponseDataRelationshipsTags struct for POSTAddresses201ResponseDataRelationshipsTags
type POSTAddresses201ResponseDataRelationshipsTags struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *POSTAddresses201ResponseDataRelationshipsTagsData      `json:"data,omitempty"`
}

// NewPOSTAddresses201ResponseDataRelationshipsTags instantiates a new POSTAddresses201ResponseDataRelationshipsTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAddresses201ResponseDataRelationshipsTags() *POSTAddresses201ResponseDataRelationshipsTags {
	this := POSTAddresses201ResponseDataRelationshipsTags{}
	return &this
}

// NewPOSTAddresses201ResponseDataRelationshipsTagsWithDefaults instantiates a new POSTAddresses201ResponseDataRelationshipsTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAddresses201ResponseDataRelationshipsTagsWithDefaults() *POSTAddresses201ResponseDataRelationshipsTags {
	this := POSTAddresses201ResponseDataRelationshipsTags{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTAddresses201ResponseDataRelationshipsTags) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAddresses201ResponseDataRelationshipsTags) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataRelationshipsTags) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTAddresses201ResponseDataRelationshipsTags) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTAddresses201ResponseDataRelationshipsTags) GetData() POSTAddresses201ResponseDataRelationshipsTagsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTAddresses201ResponseDataRelationshipsTagsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAddresses201ResponseDataRelationshipsTags) GetDataOk() (*POSTAddresses201ResponseDataRelationshipsTagsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataRelationshipsTags) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTAddresses201ResponseDataRelationshipsTagsData and assigns it to the Data field.
func (o *POSTAddresses201ResponseDataRelationshipsTags) SetData(v POSTAddresses201ResponseDataRelationshipsTagsData) {
	o.Data = &v
}

func (o POSTAddresses201ResponseDataRelationshipsTags) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAddresses201ResponseDataRelationshipsTags) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTAddresses201ResponseDataRelationshipsTags struct {
	value *POSTAddresses201ResponseDataRelationshipsTags
	isSet bool
}

func (v NullablePOSTAddresses201ResponseDataRelationshipsTags) Get() *POSTAddresses201ResponseDataRelationshipsTags {
	return v.value
}

func (v *NullablePOSTAddresses201ResponseDataRelationshipsTags) Set(val *POSTAddresses201ResponseDataRelationshipsTags) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAddresses201ResponseDataRelationshipsTags) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAddresses201ResponseDataRelationshipsTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAddresses201ResponseDataRelationshipsTags(val *POSTAddresses201ResponseDataRelationshipsTags) *NullablePOSTAddresses201ResponseDataRelationshipsTags {
	return &NullablePOSTAddresses201ResponseDataRelationshipsTags{value: val, isSet: true}
}

func (v NullablePOSTAddresses201ResponseDataRelationshipsTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAddresses201ResponseDataRelationshipsTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
