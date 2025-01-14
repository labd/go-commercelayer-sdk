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

// checks if the POSTMarkets201ResponseDataRelationshipsStoresData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTMarkets201ResponseDataRelationshipsStoresData{}

// POSTMarkets201ResponseDataRelationshipsStoresData struct for POSTMarkets201ResponseDataRelationshipsStoresData
type POSTMarkets201ResponseDataRelationshipsStoresData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTMarkets201ResponseDataRelationshipsStoresData instantiates a new POSTMarkets201ResponseDataRelationshipsStoresData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTMarkets201ResponseDataRelationshipsStoresData() *POSTMarkets201ResponseDataRelationshipsStoresData {
	this := POSTMarkets201ResponseDataRelationshipsStoresData{}
	return &this
}

// NewPOSTMarkets201ResponseDataRelationshipsStoresDataWithDefaults instantiates a new POSTMarkets201ResponseDataRelationshipsStoresData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTMarkets201ResponseDataRelationshipsStoresDataWithDefaults() *POSTMarkets201ResponseDataRelationshipsStoresData {
	this := POSTMarkets201ResponseDataRelationshipsStoresData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataRelationshipsStoresData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataRelationshipsStoresData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationshipsStoresData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTMarkets201ResponseDataRelationshipsStoresData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataRelationshipsStoresData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataRelationshipsStoresData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationshipsStoresData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTMarkets201ResponseDataRelationshipsStoresData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTMarkets201ResponseDataRelationshipsStoresData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTMarkets201ResponseDataRelationshipsStoresData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTMarkets201ResponseDataRelationshipsStoresData struct {
	value *POSTMarkets201ResponseDataRelationshipsStoresData
	isSet bool
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsStoresData) Get() *POSTMarkets201ResponseDataRelationshipsStoresData {
	return v.value
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsStoresData) Set(val *POSTMarkets201ResponseDataRelationshipsStoresData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsStoresData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsStoresData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTMarkets201ResponseDataRelationshipsStoresData(val *POSTMarkets201ResponseDataRelationshipsStoresData) *NullablePOSTMarkets201ResponseDataRelationshipsStoresData {
	return &NullablePOSTMarkets201ResponseDataRelationshipsStoresData{value: val, isSet: true}
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsStoresData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsStoresData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
