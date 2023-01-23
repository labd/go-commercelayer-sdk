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

// GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData struct for GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData
type GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData instantiates a new GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData() *GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData {
	this := GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData{}
	return &this
}

// NewGETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesDataWithDefaults instantiates a new GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesDataWithDefaults() *GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData {
	this := GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData) SetId(v string) {
	o.Id = &v
}

func (o GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData struct {
	value *GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData
	isSet bool
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData) Get() *GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData {
	return v.value
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData) Set(val *GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData(val *GETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData) *NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData {
	return &NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData{value: val, isSet: true}
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeBundlesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
