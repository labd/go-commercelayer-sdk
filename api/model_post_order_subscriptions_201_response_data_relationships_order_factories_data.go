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

// checks if the POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData{}

// POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData struct for POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData
type POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData instantiates a new POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData() *POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData {
	this := POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData{}
	return &this
}

// NewPOSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesDataWithDefaults instantiates a new POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesDataWithDefaults() *POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData {
	this := POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData struct {
	value *POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData
	isSet bool
}

func (v NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData) Get() *POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData {
	return v.value
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData) Set(val *POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData(val *POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData) *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData {
	return &NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData{value: val, isSet: true}
}

func (v NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderFactoriesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
