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

// checks if the POSTOrders201ResponseDataRelationshipsResourceErrors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrders201ResponseDataRelationshipsResourceErrors{}

// POSTOrders201ResponseDataRelationshipsResourceErrors struct for POSTOrders201ResponseDataRelationshipsResourceErrors
type POSTOrders201ResponseDataRelationshipsResourceErrors struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks   `json:"links,omitempty"`
	Data  *POSTOrders201ResponseDataRelationshipsResourceErrorsData `json:"data,omitempty"`
}

// NewPOSTOrders201ResponseDataRelationshipsResourceErrors instantiates a new POSTOrders201ResponseDataRelationshipsResourceErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrders201ResponseDataRelationshipsResourceErrors() *POSTOrders201ResponseDataRelationshipsResourceErrors {
	this := POSTOrders201ResponseDataRelationshipsResourceErrors{}
	return &this
}

// NewPOSTOrders201ResponseDataRelationshipsResourceErrorsWithDefaults instantiates a new POSTOrders201ResponseDataRelationshipsResourceErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrders201ResponseDataRelationshipsResourceErrorsWithDefaults() *POSTOrders201ResponseDataRelationshipsResourceErrors {
	this := POSTOrders201ResponseDataRelationshipsResourceErrors{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsResourceErrors) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsResourceErrors) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsResourceErrors) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTOrders201ResponseDataRelationshipsResourceErrors) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsResourceErrors) GetData() POSTOrders201ResponseDataRelationshipsResourceErrorsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTOrders201ResponseDataRelationshipsResourceErrorsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsResourceErrors) GetDataOk() (*POSTOrders201ResponseDataRelationshipsResourceErrorsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsResourceErrors) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTOrders201ResponseDataRelationshipsResourceErrorsData and assigns it to the Data field.
func (o *POSTOrders201ResponseDataRelationshipsResourceErrors) SetData(v POSTOrders201ResponseDataRelationshipsResourceErrorsData) {
	o.Data = &v
}

func (o POSTOrders201ResponseDataRelationshipsResourceErrors) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrders201ResponseDataRelationshipsResourceErrors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTOrders201ResponseDataRelationshipsResourceErrors struct {
	value *POSTOrders201ResponseDataRelationshipsResourceErrors
	isSet bool
}

func (v NullablePOSTOrders201ResponseDataRelationshipsResourceErrors) Get() *POSTOrders201ResponseDataRelationshipsResourceErrors {
	return v.value
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsResourceErrors) Set(val *POSTOrders201ResponseDataRelationshipsResourceErrors) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrders201ResponseDataRelationshipsResourceErrors) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsResourceErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrders201ResponseDataRelationshipsResourceErrors(val *POSTOrders201ResponseDataRelationshipsResourceErrors) *NullablePOSTOrders201ResponseDataRelationshipsResourceErrors {
	return &NullablePOSTOrders201ResponseDataRelationshipsResourceErrors{value: val, isSet: true}
}

func (v NullablePOSTOrders201ResponseDataRelationshipsResourceErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsResourceErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
