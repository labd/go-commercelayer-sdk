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

// GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData struct for GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData
type GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData instantiates a new GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData() *GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData {
	this := GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData{}
	return &this
}

// NewGETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentDataWithDefaults instantiates a new GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentDataWithDefaults() *GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData {
	this := GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData) SetId(v string) {
	o.Id = &v
}

func (o GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData struct {
	value *GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData
	isSet bool
}

func (v NullableGETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData) Get() *GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData {
	return v.value
}

func (v *NullableGETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData) Set(val *GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData(val *GETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData) *NullableGETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData {
	return &NullableGETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData{value: val, isSet: true}
}

func (v NullableGETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShippingMethods200ResponseDataInnerRelationshipsDeliveryLeadTimeForShipmentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
