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

// checks if the Store type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Store{}

// Store struct for Store
type Store struct {
	Data *StoreData `json:"data,omitempty"`
}

// NewStore instantiates a new Store object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStore() *Store {
	this := Store{}
	return &this
}

// NewStoreWithDefaults instantiates a new Store object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreWithDefaults() *Store {
	this := Store{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Store) GetData() StoreData {
	if o == nil || IsNil(o.Data) {
		var ret StoreData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Store) GetDataOk() (*StoreData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Store) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given StoreData and assigns it to the Data field.
func (o *Store) SetData(v StoreData) {
	o.Data = &v
}

func (o Store) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Store) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableStore struct {
	value *Store
	isSet bool
}

func (v NullableStore) Get() *Store {
	return v.value
}

func (v *NullableStore) Set(val *Store) {
	v.value = val
	v.isSet = true
}

func (v NullableStore) IsSet() bool {
	return v.isSet
}

func (v *NullableStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStore(val *Store) *NullableStore {
	return &NullableStore{value: val, isSet: true}
}

func (v NullableStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}