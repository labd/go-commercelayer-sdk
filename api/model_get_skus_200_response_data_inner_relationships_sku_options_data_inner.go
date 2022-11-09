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

// GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner struct for GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner
type GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner instantiates a new GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner() *GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner {
	this := GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner{}
	return &this
}

// NewGETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInnerWithDefaults instantiates a new GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInnerWithDefaults() *GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner {
	this := GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner) SetId(v string) {
	o.Id = &v
}

func (o GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner struct {
	value *GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner
	isSet bool
}

func (v NullableGETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner) Get() *GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner {
	return v.value
}

func (v *NullableGETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner) Set(val *GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner(val *GETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner) *NullableGETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner {
	return &NullableGETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner{value: val, isSet: true}
}

func (v NullableGETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSkus200ResponseDataInnerRelationshipsSkuOptionsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
