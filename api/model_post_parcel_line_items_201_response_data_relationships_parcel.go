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

// checks if the POSTParcelLineItems201ResponseDataRelationshipsParcel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTParcelLineItems201ResponseDataRelationshipsParcel{}

// POSTParcelLineItems201ResponseDataRelationshipsParcel struct for POSTParcelLineItems201ResponseDataRelationshipsParcel
type POSTParcelLineItems201ResponseDataRelationshipsParcel struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks    `json:"links,omitempty"`
	Data  *POSTParcelLineItems201ResponseDataRelationshipsParcelData `json:"data,omitempty"`
}

// NewPOSTParcelLineItems201ResponseDataRelationshipsParcel instantiates a new POSTParcelLineItems201ResponseDataRelationshipsParcel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTParcelLineItems201ResponseDataRelationshipsParcel() *POSTParcelLineItems201ResponseDataRelationshipsParcel {
	this := POSTParcelLineItems201ResponseDataRelationshipsParcel{}
	return &this
}

// NewPOSTParcelLineItems201ResponseDataRelationshipsParcelWithDefaults instantiates a new POSTParcelLineItems201ResponseDataRelationshipsParcel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTParcelLineItems201ResponseDataRelationshipsParcelWithDefaults() *POSTParcelLineItems201ResponseDataRelationshipsParcel {
	this := POSTParcelLineItems201ResponseDataRelationshipsParcel{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTParcelLineItems201ResponseDataRelationshipsParcel) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTParcelLineItems201ResponseDataRelationshipsParcel) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTParcelLineItems201ResponseDataRelationshipsParcel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTParcelLineItems201ResponseDataRelationshipsParcel) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTParcelLineItems201ResponseDataRelationshipsParcel) GetData() POSTParcelLineItems201ResponseDataRelationshipsParcelData {
	if o == nil || IsNil(o.Data) {
		var ret POSTParcelLineItems201ResponseDataRelationshipsParcelData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTParcelLineItems201ResponseDataRelationshipsParcel) GetDataOk() (*POSTParcelLineItems201ResponseDataRelationshipsParcelData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTParcelLineItems201ResponseDataRelationshipsParcel) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTParcelLineItems201ResponseDataRelationshipsParcelData and assigns it to the Data field.
func (o *POSTParcelLineItems201ResponseDataRelationshipsParcel) SetData(v POSTParcelLineItems201ResponseDataRelationshipsParcelData) {
	o.Data = &v
}

func (o POSTParcelLineItems201ResponseDataRelationshipsParcel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTParcelLineItems201ResponseDataRelationshipsParcel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTParcelLineItems201ResponseDataRelationshipsParcel struct {
	value *POSTParcelLineItems201ResponseDataRelationshipsParcel
	isSet bool
}

func (v NullablePOSTParcelLineItems201ResponseDataRelationshipsParcel) Get() *POSTParcelLineItems201ResponseDataRelationshipsParcel {
	return v.value
}

func (v *NullablePOSTParcelLineItems201ResponseDataRelationshipsParcel) Set(val *POSTParcelLineItems201ResponseDataRelationshipsParcel) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTParcelLineItems201ResponseDataRelationshipsParcel) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTParcelLineItems201ResponseDataRelationshipsParcel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTParcelLineItems201ResponseDataRelationshipsParcel(val *POSTParcelLineItems201ResponseDataRelationshipsParcel) *NullablePOSTParcelLineItems201ResponseDataRelationshipsParcel {
	return &NullablePOSTParcelLineItems201ResponseDataRelationshipsParcel{value: val, isSet: true}
}

func (v NullablePOSTParcelLineItems201ResponseDataRelationshipsParcel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTParcelLineItems201ResponseDataRelationshipsParcel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
