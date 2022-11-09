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

// GETParcels200ResponseDataInnerRelationshipsPackage struct for GETParcels200ResponseDataInnerRelationshipsPackage
type GETParcels200ResponseDataInnerRelationshipsPackage struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *GETParcels200ResponseDataInnerRelationshipsPackageData     `json:"data,omitempty"`
}

// NewGETParcels200ResponseDataInnerRelationshipsPackage instantiates a new GETParcels200ResponseDataInnerRelationshipsPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETParcels200ResponseDataInnerRelationshipsPackage() *GETParcels200ResponseDataInnerRelationshipsPackage {
	this := GETParcels200ResponseDataInnerRelationshipsPackage{}
	return &this
}

// NewGETParcels200ResponseDataInnerRelationshipsPackageWithDefaults instantiates a new GETParcels200ResponseDataInnerRelationshipsPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETParcels200ResponseDataInnerRelationshipsPackageWithDefaults() *GETParcels200ResponseDataInnerRelationshipsPackage {
	this := GETParcels200ResponseDataInnerRelationshipsPackage{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETParcels200ResponseDataInnerRelationshipsPackage) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcels200ResponseDataInnerRelationshipsPackage) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETParcels200ResponseDataInnerRelationshipsPackage) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETParcels200ResponseDataInnerRelationshipsPackage) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETParcels200ResponseDataInnerRelationshipsPackage) GetData() GETParcels200ResponseDataInnerRelationshipsPackageData {
	if o == nil || o.Data == nil {
		var ret GETParcels200ResponseDataInnerRelationshipsPackageData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcels200ResponseDataInnerRelationshipsPackage) GetDataOk() (*GETParcels200ResponseDataInnerRelationshipsPackageData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETParcels200ResponseDataInnerRelationshipsPackage) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETParcels200ResponseDataInnerRelationshipsPackageData and assigns it to the Data field.
func (o *GETParcels200ResponseDataInnerRelationshipsPackage) SetData(v GETParcels200ResponseDataInnerRelationshipsPackageData) {
	o.Data = &v
}

func (o GETParcels200ResponseDataInnerRelationshipsPackage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETParcels200ResponseDataInnerRelationshipsPackage struct {
	value *GETParcels200ResponseDataInnerRelationshipsPackage
	isSet bool
}

func (v NullableGETParcels200ResponseDataInnerRelationshipsPackage) Get() *GETParcels200ResponseDataInnerRelationshipsPackage {
	return v.value
}

func (v *NullableGETParcels200ResponseDataInnerRelationshipsPackage) Set(val *GETParcels200ResponseDataInnerRelationshipsPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableGETParcels200ResponseDataInnerRelationshipsPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableGETParcels200ResponseDataInnerRelationshipsPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETParcels200ResponseDataInnerRelationshipsPackage(val *GETParcels200ResponseDataInnerRelationshipsPackage) *NullableGETParcels200ResponseDataInnerRelationshipsPackage {
	return &NullableGETParcels200ResponseDataInnerRelationshipsPackage{value: val, isSet: true}
}

func (v NullableGETParcels200ResponseDataInnerRelationshipsPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETParcels200ResponseDataInnerRelationshipsPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
