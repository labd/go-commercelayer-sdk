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

// GETCustomers200ResponseDataInnerRelationshipsReturns struct for GETCustomers200ResponseDataInnerRelationshipsReturns
type GETCustomers200ResponseDataInnerRelationshipsReturns struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks     `json:"links,omitempty"`
	Data  []GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner `json:"data,omitempty"`
}

// NewGETCustomers200ResponseDataInnerRelationshipsReturns instantiates a new GETCustomers200ResponseDataInnerRelationshipsReturns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomers200ResponseDataInnerRelationshipsReturns() *GETCustomers200ResponseDataInnerRelationshipsReturns {
	this := GETCustomers200ResponseDataInnerRelationshipsReturns{}
	return &this
}

// NewGETCustomers200ResponseDataInnerRelationshipsReturnsWithDefaults instantiates a new GETCustomers200ResponseDataInnerRelationshipsReturns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomers200ResponseDataInnerRelationshipsReturnsWithDefaults() *GETCustomers200ResponseDataInnerRelationshipsReturns {
	this := GETCustomers200ResponseDataInnerRelationshipsReturns{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETCustomers200ResponseDataInnerRelationshipsReturns) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsReturns) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsReturns) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETCustomers200ResponseDataInnerRelationshipsReturns) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCustomers200ResponseDataInnerRelationshipsReturns) GetData() []GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner {
	if o == nil || o.Data == nil {
		var ret []GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsReturns) GetDataOk() ([]GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsReturns) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner and assigns it to the Data field.
func (o *GETCustomers200ResponseDataInnerRelationshipsReturns) SetData(v []GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner) {
	o.Data = v
}

func (o GETCustomers200ResponseDataInnerRelationshipsReturns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomers200ResponseDataInnerRelationshipsReturns struct {
	value *GETCustomers200ResponseDataInnerRelationshipsReturns
	isSet bool
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsReturns) Get() *GETCustomers200ResponseDataInnerRelationshipsReturns {
	return v.value
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsReturns) Set(val *GETCustomers200ResponseDataInnerRelationshipsReturns) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsReturns) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsReturns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomers200ResponseDataInnerRelationshipsReturns(val *GETCustomers200ResponseDataInnerRelationshipsReturns) *NullableGETCustomers200ResponseDataInnerRelationshipsReturns {
	return &NullableGETCustomers200ResponseDataInnerRelationshipsReturns{value: val, isSet: true}
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsReturns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsReturns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
