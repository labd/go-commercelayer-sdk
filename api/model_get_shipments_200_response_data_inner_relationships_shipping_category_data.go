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

// GETShipments200ResponseDataInnerRelationshipsShippingCategoryData struct for GETShipments200ResponseDataInnerRelationshipsShippingCategoryData
type GETShipments200ResponseDataInnerRelationshipsShippingCategoryData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETShipments200ResponseDataInnerRelationshipsShippingCategoryData instantiates a new GETShipments200ResponseDataInnerRelationshipsShippingCategoryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShipments200ResponseDataInnerRelationshipsShippingCategoryData() *GETShipments200ResponseDataInnerRelationshipsShippingCategoryData {
	this := GETShipments200ResponseDataInnerRelationshipsShippingCategoryData{}
	return &this
}

// NewGETShipments200ResponseDataInnerRelationshipsShippingCategoryDataWithDefaults instantiates a new GETShipments200ResponseDataInnerRelationshipsShippingCategoryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShipments200ResponseDataInnerRelationshipsShippingCategoryDataWithDefaults() *GETShipments200ResponseDataInnerRelationshipsShippingCategoryData {
	this := GETShipments200ResponseDataInnerRelationshipsShippingCategoryData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationshipsShippingCategoryData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationshipsShippingCategoryData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationshipsShippingCategoryData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETShipments200ResponseDataInnerRelationshipsShippingCategoryData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationshipsShippingCategoryData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationshipsShippingCategoryData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationshipsShippingCategoryData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETShipments200ResponseDataInnerRelationshipsShippingCategoryData) SetId(v string) {
	o.Id = &v
}

func (o GETShipments200ResponseDataInnerRelationshipsShippingCategoryData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETShipments200ResponseDataInnerRelationshipsShippingCategoryData struct {
	value *GETShipments200ResponseDataInnerRelationshipsShippingCategoryData
	isSet bool
}

func (v NullableGETShipments200ResponseDataInnerRelationshipsShippingCategoryData) Get() *GETShipments200ResponseDataInnerRelationshipsShippingCategoryData {
	return v.value
}

func (v *NullableGETShipments200ResponseDataInnerRelationshipsShippingCategoryData) Set(val *GETShipments200ResponseDataInnerRelationshipsShippingCategoryData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShipments200ResponseDataInnerRelationshipsShippingCategoryData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShipments200ResponseDataInnerRelationshipsShippingCategoryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShipments200ResponseDataInnerRelationshipsShippingCategoryData(val *GETShipments200ResponseDataInnerRelationshipsShippingCategoryData) *NullableGETShipments200ResponseDataInnerRelationshipsShippingCategoryData {
	return &NullableGETShipments200ResponseDataInnerRelationshipsShippingCategoryData{value: val, isSet: true}
}

func (v NullableGETShipments200ResponseDataInnerRelationshipsShippingCategoryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShipments200ResponseDataInnerRelationshipsShippingCategoryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
