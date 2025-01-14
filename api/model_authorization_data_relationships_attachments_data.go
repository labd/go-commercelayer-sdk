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

// checks if the AuthorizationDataRelationshipsAttachmentsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationDataRelationshipsAttachmentsData{}

// AuthorizationDataRelationshipsAttachmentsData struct for AuthorizationDataRelationshipsAttachmentsData
type AuthorizationDataRelationshipsAttachmentsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource's id
	Id interface{} `json:"id,omitempty"`
}

// NewAuthorizationDataRelationshipsAttachmentsData instantiates a new AuthorizationDataRelationshipsAttachmentsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationDataRelationshipsAttachmentsData() *AuthorizationDataRelationshipsAttachmentsData {
	this := AuthorizationDataRelationshipsAttachmentsData{}
	return &this
}

// NewAuthorizationDataRelationshipsAttachmentsDataWithDefaults instantiates a new AuthorizationDataRelationshipsAttachmentsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationDataRelationshipsAttachmentsDataWithDefaults() *AuthorizationDataRelationshipsAttachmentsData {
	this := AuthorizationDataRelationshipsAttachmentsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthorizationDataRelationshipsAttachmentsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthorizationDataRelationshipsAttachmentsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuthorizationDataRelationshipsAttachmentsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *AuthorizationDataRelationshipsAttachmentsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthorizationDataRelationshipsAttachmentsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthorizationDataRelationshipsAttachmentsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthorizationDataRelationshipsAttachmentsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *AuthorizationDataRelationshipsAttachmentsData) SetId(v interface{}) {
	o.Id = v
}

func (o AuthorizationDataRelationshipsAttachmentsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationDataRelationshipsAttachmentsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableAuthorizationDataRelationshipsAttachmentsData struct {
	value *AuthorizationDataRelationshipsAttachmentsData
	isSet bool
}

func (v NullableAuthorizationDataRelationshipsAttachmentsData) Get() *AuthorizationDataRelationshipsAttachmentsData {
	return v.value
}

func (v *NullableAuthorizationDataRelationshipsAttachmentsData) Set(val *AuthorizationDataRelationshipsAttachmentsData) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationDataRelationshipsAttachmentsData) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationDataRelationshipsAttachmentsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationDataRelationshipsAttachmentsData(val *AuthorizationDataRelationshipsAttachmentsData) *NullableAuthorizationDataRelationshipsAttachmentsData {
	return &NullableAuthorizationDataRelationshipsAttachmentsData{value: val, isSet: true}
}

func (v NullableAuthorizationDataRelationshipsAttachmentsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationDataRelationshipsAttachmentsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
