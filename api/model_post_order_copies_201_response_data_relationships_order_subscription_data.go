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

// checks if the POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData{}

// POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData struct for POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData
type POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData instantiates a new POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData() *POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData {
	this := POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData{}
	return &this
}

// NewPOSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionDataWithDefaults instantiates a new POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionDataWithDefaults() *POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData {
	this := POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData struct {
	value *POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData
	isSet bool
}

func (v NullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) Get() *POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData {
	return v.value
}

func (v *NullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) Set(val *POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData(val *POSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) *NullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData {
	return &NullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData{value: val, isSet: true}
}

func (v NullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderCopies201ResponseDataRelationshipsOrderSubscriptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
