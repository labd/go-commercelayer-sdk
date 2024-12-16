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

// checks if the POSTSkuLists201ResponseDataRelationshipsBundles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTSkuLists201ResponseDataRelationshipsBundles{}

// POSTSkuLists201ResponseDataRelationshipsBundles struct for POSTSkuLists201ResponseDataRelationshipsBundles
type POSTSkuLists201ResponseDataRelationshipsBundles struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *POSTSkuLists201ResponseDataRelationshipsBundlesData    `json:"data,omitempty"`
}

// NewPOSTSkuLists201ResponseDataRelationshipsBundles instantiates a new POSTSkuLists201ResponseDataRelationshipsBundles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSkuLists201ResponseDataRelationshipsBundles() *POSTSkuLists201ResponseDataRelationshipsBundles {
	this := POSTSkuLists201ResponseDataRelationshipsBundles{}
	return &this
}

// NewPOSTSkuLists201ResponseDataRelationshipsBundlesWithDefaults instantiates a new POSTSkuLists201ResponseDataRelationshipsBundles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSkuLists201ResponseDataRelationshipsBundlesWithDefaults() *POSTSkuLists201ResponseDataRelationshipsBundles {
	this := POSTSkuLists201ResponseDataRelationshipsBundles{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTSkuLists201ResponseDataRelationshipsBundles) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSkuLists201ResponseDataRelationshipsBundles) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTSkuLists201ResponseDataRelationshipsBundles) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTSkuLists201ResponseDataRelationshipsBundles) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTSkuLists201ResponseDataRelationshipsBundles) GetData() POSTSkuLists201ResponseDataRelationshipsBundlesData {
	if o == nil || IsNil(o.Data) {
		var ret POSTSkuLists201ResponseDataRelationshipsBundlesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSkuLists201ResponseDataRelationshipsBundles) GetDataOk() (*POSTSkuLists201ResponseDataRelationshipsBundlesData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTSkuLists201ResponseDataRelationshipsBundles) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTSkuLists201ResponseDataRelationshipsBundlesData and assigns it to the Data field.
func (o *POSTSkuLists201ResponseDataRelationshipsBundles) SetData(v POSTSkuLists201ResponseDataRelationshipsBundlesData) {
	o.Data = &v
}

func (o POSTSkuLists201ResponseDataRelationshipsBundles) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTSkuLists201ResponseDataRelationshipsBundles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTSkuLists201ResponseDataRelationshipsBundles struct {
	value *POSTSkuLists201ResponseDataRelationshipsBundles
	isSet bool
}

func (v NullablePOSTSkuLists201ResponseDataRelationshipsBundles) Get() *POSTSkuLists201ResponseDataRelationshipsBundles {
	return v.value
}

func (v *NullablePOSTSkuLists201ResponseDataRelationshipsBundles) Set(val *POSTSkuLists201ResponseDataRelationshipsBundles) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSkuLists201ResponseDataRelationshipsBundles) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSkuLists201ResponseDataRelationshipsBundles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSkuLists201ResponseDataRelationshipsBundles(val *POSTSkuLists201ResponseDataRelationshipsBundles) *NullablePOSTSkuLists201ResponseDataRelationshipsBundles {
	return &NullablePOSTSkuLists201ResponseDataRelationshipsBundles{value: val, isSet: true}
}

func (v NullablePOSTSkuLists201ResponseDataRelationshipsBundles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSkuLists201ResponseDataRelationshipsBundles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
