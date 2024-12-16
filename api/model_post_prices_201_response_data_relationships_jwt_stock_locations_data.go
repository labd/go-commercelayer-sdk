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

// checks if the POSTPrices201ResponseDataRelationshipsJwtStockLocationsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPrices201ResponseDataRelationshipsJwtStockLocationsData{}

// POSTPrices201ResponseDataRelationshipsJwtStockLocationsData struct for POSTPrices201ResponseDataRelationshipsJwtStockLocationsData
type POSTPrices201ResponseDataRelationshipsJwtStockLocationsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTPrices201ResponseDataRelationshipsJwtStockLocationsData instantiates a new POSTPrices201ResponseDataRelationshipsJwtStockLocationsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPrices201ResponseDataRelationshipsJwtStockLocationsData() *POSTPrices201ResponseDataRelationshipsJwtStockLocationsData {
	this := POSTPrices201ResponseDataRelationshipsJwtStockLocationsData{}
	return &this
}

// NewPOSTPrices201ResponseDataRelationshipsJwtStockLocationsDataWithDefaults instantiates a new POSTPrices201ResponseDataRelationshipsJwtStockLocationsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPrices201ResponseDataRelationshipsJwtStockLocationsDataWithDefaults() *POSTPrices201ResponseDataRelationshipsJwtStockLocationsData {
	this := POSTPrices201ResponseDataRelationshipsJwtStockLocationsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPrices201ResponseDataRelationshipsJwtStockLocationsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPrices201ResponseDataRelationshipsJwtStockLocationsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationshipsJwtStockLocationsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTPrices201ResponseDataRelationshipsJwtStockLocationsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPrices201ResponseDataRelationshipsJwtStockLocationsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPrices201ResponseDataRelationshipsJwtStockLocationsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationshipsJwtStockLocationsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTPrices201ResponseDataRelationshipsJwtStockLocationsData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTPrices201ResponseDataRelationshipsJwtStockLocationsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPrices201ResponseDataRelationshipsJwtStockLocationsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTPrices201ResponseDataRelationshipsJwtStockLocationsData struct {
	value *POSTPrices201ResponseDataRelationshipsJwtStockLocationsData
	isSet bool
}

func (v NullablePOSTPrices201ResponseDataRelationshipsJwtStockLocationsData) Get() *POSTPrices201ResponseDataRelationshipsJwtStockLocationsData {
	return v.value
}

func (v *NullablePOSTPrices201ResponseDataRelationshipsJwtStockLocationsData) Set(val *POSTPrices201ResponseDataRelationshipsJwtStockLocationsData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPrices201ResponseDataRelationshipsJwtStockLocationsData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPrices201ResponseDataRelationshipsJwtStockLocationsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPrices201ResponseDataRelationshipsJwtStockLocationsData(val *POSTPrices201ResponseDataRelationshipsJwtStockLocationsData) *NullablePOSTPrices201ResponseDataRelationshipsJwtStockLocationsData {
	return &NullablePOSTPrices201ResponseDataRelationshipsJwtStockLocationsData{value: val, isSet: true}
}

func (v NullablePOSTPrices201ResponseDataRelationshipsJwtStockLocationsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPrices201ResponseDataRelationshipsJwtStockLocationsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
