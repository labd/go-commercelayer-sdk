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

// checks if the PackageDataRelationshipsParcelsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageDataRelationshipsParcelsData{}

// PackageDataRelationshipsParcelsData struct for PackageDataRelationshipsParcelsData
type PackageDataRelationshipsParcelsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource's id
	Id interface{} `json:"id,omitempty"`
}

// NewPackageDataRelationshipsParcelsData instantiates a new PackageDataRelationshipsParcelsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageDataRelationshipsParcelsData() *PackageDataRelationshipsParcelsData {
	this := PackageDataRelationshipsParcelsData{}
	return &this
}

// NewPackageDataRelationshipsParcelsDataWithDefaults instantiates a new PackageDataRelationshipsParcelsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageDataRelationshipsParcelsDataWithDefaults() *PackageDataRelationshipsParcelsData {
	this := PackageDataRelationshipsParcelsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageDataRelationshipsParcelsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageDataRelationshipsParcelsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PackageDataRelationshipsParcelsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *PackageDataRelationshipsParcelsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageDataRelationshipsParcelsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageDataRelationshipsParcelsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PackageDataRelationshipsParcelsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *PackageDataRelationshipsParcelsData) SetId(v interface{}) {
	o.Id = v
}

func (o PackageDataRelationshipsParcelsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageDataRelationshipsParcelsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePackageDataRelationshipsParcelsData struct {
	value *PackageDataRelationshipsParcelsData
	isSet bool
}

func (v NullablePackageDataRelationshipsParcelsData) Get() *PackageDataRelationshipsParcelsData {
	return v.value
}

func (v *NullablePackageDataRelationshipsParcelsData) Set(val *PackageDataRelationshipsParcelsData) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageDataRelationshipsParcelsData) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageDataRelationshipsParcelsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageDataRelationshipsParcelsData(val *PackageDataRelationshipsParcelsData) *NullablePackageDataRelationshipsParcelsData {
	return &NullablePackageDataRelationshipsParcelsData{value: val, isSet: true}
}

func (v NullablePackageDataRelationshipsParcelsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageDataRelationshipsParcelsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
