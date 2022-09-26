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

// ExternalTaxCalculatorData struct for ExternalTaxCalculatorData
type ExternalTaxCalculatorData struct {
	// The resource's type
	Type          string                              `json:"type"`
	Attributes    ExternalTaxCalculatorDataAttributes `json:"attributes"`
	Relationships *AvalaraAccountDataRelationships    `json:"relationships,omitempty"`
}

// NewExternalTaxCalculatorData instantiates a new ExternalTaxCalculatorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalTaxCalculatorData(type_ string, attributes ExternalTaxCalculatorDataAttributes) *ExternalTaxCalculatorData {
	this := ExternalTaxCalculatorData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewExternalTaxCalculatorDataWithDefaults instantiates a new ExternalTaxCalculatorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalTaxCalculatorDataWithDefaults() *ExternalTaxCalculatorData {
	this := ExternalTaxCalculatorData{}
	return &this
}

// GetType returns the Type field value
func (o *ExternalTaxCalculatorData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExternalTaxCalculatorData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalTaxCalculatorData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ExternalTaxCalculatorData) GetAttributes() ExternalTaxCalculatorDataAttributes {
	if o == nil {
		var ret ExternalTaxCalculatorDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ExternalTaxCalculatorData) GetAttributesOk() (*ExternalTaxCalculatorDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ExternalTaxCalculatorData) SetAttributes(v ExternalTaxCalculatorDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ExternalTaxCalculatorData) GetRelationships() AvalaraAccountDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AvalaraAccountDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalTaxCalculatorData) GetRelationshipsOk() (*AvalaraAccountDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ExternalTaxCalculatorData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AvalaraAccountDataRelationships and assigns it to the Relationships field.
func (o *ExternalTaxCalculatorData) SetRelationships(v AvalaraAccountDataRelationships) {
	o.Relationships = &v
}

func (o ExternalTaxCalculatorData) MarshalJSON() ([]byte, error) {
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

type NullableExternalTaxCalculatorData struct {
	value *ExternalTaxCalculatorData
	isSet bool
}

func (v NullableExternalTaxCalculatorData) Get() *ExternalTaxCalculatorData {
	return v.value
}

func (v *NullableExternalTaxCalculatorData) Set(val *ExternalTaxCalculatorData) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalTaxCalculatorData) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalTaxCalculatorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalTaxCalculatorData(val *ExternalTaxCalculatorData) *NullableExternalTaxCalculatorData {
	return &NullableExternalTaxCalculatorData{value: val, isSet: true}
}

func (v NullableExternalTaxCalculatorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalTaxCalculatorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
