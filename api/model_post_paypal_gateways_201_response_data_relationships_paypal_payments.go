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

// checks if the POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments{}

// POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments struct for POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments
type POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks           `json:"links,omitempty"`
	Data  *POSTPaypalGateways201ResponseDataRelationshipsPaypalPaymentsData `json:"data,omitempty"`
}

// NewPOSTPaypalGateways201ResponseDataRelationshipsPaypalPayments instantiates a new POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPaypalGateways201ResponseDataRelationshipsPaypalPayments() *POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments {
	this := POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments{}
	return &this
}

// NewPOSTPaypalGateways201ResponseDataRelationshipsPaypalPaymentsWithDefaults instantiates a new POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPaypalGateways201ResponseDataRelationshipsPaypalPaymentsWithDefaults() *POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments {
	this := POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) GetData() POSTPaypalGateways201ResponseDataRelationshipsPaypalPaymentsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTPaypalGateways201ResponseDataRelationshipsPaypalPaymentsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) GetDataOk() (*POSTPaypalGateways201ResponseDataRelationshipsPaypalPaymentsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTPaypalGateways201ResponseDataRelationshipsPaypalPaymentsData and assigns it to the Data field.
func (o *POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) SetData(v POSTPaypalGateways201ResponseDataRelationshipsPaypalPaymentsData) {
	o.Data = &v
}

func (o POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTPaypalGateways201ResponseDataRelationshipsPaypalPayments struct {
	value *POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments
	isSet bool
}

func (v NullablePOSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) Get() *POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments {
	return v.value
}

func (v *NullablePOSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) Set(val *POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPaypalGateways201ResponseDataRelationshipsPaypalPayments(val *POSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) *NullablePOSTPaypalGateways201ResponseDataRelationshipsPaypalPayments {
	return &NullablePOSTPaypalGateways201ResponseDataRelationshipsPaypalPayments{value: val, isSet: true}
}

func (v NullablePOSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPaypalGateways201ResponseDataRelationshipsPaypalPayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
