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

// checks if the POSTShippingZones201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTShippingZones201ResponseDataRelationships{}

// POSTShippingZones201ResponseDataRelationships struct for POSTShippingZones201ResponseDataRelationships
type POSTShippingZones201ResponseDataRelationships struct {
	Attachments *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments `json:"attachments,omitempty"`
	Versions    *POSTAddresses201ResponseDataRelationshipsVersions                       `json:"versions,omitempty"`
}

// NewPOSTShippingZones201ResponseDataRelationships instantiates a new POSTShippingZones201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShippingZones201ResponseDataRelationships() *POSTShippingZones201ResponseDataRelationships {
	this := POSTShippingZones201ResponseDataRelationships{}
	return &this
}

// NewPOSTShippingZones201ResponseDataRelationshipsWithDefaults instantiates a new POSTShippingZones201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShippingZones201ResponseDataRelationshipsWithDefaults() *POSTShippingZones201ResponseDataRelationships {
	this := POSTShippingZones201ResponseDataRelationships{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *POSTShippingZones201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingZones201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *POSTShippingZones201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTShippingZones201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingZones201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTShippingZones201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

func (o POSTShippingZones201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTShippingZones201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullablePOSTShippingZones201ResponseDataRelationships struct {
	value *POSTShippingZones201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTShippingZones201ResponseDataRelationships) Get() *POSTShippingZones201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTShippingZones201ResponseDataRelationships) Set(val *POSTShippingZones201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShippingZones201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShippingZones201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShippingZones201ResponseDataRelationships(val *POSTShippingZones201ResponseDataRelationships) *NullablePOSTShippingZones201ResponseDataRelationships {
	return &NullablePOSTShippingZones201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTShippingZones201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShippingZones201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
