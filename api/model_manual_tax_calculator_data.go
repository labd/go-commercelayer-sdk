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

// ManualTaxCalculatorData struct for ManualTaxCalculatorData
type ManualTaxCalculatorData struct {
	// The resource's type
	Type          string                                                `json:"type"`
	Attributes    GETManualTaxCalculators200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *ManualTaxCalculatorDataRelationships                 `json:"relationships,omitempty"`
}

// NewManualTaxCalculatorData instantiates a new ManualTaxCalculatorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualTaxCalculatorData(type_ string, attributes GETManualTaxCalculators200ResponseDataInnerAttributes) *ManualTaxCalculatorData {
	this := ManualTaxCalculatorData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewManualTaxCalculatorDataWithDefaults instantiates a new ManualTaxCalculatorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualTaxCalculatorDataWithDefaults() *ManualTaxCalculatorData {
	this := ManualTaxCalculatorData{}
	return &this
}

// GetType returns the Type field value
func (o *ManualTaxCalculatorData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ManualTaxCalculatorData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ManualTaxCalculatorData) GetAttributes() GETManualTaxCalculators200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETManualTaxCalculators200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorData) GetAttributesOk() (*GETManualTaxCalculators200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ManualTaxCalculatorData) SetAttributes(v GETManualTaxCalculators200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ManualTaxCalculatorData) GetRelationships() ManualTaxCalculatorDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ManualTaxCalculatorDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorData) GetRelationshipsOk() (*ManualTaxCalculatorDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ManualTaxCalculatorData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ManualTaxCalculatorDataRelationships and assigns it to the Relationships field.
func (o *ManualTaxCalculatorData) SetRelationships(v ManualTaxCalculatorDataRelationships) {
	o.Relationships = &v
}

func (o ManualTaxCalculatorData) MarshalJSON() ([]byte, error) {
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

type NullableManualTaxCalculatorData struct {
	value *ManualTaxCalculatorData
	isSet bool
}

func (v NullableManualTaxCalculatorData) Get() *ManualTaxCalculatorData {
	return v.value
}

func (v *NullableManualTaxCalculatorData) Set(val *ManualTaxCalculatorData) {
	v.value = val
	v.isSet = true
}

func (v NullableManualTaxCalculatorData) IsSet() bool {
	return v.isSet
}

func (v *NullableManualTaxCalculatorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualTaxCalculatorData(val *ManualTaxCalculatorData) *NullableManualTaxCalculatorData {
	return &NullableManualTaxCalculatorData{value: val, isSet: true}
}

func (v NullableManualTaxCalculatorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualTaxCalculatorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
