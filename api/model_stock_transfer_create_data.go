/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// StockTransferCreateData struct for StockTransferCreateData
type StockTransferCreateData struct {
	// The resource's type
	Type          string                                `json:"type"`
	Attributes    StockTransferCreateDataAttributes     `json:"attributes"`
	Relationships *StockTransferCreateDataRelationships `json:"relationships,omitempty"`
}

// NewStockTransferCreateData instantiates a new StockTransferCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockTransferCreateData(type_ string, attributes StockTransferCreateDataAttributes) *StockTransferCreateData {
	this := StockTransferCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewStockTransferCreateDataWithDefaults instantiates a new StockTransferCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockTransferCreateDataWithDefaults() *StockTransferCreateData {
	this := StockTransferCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *StockTransferCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StockTransferCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StockTransferCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *StockTransferCreateData) GetAttributes() StockTransferCreateDataAttributes {
	if o == nil {
		var ret StockTransferCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *StockTransferCreateData) GetAttributesOk() (*StockTransferCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *StockTransferCreateData) SetAttributes(v StockTransferCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StockTransferCreateData) GetRelationships() StockTransferCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret StockTransferCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferCreateData) GetRelationshipsOk() (*StockTransferCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StockTransferCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given StockTransferCreateDataRelationships and assigns it to the Relationships field.
func (o *StockTransferCreateData) SetRelationships(v StockTransferCreateDataRelationships) {
	o.Relationships = &v
}

func (o StockTransferCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableStockTransferCreateData struct {
	value *StockTransferCreateData
	isSet bool
}

func (v NullableStockTransferCreateData) Get() *StockTransferCreateData {
	return v.value
}

func (v *NullableStockTransferCreateData) Set(val *StockTransferCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableStockTransferCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableStockTransferCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockTransferCreateData(val *StockTransferCreateData) *NullableStockTransferCreateData {
	return &NullableStockTransferCreateData{value: val, isSet: true}
}

func (v NullableStockTransferCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockTransferCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
