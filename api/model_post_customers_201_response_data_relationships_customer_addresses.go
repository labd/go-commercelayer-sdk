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

// checks if the POSTCustomers201ResponseDataRelationshipsCustomerAddresses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCustomers201ResponseDataRelationshipsCustomerAddresses{}

// POSTCustomers201ResponseDataRelationshipsCustomerAddresses struct for POSTCustomers201ResponseDataRelationshipsCustomerAddresses
type POSTCustomers201ResponseDataRelationshipsCustomerAddresses struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks         `json:"links,omitempty"`
	Data  *POSTCustomers201ResponseDataRelationshipsCustomerAddressesData `json:"data,omitempty"`
}

// NewPOSTCustomers201ResponseDataRelationshipsCustomerAddresses instantiates a new POSTCustomers201ResponseDataRelationshipsCustomerAddresses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomers201ResponseDataRelationshipsCustomerAddresses() *POSTCustomers201ResponseDataRelationshipsCustomerAddresses {
	this := POSTCustomers201ResponseDataRelationshipsCustomerAddresses{}
	return &this
}

// NewPOSTCustomers201ResponseDataRelationshipsCustomerAddressesWithDefaults instantiates a new POSTCustomers201ResponseDataRelationshipsCustomerAddresses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomers201ResponseDataRelationshipsCustomerAddressesWithDefaults() *POSTCustomers201ResponseDataRelationshipsCustomerAddresses {
	this := POSTCustomers201ResponseDataRelationshipsCustomerAddresses{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerAddresses) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerAddresses) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerAddresses) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerAddresses) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerAddresses) GetData() POSTCustomers201ResponseDataRelationshipsCustomerAddressesData {
	if o == nil || IsNil(o.Data) {
		var ret POSTCustomers201ResponseDataRelationshipsCustomerAddressesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerAddresses) GetDataOk() (*POSTCustomers201ResponseDataRelationshipsCustomerAddressesData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerAddresses) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCustomers201ResponseDataRelationshipsCustomerAddressesData and assigns it to the Data field.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerAddresses) SetData(v POSTCustomers201ResponseDataRelationshipsCustomerAddressesData) {
	o.Data = &v
}

func (o POSTCustomers201ResponseDataRelationshipsCustomerAddresses) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCustomers201ResponseDataRelationshipsCustomerAddresses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTCustomers201ResponseDataRelationshipsCustomerAddresses struct {
	value *POSTCustomers201ResponseDataRelationshipsCustomerAddresses
	isSet bool
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsCustomerAddresses) Get() *POSTCustomers201ResponseDataRelationshipsCustomerAddresses {
	return v.value
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsCustomerAddresses) Set(val *POSTCustomers201ResponseDataRelationshipsCustomerAddresses) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsCustomerAddresses) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsCustomerAddresses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomers201ResponseDataRelationshipsCustomerAddresses(val *POSTCustomers201ResponseDataRelationshipsCustomerAddresses) *NullablePOSTCustomers201ResponseDataRelationshipsCustomerAddresses {
	return &NullablePOSTCustomers201ResponseDataRelationshipsCustomerAddresses{value: val, isSet: true}
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsCustomerAddresses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsCustomerAddresses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
