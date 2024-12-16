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

// checks if the POSTAddresses201ResponseDataRelationshipsGeocoderData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAddresses201ResponseDataRelationshipsGeocoderData{}

// POSTAddresses201ResponseDataRelationshipsGeocoderData struct for POSTAddresses201ResponseDataRelationshipsGeocoderData
type POSTAddresses201ResponseDataRelationshipsGeocoderData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTAddresses201ResponseDataRelationshipsGeocoderData instantiates a new POSTAddresses201ResponseDataRelationshipsGeocoderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAddresses201ResponseDataRelationshipsGeocoderData() *POSTAddresses201ResponseDataRelationshipsGeocoderData {
	this := POSTAddresses201ResponseDataRelationshipsGeocoderData{}
	return &this
}

// NewPOSTAddresses201ResponseDataRelationshipsGeocoderDataWithDefaults instantiates a new POSTAddresses201ResponseDataRelationshipsGeocoderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAddresses201ResponseDataRelationshipsGeocoderDataWithDefaults() *POSTAddresses201ResponseDataRelationshipsGeocoderData {
	this := POSTAddresses201ResponseDataRelationshipsGeocoderData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAddresses201ResponseDataRelationshipsGeocoderData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataRelationshipsGeocoderData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataRelationshipsGeocoderData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTAddresses201ResponseDataRelationshipsGeocoderData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAddresses201ResponseDataRelationshipsGeocoderData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAddresses201ResponseDataRelationshipsGeocoderData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTAddresses201ResponseDataRelationshipsGeocoderData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTAddresses201ResponseDataRelationshipsGeocoderData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTAddresses201ResponseDataRelationshipsGeocoderData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAddresses201ResponseDataRelationshipsGeocoderData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTAddresses201ResponseDataRelationshipsGeocoderData struct {
	value *POSTAddresses201ResponseDataRelationshipsGeocoderData
	isSet bool
}

func (v NullablePOSTAddresses201ResponseDataRelationshipsGeocoderData) Get() *POSTAddresses201ResponseDataRelationshipsGeocoderData {
	return v.value
}

func (v *NullablePOSTAddresses201ResponseDataRelationshipsGeocoderData) Set(val *POSTAddresses201ResponseDataRelationshipsGeocoderData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAddresses201ResponseDataRelationshipsGeocoderData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAddresses201ResponseDataRelationshipsGeocoderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAddresses201ResponseDataRelationshipsGeocoderData(val *POSTAddresses201ResponseDataRelationshipsGeocoderData) *NullablePOSTAddresses201ResponseDataRelationshipsGeocoderData {
	return &NullablePOSTAddresses201ResponseDataRelationshipsGeocoderData{value: val, isSet: true}
}

func (v NullablePOSTAddresses201ResponseDataRelationshipsGeocoderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAddresses201ResponseDataRelationshipsGeocoderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
