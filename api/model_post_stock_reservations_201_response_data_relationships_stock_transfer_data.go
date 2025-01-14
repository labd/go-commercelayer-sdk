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

// checks if the POSTStockReservations201ResponseDataRelationshipsStockTransferData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTStockReservations201ResponseDataRelationshipsStockTransferData{}

// POSTStockReservations201ResponseDataRelationshipsStockTransferData struct for POSTStockReservations201ResponseDataRelationshipsStockTransferData
type POSTStockReservations201ResponseDataRelationshipsStockTransferData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTStockReservations201ResponseDataRelationshipsStockTransferData instantiates a new POSTStockReservations201ResponseDataRelationshipsStockTransferData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStockReservations201ResponseDataRelationshipsStockTransferData() *POSTStockReservations201ResponseDataRelationshipsStockTransferData {
	this := POSTStockReservations201ResponseDataRelationshipsStockTransferData{}
	return &this
}

// NewPOSTStockReservations201ResponseDataRelationshipsStockTransferDataWithDefaults instantiates a new POSTStockReservations201ResponseDataRelationshipsStockTransferData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStockReservations201ResponseDataRelationshipsStockTransferDataWithDefaults() *POSTStockReservations201ResponseDataRelationshipsStockTransferData {
	this := POSTStockReservations201ResponseDataRelationshipsStockTransferData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStockReservations201ResponseDataRelationshipsStockTransferData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStockReservations201ResponseDataRelationshipsStockTransferData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTStockReservations201ResponseDataRelationshipsStockTransferData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTStockReservations201ResponseDataRelationshipsStockTransferData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStockReservations201ResponseDataRelationshipsStockTransferData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStockReservations201ResponseDataRelationshipsStockTransferData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTStockReservations201ResponseDataRelationshipsStockTransferData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTStockReservations201ResponseDataRelationshipsStockTransferData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTStockReservations201ResponseDataRelationshipsStockTransferData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTStockReservations201ResponseDataRelationshipsStockTransferData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTStockReservations201ResponseDataRelationshipsStockTransferData struct {
	value *POSTStockReservations201ResponseDataRelationshipsStockTransferData
	isSet bool
}

func (v NullablePOSTStockReservations201ResponseDataRelationshipsStockTransferData) Get() *POSTStockReservations201ResponseDataRelationshipsStockTransferData {
	return v.value
}

func (v *NullablePOSTStockReservations201ResponseDataRelationshipsStockTransferData) Set(val *POSTStockReservations201ResponseDataRelationshipsStockTransferData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStockReservations201ResponseDataRelationshipsStockTransferData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStockReservations201ResponseDataRelationshipsStockTransferData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStockReservations201ResponseDataRelationshipsStockTransferData(val *POSTStockReservations201ResponseDataRelationshipsStockTransferData) *NullablePOSTStockReservations201ResponseDataRelationshipsStockTransferData {
	return &NullablePOSTStockReservations201ResponseDataRelationshipsStockTransferData{value: val, isSet: true}
}

func (v NullablePOSTStockReservations201ResponseDataRelationshipsStockTransferData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStockReservations201ResponseDataRelationshipsStockTransferData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
