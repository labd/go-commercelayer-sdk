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

// AdyenPaymentCreateData struct for AdyenPaymentCreateData
type AdyenPaymentCreateData struct {
	// The resource's type
	Type          string                                     `json:"type"`
	Attributes    POSTAdyenPayments201ResponseDataAttributes `json:"attributes"`
	Relationships *AdyenPaymentCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewAdyenPaymentCreateData instantiates a new AdyenPaymentCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenPaymentCreateData(type_ string, attributes POSTAdyenPayments201ResponseDataAttributes) *AdyenPaymentCreateData {
	this := AdyenPaymentCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAdyenPaymentCreateDataWithDefaults instantiates a new AdyenPaymentCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenPaymentCreateDataWithDefaults() *AdyenPaymentCreateData {
	this := AdyenPaymentCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *AdyenPaymentCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AdyenPaymentCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AdyenPaymentCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AdyenPaymentCreateData) GetAttributes() POSTAdyenPayments201ResponseDataAttributes {
	if o == nil {
		var ret POSTAdyenPayments201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AdyenPaymentCreateData) GetAttributesOk() (*POSTAdyenPayments201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AdyenPaymentCreateData) SetAttributes(v POSTAdyenPayments201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AdyenPaymentCreateData) GetRelationships() AdyenPaymentCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AdyenPaymentCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentCreateData) GetRelationshipsOk() (*AdyenPaymentCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AdyenPaymentCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentCreateDataRelationships and assigns it to the Relationships field.
func (o *AdyenPaymentCreateData) SetRelationships(v AdyenPaymentCreateDataRelationships) {
	o.Relationships = &v
}

func (o AdyenPaymentCreateData) MarshalJSON() ([]byte, error) {
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

type NullableAdyenPaymentCreateData struct {
	value *AdyenPaymentCreateData
	isSet bool
}

func (v NullableAdyenPaymentCreateData) Get() *AdyenPaymentCreateData {
	return v.value
}

func (v *NullableAdyenPaymentCreateData) Set(val *AdyenPaymentCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenPaymentCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenPaymentCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenPaymentCreateData(val *AdyenPaymentCreateData) *NullableAdyenPaymentCreateData {
	return &NullableAdyenPaymentCreateData{value: val, isSet: true}
}

func (v NullableAdyenPaymentCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenPaymentCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
