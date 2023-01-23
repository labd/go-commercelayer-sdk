/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETReturns200ResponseDataInnerRelationshipsDestinationAddressData struct for GETReturns200ResponseDataInnerRelationshipsDestinationAddressData
type GETReturns200ResponseDataInnerRelationshipsDestinationAddressData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETReturns200ResponseDataInnerRelationshipsDestinationAddressData instantiates a new GETReturns200ResponseDataInnerRelationshipsDestinationAddressData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETReturns200ResponseDataInnerRelationshipsDestinationAddressData() *GETReturns200ResponseDataInnerRelationshipsDestinationAddressData {
	this := GETReturns200ResponseDataInnerRelationshipsDestinationAddressData{}
	return &this
}

// NewGETReturns200ResponseDataInnerRelationshipsDestinationAddressDataWithDefaults instantiates a new GETReturns200ResponseDataInnerRelationshipsDestinationAddressData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETReturns200ResponseDataInnerRelationshipsDestinationAddressDataWithDefaults() *GETReturns200ResponseDataInnerRelationshipsDestinationAddressData {
	this := GETReturns200ResponseDataInnerRelationshipsDestinationAddressData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETReturns200ResponseDataInnerRelationshipsDestinationAddressData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturns200ResponseDataInnerRelationshipsDestinationAddressData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETReturns200ResponseDataInnerRelationshipsDestinationAddressData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETReturns200ResponseDataInnerRelationshipsDestinationAddressData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETReturns200ResponseDataInnerRelationshipsDestinationAddressData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturns200ResponseDataInnerRelationshipsDestinationAddressData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETReturns200ResponseDataInnerRelationshipsDestinationAddressData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETReturns200ResponseDataInnerRelationshipsDestinationAddressData) SetId(v string) {
	o.Id = &v
}

func (o GETReturns200ResponseDataInnerRelationshipsDestinationAddressData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETReturns200ResponseDataInnerRelationshipsDestinationAddressData struct {
	value *GETReturns200ResponseDataInnerRelationshipsDestinationAddressData
	isSet bool
}

func (v NullableGETReturns200ResponseDataInnerRelationshipsDestinationAddressData) Get() *GETReturns200ResponseDataInnerRelationshipsDestinationAddressData {
	return v.value
}

func (v *NullableGETReturns200ResponseDataInnerRelationshipsDestinationAddressData) Set(val *GETReturns200ResponseDataInnerRelationshipsDestinationAddressData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETReturns200ResponseDataInnerRelationshipsDestinationAddressData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETReturns200ResponseDataInnerRelationshipsDestinationAddressData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETReturns200ResponseDataInnerRelationshipsDestinationAddressData(val *GETReturns200ResponseDataInnerRelationshipsDestinationAddressData) *NullableGETReturns200ResponseDataInnerRelationshipsDestinationAddressData {
	return &NullableGETReturns200ResponseDataInnerRelationshipsDestinationAddressData{value: val, isSet: true}
}

func (v NullableGETReturns200ResponseDataInnerRelationshipsDestinationAddressData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETReturns200ResponseDataInnerRelationshipsDestinationAddressData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
