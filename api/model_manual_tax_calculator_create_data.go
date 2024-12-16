/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ManualTaxCalculatorCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManualTaxCalculatorCreateData{}

// ManualTaxCalculatorCreateData struct for ManualTaxCalculatorCreateData
type ManualTaxCalculatorCreateData struct {
	// The resource's type
	Type          interface{}                                       `json:"type"`
	Attributes    POSTManualTaxCalculators201ResponseDataAttributes `json:"attributes"`
	Relationships *ManualTaxCalculatorCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewManualTaxCalculatorCreateData instantiates a new ManualTaxCalculatorCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualTaxCalculatorCreateData(type_ interface{}, attributes POSTManualTaxCalculators201ResponseDataAttributes) *ManualTaxCalculatorCreateData {
	this := ManualTaxCalculatorCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewManualTaxCalculatorCreateDataWithDefaults instantiates a new ManualTaxCalculatorCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualTaxCalculatorCreateDataWithDefaults() *ManualTaxCalculatorCreateData {
	this := ManualTaxCalculatorCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ManualTaxCalculatorCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualTaxCalculatorCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ManualTaxCalculatorCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ManualTaxCalculatorCreateData) GetAttributes() POSTManualTaxCalculators201ResponseDataAttributes {
	if o == nil {
		var ret POSTManualTaxCalculators201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorCreateData) GetAttributesOk() (*POSTManualTaxCalculators201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ManualTaxCalculatorCreateData) SetAttributes(v POSTManualTaxCalculators201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ManualTaxCalculatorCreateData) GetRelationships() ManualTaxCalculatorCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ManualTaxCalculatorCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorCreateData) GetRelationshipsOk() (*ManualTaxCalculatorCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ManualTaxCalculatorCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ManualTaxCalculatorCreateDataRelationships and assigns it to the Relationships field.
func (o *ManualTaxCalculatorCreateData) SetRelationships(v ManualTaxCalculatorCreateDataRelationships) {
	o.Relationships = &v
}

func (o ManualTaxCalculatorCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManualTaxCalculatorCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableManualTaxCalculatorCreateData struct {
	value *ManualTaxCalculatorCreateData
	isSet bool
}

func (v NullableManualTaxCalculatorCreateData) Get() *ManualTaxCalculatorCreateData {
	return v.value
}

func (v *NullableManualTaxCalculatorCreateData) Set(val *ManualTaxCalculatorCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableManualTaxCalculatorCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableManualTaxCalculatorCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualTaxCalculatorCreateData(val *ManualTaxCalculatorCreateData) *NullableManualTaxCalculatorCreateData {
	return &NullableManualTaxCalculatorCreateData{value: val, isSet: true}
}

func (v NullableManualTaxCalculatorCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualTaxCalculatorCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
