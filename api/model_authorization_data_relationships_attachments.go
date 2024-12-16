/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AuthorizationDataRelationshipsAttachments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationDataRelationshipsAttachments{}

// AuthorizationDataRelationshipsAttachments struct for AuthorizationDataRelationshipsAttachments
type AuthorizationDataRelationshipsAttachments struct {
	Data *AuthorizationDataRelationshipsAttachmentsData `json:"data,omitempty"`
}

// NewAuthorizationDataRelationshipsAttachments instantiates a new AuthorizationDataRelationshipsAttachments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationDataRelationshipsAttachments() *AuthorizationDataRelationshipsAttachments {
	this := AuthorizationDataRelationshipsAttachments{}
	return &this
}

// NewAuthorizationDataRelationshipsAttachmentsWithDefaults instantiates a new AuthorizationDataRelationshipsAttachments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationDataRelationshipsAttachmentsWithDefaults() *AuthorizationDataRelationshipsAttachments {
	this := AuthorizationDataRelationshipsAttachments{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AuthorizationDataRelationshipsAttachments) GetData() AuthorizationDataRelationshipsAttachmentsData {
	if o == nil || IsNil(o.Data) {
		var ret AuthorizationDataRelationshipsAttachmentsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDataRelationshipsAttachments) GetDataOk() (*AuthorizationDataRelationshipsAttachmentsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AuthorizationDataRelationshipsAttachments) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AuthorizationDataRelationshipsAttachmentsData and assigns it to the Data field.
func (o *AuthorizationDataRelationshipsAttachments) SetData(v AuthorizationDataRelationshipsAttachmentsData) {
	o.Data = &v
}

func (o AuthorizationDataRelationshipsAttachments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationDataRelationshipsAttachments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAuthorizationDataRelationshipsAttachments struct {
	value *AuthorizationDataRelationshipsAttachments
	isSet bool
}

func (v NullableAuthorizationDataRelationshipsAttachments) Get() *AuthorizationDataRelationshipsAttachments {
	return v.value
}

func (v *NullableAuthorizationDataRelationshipsAttachments) Set(val *AuthorizationDataRelationshipsAttachments) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationDataRelationshipsAttachments) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationDataRelationshipsAttachments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationDataRelationshipsAttachments(val *AuthorizationDataRelationshipsAttachments) *NullableAuthorizationDataRelationshipsAttachments {
	return &NullableAuthorizationDataRelationshipsAttachments{value: val, isSet: true}
}

func (v NullableAuthorizationDataRelationshipsAttachments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationDataRelationshipsAttachments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
