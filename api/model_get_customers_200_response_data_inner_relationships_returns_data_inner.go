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

// GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner struct for GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner
type GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETCustomers200ResponseDataInnerRelationshipsReturnsDataInner instantiates a new GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomers200ResponseDataInnerRelationshipsReturnsDataInner() *GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner {
	this := GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner{}
	return &this
}

// NewGETCustomers200ResponseDataInnerRelationshipsReturnsDataInnerWithDefaults instantiates a new GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomers200ResponseDataInnerRelationshipsReturnsDataInnerWithDefaults() *GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner {
	this := GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner) SetId(v string) {
	o.Id = &v
}

func (o GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomers200ResponseDataInnerRelationshipsReturnsDataInner struct {
	value *GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner
	isSet bool
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsReturnsDataInner) Get() *GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner {
	return v.value
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsReturnsDataInner) Set(val *GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsReturnsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsReturnsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomers200ResponseDataInnerRelationshipsReturnsDataInner(val *GETCustomers200ResponseDataInnerRelationshipsReturnsDataInner) *NullableGETCustomers200ResponseDataInnerRelationshipsReturnsDataInner {
	return &NullableGETCustomers200ResponseDataInnerRelationshipsReturnsDataInner{value: val, isSet: true}
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsReturnsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsReturnsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
