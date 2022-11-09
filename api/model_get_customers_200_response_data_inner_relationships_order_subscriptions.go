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

// GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions struct for GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions
type GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks                `json:"links,omitempty"`
	Data  []GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptionsDataInner `json:"data,omitempty"`
}

// NewGETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions instantiates a new GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions() *GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions {
	this := GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions{}
	return &this
}

// NewGETCustomers200ResponseDataInnerRelationshipsOrderSubscriptionsWithDefaults instantiates a new GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomers200ResponseDataInnerRelationshipsOrderSubscriptionsWithDefaults() *GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions {
	this := GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions) GetData() []GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptionsDataInner {
	if o == nil || o.Data == nil {
		var ret []GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptionsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions) GetDataOk() ([]GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptionsDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptionsDataInner and assigns it to the Data field.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions) SetData(v []GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptionsDataInner) {
	o.Data = v
}

func (o GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions struct {
	value *GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions
	isSet bool
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions) Get() *GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions {
	return v.value
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions) Set(val *GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions(val *GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions) *NullableGETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions {
	return &NullableGETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions{value: val, isSet: true}
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}