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

// checks if the POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData{}

// POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData struct for POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData
type POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTPrices201ResponseDataRelationshipsPriceVolumeTiersData instantiates a new POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPrices201ResponseDataRelationshipsPriceVolumeTiersData() *POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData {
	this := POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData{}
	return &this
}

// NewPOSTPrices201ResponseDataRelationshipsPriceVolumeTiersDataWithDefaults instantiates a new POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPrices201ResponseDataRelationshipsPriceVolumeTiersDataWithDefaults() *POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData {
	this := POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTPrices201ResponseDataRelationshipsPriceVolumeTiersData struct {
	value *POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData
	isSet bool
}

func (v NullablePOSTPrices201ResponseDataRelationshipsPriceVolumeTiersData) Get() *POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData {
	return v.value
}

func (v *NullablePOSTPrices201ResponseDataRelationshipsPriceVolumeTiersData) Set(val *POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPrices201ResponseDataRelationshipsPriceVolumeTiersData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPrices201ResponseDataRelationshipsPriceVolumeTiersData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPrices201ResponseDataRelationshipsPriceVolumeTiersData(val *POSTPrices201ResponseDataRelationshipsPriceVolumeTiersData) *NullablePOSTPrices201ResponseDataRelationshipsPriceVolumeTiersData {
	return &NullablePOSTPrices201ResponseDataRelationshipsPriceVolumeTiersData{value: val, isSet: true}
}

func (v NullablePOSTPrices201ResponseDataRelationshipsPriceVolumeTiersData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPrices201ResponseDataRelationshipsPriceVolumeTiersData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
