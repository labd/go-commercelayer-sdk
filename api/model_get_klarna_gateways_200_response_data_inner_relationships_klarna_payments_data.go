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

// GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData struct for GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData
type GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData instantiates a new GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData() *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData {
	this := GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData{}
	return &this
}

// NewGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsDataWithDefaults instantiates a new GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsDataWithDefaults() *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData {
	this := GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData) SetId(v string) {
	o.Id = &v
}

func (o GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData struct {
	value *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData
	isSet bool
}

func (v NullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData) Get() *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData {
	return v.value
}

func (v *NullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData) Set(val *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData(val *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData) *NullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData {
	return &NullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData{value: val, isSet: true}
}

func (v NullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
