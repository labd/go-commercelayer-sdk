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

// GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner struct for GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner
type GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETPriceLists200ResponseDataInnerRelationshipsPricesDataInner instantiates a new GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPriceLists200ResponseDataInnerRelationshipsPricesDataInner() *GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner {
	this := GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner{}
	return &this
}

// NewGETPriceLists200ResponseDataInnerRelationshipsPricesDataInnerWithDefaults instantiates a new GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPriceLists200ResponseDataInnerRelationshipsPricesDataInnerWithDefaults() *GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner {
	this := GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner) SetId(v string) {
	o.Id = &v
}

func (o GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETPriceLists200ResponseDataInnerRelationshipsPricesDataInner struct {
	value *GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner
	isSet bool
}

func (v NullableGETPriceLists200ResponseDataInnerRelationshipsPricesDataInner) Get() *GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner {
	return v.value
}

func (v *NullableGETPriceLists200ResponseDataInnerRelationshipsPricesDataInner) Set(val *GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPriceLists200ResponseDataInnerRelationshipsPricesDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPriceLists200ResponseDataInnerRelationshipsPricesDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPriceLists200ResponseDataInnerRelationshipsPricesDataInner(val *GETPriceLists200ResponseDataInnerRelationshipsPricesDataInner) *NullableGETPriceLists200ResponseDataInnerRelationshipsPricesDataInner {
	return &NullableGETPriceLists200ResponseDataInnerRelationshipsPricesDataInner{value: val, isSet: true}
}

func (v NullableGETPriceLists200ResponseDataInnerRelationshipsPricesDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPriceLists200ResponseDataInnerRelationshipsPricesDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
