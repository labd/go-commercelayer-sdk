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

// checks if the ShippingZoneDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingZoneDataRelationships{}

// ShippingZoneDataRelationships struct for ShippingZoneDataRelationships
type ShippingZoneDataRelationships struct {
	Attachments *AuthorizationDataRelationshipsAttachments `json:"attachments,omitempty"`
	Versions    *AddressDataRelationshipsVersions          `json:"versions,omitempty"`
}

// NewShippingZoneDataRelationships instantiates a new ShippingZoneDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingZoneDataRelationships() *ShippingZoneDataRelationships {
	this := ShippingZoneDataRelationships{}
	return &this
}

// NewShippingZoneDataRelationshipsWithDefaults instantiates a new ShippingZoneDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingZoneDataRelationshipsWithDefaults() *ShippingZoneDataRelationships {
	this := ShippingZoneDataRelationships{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *ShippingZoneDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AuthorizationDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingZoneDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *ShippingZoneDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AuthorizationDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *ShippingZoneDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *ShippingZoneDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingZoneDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *ShippingZoneDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *ShippingZoneDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o ShippingZoneDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingZoneDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableShippingZoneDataRelationships struct {
	value *ShippingZoneDataRelationships
	isSet bool
}

func (v NullableShippingZoneDataRelationships) Get() *ShippingZoneDataRelationships {
	return v.value
}

func (v *NullableShippingZoneDataRelationships) Set(val *ShippingZoneDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingZoneDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingZoneDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingZoneDataRelationships(val *ShippingZoneDataRelationships) *NullableShippingZoneDataRelationships {
	return &NullableShippingZoneDataRelationships{value: val, isSet: true}
}

func (v NullableShippingZoneDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingZoneDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
