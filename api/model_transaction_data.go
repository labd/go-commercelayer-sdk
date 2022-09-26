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

// TransactionData struct for TransactionData
type TransactionData struct {
	// The resource's type
	Type          string                               `json:"type"`
	Attributes    RefundDataAttributes                 `json:"attributes"`
	Relationships *AdyenPaymentUpdateDataRelationships `json:"relationships,omitempty"`
}

// NewTransactionData instantiates a new TransactionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionData(type_ string, attributes RefundDataAttributes) *TransactionData {
	this := TransactionData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewTransactionDataWithDefaults instantiates a new TransactionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionDataWithDefaults() *TransactionData {
	this := TransactionData{}
	return &this
}

// GetType returns the Type field value
func (o *TransactionData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransactionData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransactionData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *TransactionData) GetAttributes() RefundDataAttributes {
	if o == nil {
		var ret RefundDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TransactionData) GetAttributesOk() (*RefundDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TransactionData) SetAttributes(v RefundDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TransactionData) GetRelationships() AdyenPaymentUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AdyenPaymentUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionData) GetRelationshipsOk() (*AdyenPaymentUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TransactionData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentUpdateDataRelationships and assigns it to the Relationships field.
func (o *TransactionData) SetRelationships(v AdyenPaymentUpdateDataRelationships) {
	o.Relationships = &v
}

func (o TransactionData) MarshalJSON() ([]byte, error) {
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

type NullableTransactionData struct {
	value *TransactionData
	isSet bool
}

func (v NullableTransactionData) Get() *TransactionData {
	return v.value
}

func (v *NullableTransactionData) Set(val *TransactionData) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionData) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionData(val *TransactionData) *NullableTransactionData {
	return &NullableTransactionData{value: val, isSet: true}
}

func (v NullableTransactionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
