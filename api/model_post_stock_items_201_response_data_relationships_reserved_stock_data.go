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

// checks if the POSTStockItems201ResponseDataRelationshipsReservedStockData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTStockItems201ResponseDataRelationshipsReservedStockData{}

// POSTStockItems201ResponseDataRelationshipsReservedStockData struct for POSTStockItems201ResponseDataRelationshipsReservedStockData
type POSTStockItems201ResponseDataRelationshipsReservedStockData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTStockItems201ResponseDataRelationshipsReservedStockData instantiates a new POSTStockItems201ResponseDataRelationshipsReservedStockData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStockItems201ResponseDataRelationshipsReservedStockData() *POSTStockItems201ResponseDataRelationshipsReservedStockData {
	this := POSTStockItems201ResponseDataRelationshipsReservedStockData{}
	return &this
}

// NewPOSTStockItems201ResponseDataRelationshipsReservedStockDataWithDefaults instantiates a new POSTStockItems201ResponseDataRelationshipsReservedStockData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStockItems201ResponseDataRelationshipsReservedStockDataWithDefaults() *POSTStockItems201ResponseDataRelationshipsReservedStockData {
	this := POSTStockItems201ResponseDataRelationshipsReservedStockData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStockItems201ResponseDataRelationshipsReservedStockData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStockItems201ResponseDataRelationshipsReservedStockData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTStockItems201ResponseDataRelationshipsReservedStockData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTStockItems201ResponseDataRelationshipsReservedStockData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStockItems201ResponseDataRelationshipsReservedStockData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStockItems201ResponseDataRelationshipsReservedStockData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTStockItems201ResponseDataRelationshipsReservedStockData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTStockItems201ResponseDataRelationshipsReservedStockData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTStockItems201ResponseDataRelationshipsReservedStockData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTStockItems201ResponseDataRelationshipsReservedStockData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTStockItems201ResponseDataRelationshipsReservedStockData struct {
	value *POSTStockItems201ResponseDataRelationshipsReservedStockData
	isSet bool
}

func (v NullablePOSTStockItems201ResponseDataRelationshipsReservedStockData) Get() *POSTStockItems201ResponseDataRelationshipsReservedStockData {
	return v.value
}

func (v *NullablePOSTStockItems201ResponseDataRelationshipsReservedStockData) Set(val *POSTStockItems201ResponseDataRelationshipsReservedStockData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStockItems201ResponseDataRelationshipsReservedStockData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStockItems201ResponseDataRelationshipsReservedStockData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStockItems201ResponseDataRelationshipsReservedStockData(val *POSTStockItems201ResponseDataRelationshipsReservedStockData) *NullablePOSTStockItems201ResponseDataRelationshipsReservedStockData {
	return &NullablePOSTStockItems201ResponseDataRelationshipsReservedStockData{value: val, isSet: true}
}

func (v NullablePOSTStockItems201ResponseDataRelationshipsReservedStockData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStockItems201ResponseDataRelationshipsReservedStockData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
