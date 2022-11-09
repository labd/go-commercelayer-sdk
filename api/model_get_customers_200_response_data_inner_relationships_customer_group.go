/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETCustomers200ResponseDataInnerRelationshipsCustomerGroup struct for GETCustomers200ResponseDataInnerRelationshipsCustomerGroup
type GETCustomers200ResponseDataInnerRelationshipsCustomerGroup struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks     `json:"links,omitempty"`
	Data  *GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData `json:"data,omitempty"`
}

// NewGETCustomers200ResponseDataInnerRelationshipsCustomerGroup instantiates a new GETCustomers200ResponseDataInnerRelationshipsCustomerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomers200ResponseDataInnerRelationshipsCustomerGroup() *GETCustomers200ResponseDataInnerRelationshipsCustomerGroup {
	this := GETCustomers200ResponseDataInnerRelationshipsCustomerGroup{}
	return &this
}

// NewGETCustomers200ResponseDataInnerRelationshipsCustomerGroupWithDefaults instantiates a new GETCustomers200ResponseDataInnerRelationshipsCustomerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomers200ResponseDataInnerRelationshipsCustomerGroupWithDefaults() *GETCustomers200ResponseDataInnerRelationshipsCustomerGroup {
	this := GETCustomers200ResponseDataInnerRelationshipsCustomerGroup{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerGroup) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerGroup) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerGroup) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerGroup) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerGroup) GetData() GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData {
	if o == nil || o.Data == nil {
		var ret GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerGroup) GetDataOk() (*GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerGroup) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData and assigns it to the Data field.
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerGroup) SetData(v GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData) {
	o.Data = &v
}

func (o GETCustomers200ResponseDataInnerRelationshipsCustomerGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroup struct {
	value *GETCustomers200ResponseDataInnerRelationshipsCustomerGroup
	isSet bool
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroup) Get() *GETCustomers200ResponseDataInnerRelationshipsCustomerGroup {
	return v.value
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroup) Set(val *GETCustomers200ResponseDataInnerRelationshipsCustomerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroup(val *GETCustomers200ResponseDataInnerRelationshipsCustomerGroup) *NullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroup {
	return &NullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroup{value: val, isSet: true}
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
