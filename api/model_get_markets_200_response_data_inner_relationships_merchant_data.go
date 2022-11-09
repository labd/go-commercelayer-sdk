/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETMarkets200ResponseDataInnerRelationshipsMerchantData struct for GETMarkets200ResponseDataInnerRelationshipsMerchantData
type GETMarkets200ResponseDataInnerRelationshipsMerchantData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETMarkets200ResponseDataInnerRelationshipsMerchantData instantiates a new GETMarkets200ResponseDataInnerRelationshipsMerchantData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETMarkets200ResponseDataInnerRelationshipsMerchantData() *GETMarkets200ResponseDataInnerRelationshipsMerchantData {
	this := GETMarkets200ResponseDataInnerRelationshipsMerchantData{}
	return &this
}

// NewGETMarkets200ResponseDataInnerRelationshipsMerchantDataWithDefaults instantiates a new GETMarkets200ResponseDataInnerRelationshipsMerchantData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETMarkets200ResponseDataInnerRelationshipsMerchantDataWithDefaults() *GETMarkets200ResponseDataInnerRelationshipsMerchantData {
	this := GETMarkets200ResponseDataInnerRelationshipsMerchantData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchantData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchantData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchantData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchantData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchantData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchantData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchantData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchantData) SetId(v string) {
	o.Id = &v
}

func (o GETMarkets200ResponseDataInnerRelationshipsMerchantData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETMarkets200ResponseDataInnerRelationshipsMerchantData struct {
	value *GETMarkets200ResponseDataInnerRelationshipsMerchantData
	isSet bool
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsMerchantData) Get() *GETMarkets200ResponseDataInnerRelationshipsMerchantData {
	return v.value
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsMerchantData) Set(val *GETMarkets200ResponseDataInnerRelationshipsMerchantData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsMerchantData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsMerchantData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETMarkets200ResponseDataInnerRelationshipsMerchantData(val *GETMarkets200ResponseDataInnerRelationshipsMerchantData) *NullableGETMarkets200ResponseDataInnerRelationshipsMerchantData {
	return &NullableGETMarkets200ResponseDataInnerRelationshipsMerchantData{value: val, isSet: true}
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsMerchantData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsMerchantData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}