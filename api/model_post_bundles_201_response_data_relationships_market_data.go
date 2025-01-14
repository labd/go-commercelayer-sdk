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

// checks if the POSTBundles201ResponseDataRelationshipsMarketData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTBundles201ResponseDataRelationshipsMarketData{}

// POSTBundles201ResponseDataRelationshipsMarketData struct for POSTBundles201ResponseDataRelationshipsMarketData
type POSTBundles201ResponseDataRelationshipsMarketData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTBundles201ResponseDataRelationshipsMarketData instantiates a new POSTBundles201ResponseDataRelationshipsMarketData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTBundles201ResponseDataRelationshipsMarketData() *POSTBundles201ResponseDataRelationshipsMarketData {
	this := POSTBundles201ResponseDataRelationshipsMarketData{}
	return &this
}

// NewPOSTBundles201ResponseDataRelationshipsMarketDataWithDefaults instantiates a new POSTBundles201ResponseDataRelationshipsMarketData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTBundles201ResponseDataRelationshipsMarketDataWithDefaults() *POSTBundles201ResponseDataRelationshipsMarketData {
	this := POSTBundles201ResponseDataRelationshipsMarketData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTBundles201ResponseDataRelationshipsMarketData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTBundles201ResponseDataRelationshipsMarketData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTBundles201ResponseDataRelationshipsMarketData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTBundles201ResponseDataRelationshipsMarketData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTBundles201ResponseDataRelationshipsMarketData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTBundles201ResponseDataRelationshipsMarketData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTBundles201ResponseDataRelationshipsMarketData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTBundles201ResponseDataRelationshipsMarketData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTBundles201ResponseDataRelationshipsMarketData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTBundles201ResponseDataRelationshipsMarketData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTBundles201ResponseDataRelationshipsMarketData struct {
	value *POSTBundles201ResponseDataRelationshipsMarketData
	isSet bool
}

func (v NullablePOSTBundles201ResponseDataRelationshipsMarketData) Get() *POSTBundles201ResponseDataRelationshipsMarketData {
	return v.value
}

func (v *NullablePOSTBundles201ResponseDataRelationshipsMarketData) Set(val *POSTBundles201ResponseDataRelationshipsMarketData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTBundles201ResponseDataRelationshipsMarketData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTBundles201ResponseDataRelationshipsMarketData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTBundles201ResponseDataRelationshipsMarketData(val *POSTBundles201ResponseDataRelationshipsMarketData) *NullablePOSTBundles201ResponseDataRelationshipsMarketData {
	return &NullablePOSTBundles201ResponseDataRelationshipsMarketData{value: val, isSet: true}
}

func (v NullablePOSTBundles201ResponseDataRelationshipsMarketData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTBundles201ResponseDataRelationshipsMarketData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
