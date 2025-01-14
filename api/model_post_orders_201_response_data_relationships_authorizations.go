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

// checks if the POSTOrders201ResponseDataRelationshipsAuthorizations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrders201ResponseDataRelationshipsAuthorizations{}

// POSTOrders201ResponseDataRelationshipsAuthorizations struct for POSTOrders201ResponseDataRelationshipsAuthorizations
type POSTOrders201ResponseDataRelationshipsAuthorizations struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks   `json:"links,omitempty"`
	Data  *POSTOrders201ResponseDataRelationshipsAuthorizationsData `json:"data,omitempty"`
}

// NewPOSTOrders201ResponseDataRelationshipsAuthorizations instantiates a new POSTOrders201ResponseDataRelationshipsAuthorizations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrders201ResponseDataRelationshipsAuthorizations() *POSTOrders201ResponseDataRelationshipsAuthorizations {
	this := POSTOrders201ResponseDataRelationshipsAuthorizations{}
	return &this
}

// NewPOSTOrders201ResponseDataRelationshipsAuthorizationsWithDefaults instantiates a new POSTOrders201ResponseDataRelationshipsAuthorizations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrders201ResponseDataRelationshipsAuthorizationsWithDefaults() *POSTOrders201ResponseDataRelationshipsAuthorizations {
	this := POSTOrders201ResponseDataRelationshipsAuthorizations{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsAuthorizations) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsAuthorizations) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsAuthorizations) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTOrders201ResponseDataRelationshipsAuthorizations) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsAuthorizations) GetData() POSTOrders201ResponseDataRelationshipsAuthorizationsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTOrders201ResponseDataRelationshipsAuthorizationsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsAuthorizations) GetDataOk() (*POSTOrders201ResponseDataRelationshipsAuthorizationsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsAuthorizations) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTOrders201ResponseDataRelationshipsAuthorizationsData and assigns it to the Data field.
func (o *POSTOrders201ResponseDataRelationshipsAuthorizations) SetData(v POSTOrders201ResponseDataRelationshipsAuthorizationsData) {
	o.Data = &v
}

func (o POSTOrders201ResponseDataRelationshipsAuthorizations) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrders201ResponseDataRelationshipsAuthorizations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTOrders201ResponseDataRelationshipsAuthorizations struct {
	value *POSTOrders201ResponseDataRelationshipsAuthorizations
	isSet bool
}

func (v NullablePOSTOrders201ResponseDataRelationshipsAuthorizations) Get() *POSTOrders201ResponseDataRelationshipsAuthorizations {
	return v.value
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsAuthorizations) Set(val *POSTOrders201ResponseDataRelationshipsAuthorizations) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrders201ResponseDataRelationshipsAuthorizations) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsAuthorizations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrders201ResponseDataRelationshipsAuthorizations(val *POSTOrders201ResponseDataRelationshipsAuthorizations) *NullablePOSTOrders201ResponseDataRelationshipsAuthorizations {
	return &NullablePOSTOrders201ResponseDataRelationshipsAuthorizations{value: val, isSet: true}
}

func (v NullablePOSTOrders201ResponseDataRelationshipsAuthorizations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsAuthorizations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
