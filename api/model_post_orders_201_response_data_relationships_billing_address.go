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

// checks if the POSTOrders201ResponseDataRelationshipsBillingAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrders201ResponseDataRelationshipsBillingAddress{}

// POSTOrders201ResponseDataRelationshipsBillingAddress struct for POSTOrders201ResponseDataRelationshipsBillingAddress
type POSTOrders201ResponseDataRelationshipsBillingAddress struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks   `json:"links,omitempty"`
	Data  *POSTOrders201ResponseDataRelationshipsBillingAddressData `json:"data,omitempty"`
}

// NewPOSTOrders201ResponseDataRelationshipsBillingAddress instantiates a new POSTOrders201ResponseDataRelationshipsBillingAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrders201ResponseDataRelationshipsBillingAddress() *POSTOrders201ResponseDataRelationshipsBillingAddress {
	this := POSTOrders201ResponseDataRelationshipsBillingAddress{}
	return &this
}

// NewPOSTOrders201ResponseDataRelationshipsBillingAddressWithDefaults instantiates a new POSTOrders201ResponseDataRelationshipsBillingAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrders201ResponseDataRelationshipsBillingAddressWithDefaults() *POSTOrders201ResponseDataRelationshipsBillingAddress {
	this := POSTOrders201ResponseDataRelationshipsBillingAddress{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsBillingAddress) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsBillingAddress) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsBillingAddress) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTOrders201ResponseDataRelationshipsBillingAddress) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsBillingAddress) GetData() POSTOrders201ResponseDataRelationshipsBillingAddressData {
	if o == nil || IsNil(o.Data) {
		var ret POSTOrders201ResponseDataRelationshipsBillingAddressData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsBillingAddress) GetDataOk() (*POSTOrders201ResponseDataRelationshipsBillingAddressData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsBillingAddress) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTOrders201ResponseDataRelationshipsBillingAddressData and assigns it to the Data field.
func (o *POSTOrders201ResponseDataRelationshipsBillingAddress) SetData(v POSTOrders201ResponseDataRelationshipsBillingAddressData) {
	o.Data = &v
}

func (o POSTOrders201ResponseDataRelationshipsBillingAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrders201ResponseDataRelationshipsBillingAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTOrders201ResponseDataRelationshipsBillingAddress struct {
	value *POSTOrders201ResponseDataRelationshipsBillingAddress
	isSet bool
}

func (v NullablePOSTOrders201ResponseDataRelationshipsBillingAddress) Get() *POSTOrders201ResponseDataRelationshipsBillingAddress {
	return v.value
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsBillingAddress) Set(val *POSTOrders201ResponseDataRelationshipsBillingAddress) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrders201ResponseDataRelationshipsBillingAddress) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsBillingAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrders201ResponseDataRelationshipsBillingAddress(val *POSTOrders201ResponseDataRelationshipsBillingAddress) *NullablePOSTOrders201ResponseDataRelationshipsBillingAddress {
	return &NullablePOSTOrders201ResponseDataRelationshipsBillingAddress{value: val, isSet: true}
}

func (v NullablePOSTOrders201ResponseDataRelationshipsBillingAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsBillingAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
