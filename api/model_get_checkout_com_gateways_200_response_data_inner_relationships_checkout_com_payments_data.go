/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData struct for GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData
type GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData instantiates a new GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData() *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData {
	this := GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData{}
	return &this
}

// NewGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsDataWithDefaults instantiates a new GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsDataWithDefaults() *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData {
	this := GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData) SetId(v string) {
	o.Id = &v
}

func (o GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData struct {
	value *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData
	isSet bool
}

func (v NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData) Get() *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData {
	return v.value
}

func (v *NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData) Set(val *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData(val *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData) *NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData {
	return &NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData{value: val, isSet: true}
}

func (v NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
