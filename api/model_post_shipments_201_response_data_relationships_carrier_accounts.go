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

// checks if the POSTShipments201ResponseDataRelationshipsCarrierAccounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTShipments201ResponseDataRelationshipsCarrierAccounts{}

// POSTShipments201ResponseDataRelationshipsCarrierAccounts struct for POSTShipments201ResponseDataRelationshipsCarrierAccounts
type POSTShipments201ResponseDataRelationshipsCarrierAccounts struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks       `json:"links,omitempty"`
	Data  *POSTShipments201ResponseDataRelationshipsCarrierAccountsData `json:"data,omitempty"`
}

// NewPOSTShipments201ResponseDataRelationshipsCarrierAccounts instantiates a new POSTShipments201ResponseDataRelationshipsCarrierAccounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShipments201ResponseDataRelationshipsCarrierAccounts() *POSTShipments201ResponseDataRelationshipsCarrierAccounts {
	this := POSTShipments201ResponseDataRelationshipsCarrierAccounts{}
	return &this
}

// NewPOSTShipments201ResponseDataRelationshipsCarrierAccountsWithDefaults instantiates a new POSTShipments201ResponseDataRelationshipsCarrierAccounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShipments201ResponseDataRelationshipsCarrierAccountsWithDefaults() *POSTShipments201ResponseDataRelationshipsCarrierAccounts {
	this := POSTShipments201ResponseDataRelationshipsCarrierAccounts{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTShipments201ResponseDataRelationshipsCarrierAccounts) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShipments201ResponseDataRelationshipsCarrierAccounts) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTShipments201ResponseDataRelationshipsCarrierAccounts) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTShipments201ResponseDataRelationshipsCarrierAccounts) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTShipments201ResponseDataRelationshipsCarrierAccounts) GetData() POSTShipments201ResponseDataRelationshipsCarrierAccountsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTShipments201ResponseDataRelationshipsCarrierAccountsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShipments201ResponseDataRelationshipsCarrierAccounts) GetDataOk() (*POSTShipments201ResponseDataRelationshipsCarrierAccountsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTShipments201ResponseDataRelationshipsCarrierAccounts) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTShipments201ResponseDataRelationshipsCarrierAccountsData and assigns it to the Data field.
func (o *POSTShipments201ResponseDataRelationshipsCarrierAccounts) SetData(v POSTShipments201ResponseDataRelationshipsCarrierAccountsData) {
	o.Data = &v
}

func (o POSTShipments201ResponseDataRelationshipsCarrierAccounts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTShipments201ResponseDataRelationshipsCarrierAccounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTShipments201ResponseDataRelationshipsCarrierAccounts struct {
	value *POSTShipments201ResponseDataRelationshipsCarrierAccounts
	isSet bool
}

func (v NullablePOSTShipments201ResponseDataRelationshipsCarrierAccounts) Get() *POSTShipments201ResponseDataRelationshipsCarrierAccounts {
	return v.value
}

func (v *NullablePOSTShipments201ResponseDataRelationshipsCarrierAccounts) Set(val *POSTShipments201ResponseDataRelationshipsCarrierAccounts) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShipments201ResponseDataRelationshipsCarrierAccounts) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShipments201ResponseDataRelationshipsCarrierAccounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShipments201ResponseDataRelationshipsCarrierAccounts(val *POSTShipments201ResponseDataRelationshipsCarrierAccounts) *NullablePOSTShipments201ResponseDataRelationshipsCarrierAccounts {
	return &NullablePOSTShipments201ResponseDataRelationshipsCarrierAccounts{value: val, isSet: true}
}

func (v NullablePOSTShipments201ResponseDataRelationshipsCarrierAccounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShipments201ResponseDataRelationshipsCarrierAccounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
