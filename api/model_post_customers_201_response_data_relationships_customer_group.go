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

// checks if the POSTCustomers201ResponseDataRelationshipsCustomerGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCustomers201ResponseDataRelationshipsCustomerGroup{}

// POSTCustomers201ResponseDataRelationshipsCustomerGroup struct for POSTCustomers201ResponseDataRelationshipsCustomerGroup
type POSTCustomers201ResponseDataRelationshipsCustomerGroup struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks     `json:"links,omitempty"`
	Data  *POSTCustomers201ResponseDataRelationshipsCustomerGroupData `json:"data,omitempty"`
}

// NewPOSTCustomers201ResponseDataRelationshipsCustomerGroup instantiates a new POSTCustomers201ResponseDataRelationshipsCustomerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomers201ResponseDataRelationshipsCustomerGroup() *POSTCustomers201ResponseDataRelationshipsCustomerGroup {
	this := POSTCustomers201ResponseDataRelationshipsCustomerGroup{}
	return &this
}

// NewPOSTCustomers201ResponseDataRelationshipsCustomerGroupWithDefaults instantiates a new POSTCustomers201ResponseDataRelationshipsCustomerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomers201ResponseDataRelationshipsCustomerGroupWithDefaults() *POSTCustomers201ResponseDataRelationshipsCustomerGroup {
	this := POSTCustomers201ResponseDataRelationshipsCustomerGroup{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerGroup) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerGroup) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerGroup) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerGroup) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerGroup) GetData() POSTCustomers201ResponseDataRelationshipsCustomerGroupData {
	if o == nil || IsNil(o.Data) {
		var ret POSTCustomers201ResponseDataRelationshipsCustomerGroupData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerGroup) GetDataOk() (*POSTCustomers201ResponseDataRelationshipsCustomerGroupData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerGroup) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCustomers201ResponseDataRelationshipsCustomerGroupData and assigns it to the Data field.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerGroup) SetData(v POSTCustomers201ResponseDataRelationshipsCustomerGroupData) {
	o.Data = &v
}

func (o POSTCustomers201ResponseDataRelationshipsCustomerGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCustomers201ResponseDataRelationshipsCustomerGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTCustomers201ResponseDataRelationshipsCustomerGroup struct {
	value *POSTCustomers201ResponseDataRelationshipsCustomerGroup
	isSet bool
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsCustomerGroup) Get() *POSTCustomers201ResponseDataRelationshipsCustomerGroup {
	return v.value
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsCustomerGroup) Set(val *POSTCustomers201ResponseDataRelationshipsCustomerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsCustomerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsCustomerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomers201ResponseDataRelationshipsCustomerGroup(val *POSTCustomers201ResponseDataRelationshipsCustomerGroup) *NullablePOSTCustomers201ResponseDataRelationshipsCustomerGroup {
	return &NullablePOSTCustomers201ResponseDataRelationshipsCustomerGroup{value: val, isSet: true}
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsCustomerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsCustomerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
