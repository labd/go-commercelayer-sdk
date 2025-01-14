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

// checks if the POSTLineItems201ResponseDataRelationshipsGiftCardData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTLineItems201ResponseDataRelationshipsGiftCardData{}

// POSTLineItems201ResponseDataRelationshipsGiftCardData struct for POSTLineItems201ResponseDataRelationshipsGiftCardData
type POSTLineItems201ResponseDataRelationshipsGiftCardData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTLineItems201ResponseDataRelationshipsGiftCardData instantiates a new POSTLineItems201ResponseDataRelationshipsGiftCardData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTLineItems201ResponseDataRelationshipsGiftCardData() *POSTLineItems201ResponseDataRelationshipsGiftCardData {
	this := POSTLineItems201ResponseDataRelationshipsGiftCardData{}
	return &this
}

// NewPOSTLineItems201ResponseDataRelationshipsGiftCardDataWithDefaults instantiates a new POSTLineItems201ResponseDataRelationshipsGiftCardData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTLineItems201ResponseDataRelationshipsGiftCardDataWithDefaults() *POSTLineItems201ResponseDataRelationshipsGiftCardData {
	this := POSTLineItems201ResponseDataRelationshipsGiftCardData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTLineItems201ResponseDataRelationshipsGiftCardData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTLineItems201ResponseDataRelationshipsGiftCardData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationshipsGiftCardData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTLineItems201ResponseDataRelationshipsGiftCardData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTLineItems201ResponseDataRelationshipsGiftCardData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTLineItems201ResponseDataRelationshipsGiftCardData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationshipsGiftCardData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTLineItems201ResponseDataRelationshipsGiftCardData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTLineItems201ResponseDataRelationshipsGiftCardData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTLineItems201ResponseDataRelationshipsGiftCardData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTLineItems201ResponseDataRelationshipsGiftCardData struct {
	value *POSTLineItems201ResponseDataRelationshipsGiftCardData
	isSet bool
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsGiftCardData) Get() *POSTLineItems201ResponseDataRelationshipsGiftCardData {
	return v.value
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsGiftCardData) Set(val *POSTLineItems201ResponseDataRelationshipsGiftCardData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsGiftCardData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsGiftCardData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTLineItems201ResponseDataRelationshipsGiftCardData(val *POSTLineItems201ResponseDataRelationshipsGiftCardData) *NullablePOSTLineItems201ResponseDataRelationshipsGiftCardData {
	return &NullablePOSTLineItems201ResponseDataRelationshipsGiftCardData{value: val, isSet: true}
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsGiftCardData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsGiftCardData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
