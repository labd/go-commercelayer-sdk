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

// checks if the CustomerDataRelationshipsCustomerSubscriptionsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerDataRelationshipsCustomerSubscriptionsData{}

// CustomerDataRelationshipsCustomerSubscriptionsData struct for CustomerDataRelationshipsCustomerSubscriptionsData
type CustomerDataRelationshipsCustomerSubscriptionsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource's id
	Id interface{} `json:"id,omitempty"`
}

// NewCustomerDataRelationshipsCustomerSubscriptionsData instantiates a new CustomerDataRelationshipsCustomerSubscriptionsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerDataRelationshipsCustomerSubscriptionsData() *CustomerDataRelationshipsCustomerSubscriptionsData {
	this := CustomerDataRelationshipsCustomerSubscriptionsData{}
	return &this
}

// NewCustomerDataRelationshipsCustomerSubscriptionsDataWithDefaults instantiates a new CustomerDataRelationshipsCustomerSubscriptionsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerDataRelationshipsCustomerSubscriptionsDataWithDefaults() *CustomerDataRelationshipsCustomerSubscriptionsData {
	this := CustomerDataRelationshipsCustomerSubscriptionsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerDataRelationshipsCustomerSubscriptionsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerDataRelationshipsCustomerSubscriptionsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomerDataRelationshipsCustomerSubscriptionsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *CustomerDataRelationshipsCustomerSubscriptionsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerDataRelationshipsCustomerSubscriptionsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerDataRelationshipsCustomerSubscriptionsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerDataRelationshipsCustomerSubscriptionsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *CustomerDataRelationshipsCustomerSubscriptionsData) SetId(v interface{}) {
	o.Id = v
}

func (o CustomerDataRelationshipsCustomerSubscriptionsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerDataRelationshipsCustomerSubscriptionsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCustomerDataRelationshipsCustomerSubscriptionsData struct {
	value *CustomerDataRelationshipsCustomerSubscriptionsData
	isSet bool
}

func (v NullableCustomerDataRelationshipsCustomerSubscriptionsData) Get() *CustomerDataRelationshipsCustomerSubscriptionsData {
	return v.value
}

func (v *NullableCustomerDataRelationshipsCustomerSubscriptionsData) Set(val *CustomerDataRelationshipsCustomerSubscriptionsData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerDataRelationshipsCustomerSubscriptionsData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerDataRelationshipsCustomerSubscriptionsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerDataRelationshipsCustomerSubscriptionsData(val *CustomerDataRelationshipsCustomerSubscriptionsData) *NullableCustomerDataRelationshipsCustomerSubscriptionsData {
	return &NullableCustomerDataRelationshipsCustomerSubscriptionsData{value: val, isSet: true}
}

func (v NullableCustomerDataRelationshipsCustomerSubscriptionsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerDataRelationshipsCustomerSubscriptionsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
