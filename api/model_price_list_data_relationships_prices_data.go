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

// PriceListDataRelationshipsPricesData struct for PriceListDataRelationshipsPricesData
type PriceListDataRelationshipsPricesData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource's id
	Id *string `json:"id,omitempty"`
}

// NewPriceListDataRelationshipsPricesData instantiates a new PriceListDataRelationshipsPricesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceListDataRelationshipsPricesData() *PriceListDataRelationshipsPricesData {
	this := PriceListDataRelationshipsPricesData{}
	return &this
}

// NewPriceListDataRelationshipsPricesDataWithDefaults instantiates a new PriceListDataRelationshipsPricesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceListDataRelationshipsPricesDataWithDefaults() *PriceListDataRelationshipsPricesData {
	this := PriceListDataRelationshipsPricesData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PriceListDataRelationshipsPricesData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListDataRelationshipsPricesData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PriceListDataRelationshipsPricesData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PriceListDataRelationshipsPricesData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PriceListDataRelationshipsPricesData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListDataRelationshipsPricesData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PriceListDataRelationshipsPricesData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PriceListDataRelationshipsPricesData) SetId(v string) {
	o.Id = &v
}

func (o PriceListDataRelationshipsPricesData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullablePriceListDataRelationshipsPricesData struct {
	value *PriceListDataRelationshipsPricesData
	isSet bool
}

func (v NullablePriceListDataRelationshipsPricesData) Get() *PriceListDataRelationshipsPricesData {
	return v.value
}

func (v *NullablePriceListDataRelationshipsPricesData) Set(val *PriceListDataRelationshipsPricesData) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceListDataRelationshipsPricesData) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceListDataRelationshipsPricesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceListDataRelationshipsPricesData(val *PriceListDataRelationshipsPricesData) *NullablePriceListDataRelationshipsPricesData {
	return &NullablePriceListDataRelationshipsPricesData{value: val, isSet: true}
}

func (v NullablePriceListDataRelationshipsPricesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceListDataRelationshipsPricesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
