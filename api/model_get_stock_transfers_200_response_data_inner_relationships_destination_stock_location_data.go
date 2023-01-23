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

// GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData struct for GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData
type GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData instantiates a new GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData() *GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData {
	this := GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData{}
	return &this
}

// NewGETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationDataWithDefaults instantiates a new GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationDataWithDefaults() *GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData {
	this := GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData) SetId(v string) {
	o.Id = &v
}

func (o GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData struct {
	value *GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData
	isSet bool
}

func (v NullableGETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData) Get() *GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData {
	return v.value
}

func (v *NullableGETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData) Set(val *GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData(val *GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData) *NullableGETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData {
	return &NullableGETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData{value: val, isSet: true}
}

func (v NullableGETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
