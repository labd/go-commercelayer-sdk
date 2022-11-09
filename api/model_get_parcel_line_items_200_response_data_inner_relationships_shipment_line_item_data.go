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

// GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData struct for GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData
type GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData instantiates a new GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData() *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData {
	this := GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData{}
	return &this
}

// NewGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemDataWithDefaults instantiates a new GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemDataWithDefaults() *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData {
	this := GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData) SetId(v string) {
	o.Id = &v
}

func (o GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData struct {
	value *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData
	isSet bool
}

func (v NullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData) Get() *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData {
	return v.value
}

func (v *NullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData) Set(val *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData(val *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData) *NullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData {
	return &NullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData{value: val, isSet: true}
}

func (v NullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
