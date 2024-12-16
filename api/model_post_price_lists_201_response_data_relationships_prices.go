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

// checks if the POSTPriceLists201ResponseDataRelationshipsPrices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPriceLists201ResponseDataRelationshipsPrices{}

// POSTPriceLists201ResponseDataRelationshipsPrices struct for POSTPriceLists201ResponseDataRelationshipsPrices
type POSTPriceLists201ResponseDataRelationshipsPrices struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *POSTPriceLists201ResponseDataRelationshipsPricesData   `json:"data,omitempty"`
}

// NewPOSTPriceLists201ResponseDataRelationshipsPrices instantiates a new POSTPriceLists201ResponseDataRelationshipsPrices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPriceLists201ResponseDataRelationshipsPrices() *POSTPriceLists201ResponseDataRelationshipsPrices {
	this := POSTPriceLists201ResponseDataRelationshipsPrices{}
	return &this
}

// NewPOSTPriceLists201ResponseDataRelationshipsPricesWithDefaults instantiates a new POSTPriceLists201ResponseDataRelationshipsPrices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPriceLists201ResponseDataRelationshipsPricesWithDefaults() *POSTPriceLists201ResponseDataRelationshipsPrices {
	this := POSTPriceLists201ResponseDataRelationshipsPrices{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTPriceLists201ResponseDataRelationshipsPrices) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPriceLists201ResponseDataRelationshipsPrices) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTPriceLists201ResponseDataRelationshipsPrices) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTPriceLists201ResponseDataRelationshipsPrices) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTPriceLists201ResponseDataRelationshipsPrices) GetData() POSTPriceLists201ResponseDataRelationshipsPricesData {
	if o == nil || IsNil(o.Data) {
		var ret POSTPriceLists201ResponseDataRelationshipsPricesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPriceLists201ResponseDataRelationshipsPrices) GetDataOk() (*POSTPriceLists201ResponseDataRelationshipsPricesData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTPriceLists201ResponseDataRelationshipsPrices) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTPriceLists201ResponseDataRelationshipsPricesData and assigns it to the Data field.
func (o *POSTPriceLists201ResponseDataRelationshipsPrices) SetData(v POSTPriceLists201ResponseDataRelationshipsPricesData) {
	o.Data = &v
}

func (o POSTPriceLists201ResponseDataRelationshipsPrices) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPriceLists201ResponseDataRelationshipsPrices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTPriceLists201ResponseDataRelationshipsPrices struct {
	value *POSTPriceLists201ResponseDataRelationshipsPrices
	isSet bool
}

func (v NullablePOSTPriceLists201ResponseDataRelationshipsPrices) Get() *POSTPriceLists201ResponseDataRelationshipsPrices {
	return v.value
}

func (v *NullablePOSTPriceLists201ResponseDataRelationshipsPrices) Set(val *POSTPriceLists201ResponseDataRelationshipsPrices) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPriceLists201ResponseDataRelationshipsPrices) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPriceLists201ResponseDataRelationshipsPrices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPriceLists201ResponseDataRelationshipsPrices(val *POSTPriceLists201ResponseDataRelationshipsPrices) *NullablePOSTPriceLists201ResponseDataRelationshipsPrices {
	return &NullablePOSTPriceLists201ResponseDataRelationshipsPrices{value: val, isSet: true}
}

func (v NullablePOSTPriceLists201ResponseDataRelationshipsPrices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPriceLists201ResponseDataRelationshipsPrices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
