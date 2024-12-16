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

// checks if the POSTLineItems201ResponseDataRelationshipsStockTransfersData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTLineItems201ResponseDataRelationshipsStockTransfersData{}

// POSTLineItems201ResponseDataRelationshipsStockTransfersData struct for POSTLineItems201ResponseDataRelationshipsStockTransfersData
type POSTLineItems201ResponseDataRelationshipsStockTransfersData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTLineItems201ResponseDataRelationshipsStockTransfersData instantiates a new POSTLineItems201ResponseDataRelationshipsStockTransfersData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTLineItems201ResponseDataRelationshipsStockTransfersData() *POSTLineItems201ResponseDataRelationshipsStockTransfersData {
	this := POSTLineItems201ResponseDataRelationshipsStockTransfersData{}
	return &this
}

// NewPOSTLineItems201ResponseDataRelationshipsStockTransfersDataWithDefaults instantiates a new POSTLineItems201ResponseDataRelationshipsStockTransfersData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTLineItems201ResponseDataRelationshipsStockTransfersDataWithDefaults() *POSTLineItems201ResponseDataRelationshipsStockTransfersData {
	this := POSTLineItems201ResponseDataRelationshipsStockTransfersData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTLineItems201ResponseDataRelationshipsStockTransfersData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTLineItems201ResponseDataRelationshipsStockTransfersData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationshipsStockTransfersData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTLineItems201ResponseDataRelationshipsStockTransfersData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTLineItems201ResponseDataRelationshipsStockTransfersData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTLineItems201ResponseDataRelationshipsStockTransfersData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationshipsStockTransfersData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTLineItems201ResponseDataRelationshipsStockTransfersData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTLineItems201ResponseDataRelationshipsStockTransfersData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTLineItems201ResponseDataRelationshipsStockTransfersData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTLineItems201ResponseDataRelationshipsStockTransfersData struct {
	value *POSTLineItems201ResponseDataRelationshipsStockTransfersData
	isSet bool
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsStockTransfersData) Get() *POSTLineItems201ResponseDataRelationshipsStockTransfersData {
	return v.value
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsStockTransfersData) Set(val *POSTLineItems201ResponseDataRelationshipsStockTransfersData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsStockTransfersData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsStockTransfersData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTLineItems201ResponseDataRelationshipsStockTransfersData(val *POSTLineItems201ResponseDataRelationshipsStockTransfersData) *NullablePOSTLineItems201ResponseDataRelationshipsStockTransfersData {
	return &NullablePOSTLineItems201ResponseDataRelationshipsStockTransfersData{value: val, isSet: true}
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsStockTransfersData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsStockTransfersData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
