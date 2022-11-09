/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETAttachments200ResponseDataInnerRelationshipsAttachable struct for GETAttachments200ResponseDataInnerRelationshipsAttachable
type GETAttachments200ResponseDataInnerRelationshipsAttachable struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks    `json:"links,omitempty"`
	Data  *GETAttachments200ResponseDataInnerRelationshipsAttachableData `json:"data,omitempty"`
}

// NewGETAttachments200ResponseDataInnerRelationshipsAttachable instantiates a new GETAttachments200ResponseDataInnerRelationshipsAttachable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAttachments200ResponseDataInnerRelationshipsAttachable() *GETAttachments200ResponseDataInnerRelationshipsAttachable {
	this := GETAttachments200ResponseDataInnerRelationshipsAttachable{}
	return &this
}

// NewGETAttachments200ResponseDataInnerRelationshipsAttachableWithDefaults instantiates a new GETAttachments200ResponseDataInnerRelationshipsAttachable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAttachments200ResponseDataInnerRelationshipsAttachableWithDefaults() *GETAttachments200ResponseDataInnerRelationshipsAttachable {
	this := GETAttachments200ResponseDataInnerRelationshipsAttachable{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETAttachments200ResponseDataInnerRelationshipsAttachable) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAttachments200ResponseDataInnerRelationshipsAttachable) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETAttachments200ResponseDataInnerRelationshipsAttachable) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETAttachments200ResponseDataInnerRelationshipsAttachable) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETAttachments200ResponseDataInnerRelationshipsAttachable) GetData() GETAttachments200ResponseDataInnerRelationshipsAttachableData {
	if o == nil || o.Data == nil {
		var ret GETAttachments200ResponseDataInnerRelationshipsAttachableData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAttachments200ResponseDataInnerRelationshipsAttachable) GetDataOk() (*GETAttachments200ResponseDataInnerRelationshipsAttachableData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETAttachments200ResponseDataInnerRelationshipsAttachable) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETAttachments200ResponseDataInnerRelationshipsAttachableData and assigns it to the Data field.
func (o *GETAttachments200ResponseDataInnerRelationshipsAttachable) SetData(v GETAttachments200ResponseDataInnerRelationshipsAttachableData) {
	o.Data = &v
}

func (o GETAttachments200ResponseDataInnerRelationshipsAttachable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETAttachments200ResponseDataInnerRelationshipsAttachable struct {
	value *GETAttachments200ResponseDataInnerRelationshipsAttachable
	isSet bool
}

func (v NullableGETAttachments200ResponseDataInnerRelationshipsAttachable) Get() *GETAttachments200ResponseDataInnerRelationshipsAttachable {
	return v.value
}

func (v *NullableGETAttachments200ResponseDataInnerRelationshipsAttachable) Set(val *GETAttachments200ResponseDataInnerRelationshipsAttachable) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAttachments200ResponseDataInnerRelationshipsAttachable) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAttachments200ResponseDataInnerRelationshipsAttachable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAttachments200ResponseDataInnerRelationshipsAttachable(val *GETAttachments200ResponseDataInnerRelationshipsAttachable) *NullableGETAttachments200ResponseDataInnerRelationshipsAttachable {
	return &NullableGETAttachments200ResponseDataInnerRelationshipsAttachable{value: val, isSet: true}
}

func (v NullableGETAttachments200ResponseDataInnerRelationshipsAttachable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAttachments200ResponseDataInnerRelationshipsAttachable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
