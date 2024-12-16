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

// checks if the POSTExternalGateways201ResponseDataRelationshipsExternalPayments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTExternalGateways201ResponseDataRelationshipsExternalPayments{}

// POSTExternalGateways201ResponseDataRelationshipsExternalPayments struct for POSTExternalGateways201ResponseDataRelationshipsExternalPayments
type POSTExternalGateways201ResponseDataRelationshipsExternalPayments struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks               `json:"links,omitempty"`
	Data  *POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData `json:"data,omitempty"`
}

// NewPOSTExternalGateways201ResponseDataRelationshipsExternalPayments instantiates a new POSTExternalGateways201ResponseDataRelationshipsExternalPayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExternalGateways201ResponseDataRelationshipsExternalPayments() *POSTExternalGateways201ResponseDataRelationshipsExternalPayments {
	this := POSTExternalGateways201ResponseDataRelationshipsExternalPayments{}
	return &this
}

// NewPOSTExternalGateways201ResponseDataRelationshipsExternalPaymentsWithDefaults instantiates a new POSTExternalGateways201ResponseDataRelationshipsExternalPayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExternalGateways201ResponseDataRelationshipsExternalPaymentsWithDefaults() *POSTExternalGateways201ResponseDataRelationshipsExternalPayments {
	this := POSTExternalGateways201ResponseDataRelationshipsExternalPayments{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTExternalGateways201ResponseDataRelationshipsExternalPayments) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalGateways201ResponseDataRelationshipsExternalPayments) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTExternalGateways201ResponseDataRelationshipsExternalPayments) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTExternalGateways201ResponseDataRelationshipsExternalPayments) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTExternalGateways201ResponseDataRelationshipsExternalPayments) GetData() POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalGateways201ResponseDataRelationshipsExternalPayments) GetDataOk() (*POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTExternalGateways201ResponseDataRelationshipsExternalPayments) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData and assigns it to the Data field.
func (o *POSTExternalGateways201ResponseDataRelationshipsExternalPayments) SetData(v POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) {
	o.Data = &v
}

func (o POSTExternalGateways201ResponseDataRelationshipsExternalPayments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTExternalGateways201ResponseDataRelationshipsExternalPayments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTExternalGateways201ResponseDataRelationshipsExternalPayments struct {
	value *POSTExternalGateways201ResponseDataRelationshipsExternalPayments
	isSet bool
}

func (v NullablePOSTExternalGateways201ResponseDataRelationshipsExternalPayments) Get() *POSTExternalGateways201ResponseDataRelationshipsExternalPayments {
	return v.value
}

func (v *NullablePOSTExternalGateways201ResponseDataRelationshipsExternalPayments) Set(val *POSTExternalGateways201ResponseDataRelationshipsExternalPayments) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExternalGateways201ResponseDataRelationshipsExternalPayments) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExternalGateways201ResponseDataRelationshipsExternalPayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExternalGateways201ResponseDataRelationshipsExternalPayments(val *POSTExternalGateways201ResponseDataRelationshipsExternalPayments) *NullablePOSTExternalGateways201ResponseDataRelationshipsExternalPayments {
	return &NullablePOSTExternalGateways201ResponseDataRelationshipsExternalPayments{value: val, isSet: true}
}

func (v NullablePOSTExternalGateways201ResponseDataRelationshipsExternalPayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExternalGateways201ResponseDataRelationshipsExternalPayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
