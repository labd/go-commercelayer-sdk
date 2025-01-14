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

// checks if the POSTParcels201ResponseDataRelationshipsPackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTParcels201ResponseDataRelationshipsPackage{}

// POSTParcels201ResponseDataRelationshipsPackage struct for POSTParcels201ResponseDataRelationshipsPackage
type POSTParcels201ResponseDataRelationshipsPackage struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *POSTParcels201ResponseDataRelationshipsPackageData     `json:"data,omitempty"`
}

// NewPOSTParcels201ResponseDataRelationshipsPackage instantiates a new POSTParcels201ResponseDataRelationshipsPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTParcels201ResponseDataRelationshipsPackage() *POSTParcels201ResponseDataRelationshipsPackage {
	this := POSTParcels201ResponseDataRelationshipsPackage{}
	return &this
}

// NewPOSTParcels201ResponseDataRelationshipsPackageWithDefaults instantiates a new POSTParcels201ResponseDataRelationshipsPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTParcels201ResponseDataRelationshipsPackageWithDefaults() *POSTParcels201ResponseDataRelationshipsPackage {
	this := POSTParcels201ResponseDataRelationshipsPackage{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTParcels201ResponseDataRelationshipsPackage) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTParcels201ResponseDataRelationshipsPackage) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataRelationshipsPackage) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTParcels201ResponseDataRelationshipsPackage) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTParcels201ResponseDataRelationshipsPackage) GetData() POSTParcels201ResponseDataRelationshipsPackageData {
	if o == nil || IsNil(o.Data) {
		var ret POSTParcels201ResponseDataRelationshipsPackageData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTParcels201ResponseDataRelationshipsPackage) GetDataOk() (*POSTParcels201ResponseDataRelationshipsPackageData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataRelationshipsPackage) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTParcels201ResponseDataRelationshipsPackageData and assigns it to the Data field.
func (o *POSTParcels201ResponseDataRelationshipsPackage) SetData(v POSTParcels201ResponseDataRelationshipsPackageData) {
	o.Data = &v
}

func (o POSTParcels201ResponseDataRelationshipsPackage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTParcels201ResponseDataRelationshipsPackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTParcels201ResponseDataRelationshipsPackage struct {
	value *POSTParcels201ResponseDataRelationshipsPackage
	isSet bool
}

func (v NullablePOSTParcels201ResponseDataRelationshipsPackage) Get() *POSTParcels201ResponseDataRelationshipsPackage {
	return v.value
}

func (v *NullablePOSTParcels201ResponseDataRelationshipsPackage) Set(val *POSTParcels201ResponseDataRelationshipsPackage) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTParcels201ResponseDataRelationshipsPackage) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTParcels201ResponseDataRelationshipsPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTParcels201ResponseDataRelationshipsPackage(val *POSTParcels201ResponseDataRelationshipsPackage) *NullablePOSTParcels201ResponseDataRelationshipsPackage {
	return &NullablePOSTParcels201ResponseDataRelationshipsPackage{value: val, isSet: true}
}

func (v NullablePOSTParcels201ResponseDataRelationshipsPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTParcels201ResponseDataRelationshipsPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
