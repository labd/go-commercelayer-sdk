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

// checks if the POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments{}

// POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments struct for POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments
type POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks                     `json:"links,omitempty"`
	Data  *POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPaymentsData `json:"data,omitempty"`
}

// NewPOSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments instantiates a new POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments() *POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments {
	this := POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments{}
	return &this
}

// NewPOSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPaymentsWithDefaults instantiates a new POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPaymentsWithDefaults() *POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments {
	this := POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments) GetData() POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPaymentsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPaymentsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments) GetDataOk() (*POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPaymentsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPaymentsData and assigns it to the Data field.
func (o *POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments) SetData(v POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPaymentsData) {
	o.Data = &v
}

func (o POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments struct {
	value *POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments
	isSet bool
}

func (v NullablePOSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments) Get() *POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments {
	return v.value
}

func (v *NullablePOSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments) Set(val *POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments(val *POSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments) *NullablePOSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments {
	return &NullablePOSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments{value: val, isSet: true}
}

func (v NullablePOSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCheckoutComGateways201ResponseDataRelationshipsCheckoutComPayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
