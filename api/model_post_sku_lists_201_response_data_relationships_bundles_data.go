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

// checks if the POSTSkuLists201ResponseDataRelationshipsBundlesData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTSkuLists201ResponseDataRelationshipsBundlesData{}

// POSTSkuLists201ResponseDataRelationshipsBundlesData struct for POSTSkuLists201ResponseDataRelationshipsBundlesData
type POSTSkuLists201ResponseDataRelationshipsBundlesData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTSkuLists201ResponseDataRelationshipsBundlesData instantiates a new POSTSkuLists201ResponseDataRelationshipsBundlesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSkuLists201ResponseDataRelationshipsBundlesData() *POSTSkuLists201ResponseDataRelationshipsBundlesData {
	this := POSTSkuLists201ResponseDataRelationshipsBundlesData{}
	return &this
}

// NewPOSTSkuLists201ResponseDataRelationshipsBundlesDataWithDefaults instantiates a new POSTSkuLists201ResponseDataRelationshipsBundlesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSkuLists201ResponseDataRelationshipsBundlesDataWithDefaults() *POSTSkuLists201ResponseDataRelationshipsBundlesData {
	this := POSTSkuLists201ResponseDataRelationshipsBundlesData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkuLists201ResponseDataRelationshipsBundlesData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkuLists201ResponseDataRelationshipsBundlesData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTSkuLists201ResponseDataRelationshipsBundlesData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTSkuLists201ResponseDataRelationshipsBundlesData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkuLists201ResponseDataRelationshipsBundlesData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkuLists201ResponseDataRelationshipsBundlesData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTSkuLists201ResponseDataRelationshipsBundlesData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTSkuLists201ResponseDataRelationshipsBundlesData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTSkuLists201ResponseDataRelationshipsBundlesData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTSkuLists201ResponseDataRelationshipsBundlesData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTSkuLists201ResponseDataRelationshipsBundlesData struct {
	value *POSTSkuLists201ResponseDataRelationshipsBundlesData
	isSet bool
}

func (v NullablePOSTSkuLists201ResponseDataRelationshipsBundlesData) Get() *POSTSkuLists201ResponseDataRelationshipsBundlesData {
	return v.value
}

func (v *NullablePOSTSkuLists201ResponseDataRelationshipsBundlesData) Set(val *POSTSkuLists201ResponseDataRelationshipsBundlesData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSkuLists201ResponseDataRelationshipsBundlesData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSkuLists201ResponseDataRelationshipsBundlesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSkuLists201ResponseDataRelationshipsBundlesData(val *POSTSkuLists201ResponseDataRelationshipsBundlesData) *NullablePOSTSkuLists201ResponseDataRelationshipsBundlesData {
	return &NullablePOSTSkuLists201ResponseDataRelationshipsBundlesData{value: val, isSet: true}
}

func (v NullablePOSTSkuLists201ResponseDataRelationshipsBundlesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSkuLists201ResponseDataRelationshipsBundlesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
