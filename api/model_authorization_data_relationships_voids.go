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

// checks if the AuthorizationDataRelationshipsVoids type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationDataRelationshipsVoids{}

// AuthorizationDataRelationshipsVoids struct for AuthorizationDataRelationshipsVoids
type AuthorizationDataRelationshipsVoids struct {
	Data *AuthorizationDataRelationshipsVoidsData `json:"data,omitempty"`
}

// NewAuthorizationDataRelationshipsVoids instantiates a new AuthorizationDataRelationshipsVoids object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationDataRelationshipsVoids() *AuthorizationDataRelationshipsVoids {
	this := AuthorizationDataRelationshipsVoids{}
	return &this
}

// NewAuthorizationDataRelationshipsVoidsWithDefaults instantiates a new AuthorizationDataRelationshipsVoids object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationDataRelationshipsVoidsWithDefaults() *AuthorizationDataRelationshipsVoids {
	this := AuthorizationDataRelationshipsVoids{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AuthorizationDataRelationshipsVoids) GetData() AuthorizationDataRelationshipsVoidsData {
	if o == nil || IsNil(o.Data) {
		var ret AuthorizationDataRelationshipsVoidsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDataRelationshipsVoids) GetDataOk() (*AuthorizationDataRelationshipsVoidsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AuthorizationDataRelationshipsVoids) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AuthorizationDataRelationshipsVoidsData and assigns it to the Data field.
func (o *AuthorizationDataRelationshipsVoids) SetData(v AuthorizationDataRelationshipsVoidsData) {
	o.Data = &v
}

func (o AuthorizationDataRelationshipsVoids) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationDataRelationshipsVoids) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAuthorizationDataRelationshipsVoids struct {
	value *AuthorizationDataRelationshipsVoids
	isSet bool
}

func (v NullableAuthorizationDataRelationshipsVoids) Get() *AuthorizationDataRelationshipsVoids {
	return v.value
}

func (v *NullableAuthorizationDataRelationshipsVoids) Set(val *AuthorizationDataRelationshipsVoids) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationDataRelationshipsVoids) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationDataRelationshipsVoids) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationDataRelationshipsVoids(val *AuthorizationDataRelationshipsVoids) *NullableAuthorizationDataRelationshipsVoids {
	return &NullableAuthorizationDataRelationshipsVoids{value: val, isSet: true}
}

func (v NullableAuthorizationDataRelationshipsVoids) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationDataRelationshipsVoids) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
