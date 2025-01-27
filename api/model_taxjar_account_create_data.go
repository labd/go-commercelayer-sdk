/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.6.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the TaxjarAccountCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxjarAccountCreateData{}

// TaxjarAccountCreateData struct for TaxjarAccountCreateData
type TaxjarAccountCreateData struct {
	// The resource's type
	Type          interface{}                                 `json:"type"`
	Attributes    POSTTaxjarAccounts201ResponseDataAttributes `json:"attributes"`
	Relationships *AvalaraAccountCreateDataRelationships      `json:"relationships,omitempty"`
}

// NewTaxjarAccountCreateData instantiates a new TaxjarAccountCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxjarAccountCreateData(type_ interface{}, attributes POSTTaxjarAccounts201ResponseDataAttributes) *TaxjarAccountCreateData {
	this := TaxjarAccountCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewTaxjarAccountCreateDataWithDefaults instantiates a new TaxjarAccountCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxjarAccountCreateDataWithDefaults() *TaxjarAccountCreateData {
	this := TaxjarAccountCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *TaxjarAccountCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxjarAccountCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaxjarAccountCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *TaxjarAccountCreateData) GetAttributes() POSTTaxjarAccounts201ResponseDataAttributes {
	if o == nil {
		var ret POSTTaxjarAccounts201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TaxjarAccountCreateData) GetAttributesOk() (*POSTTaxjarAccounts201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TaxjarAccountCreateData) SetAttributes(v POSTTaxjarAccounts201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TaxjarAccountCreateData) GetRelationships() AvalaraAccountCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AvalaraAccountCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxjarAccountCreateData) GetRelationshipsOk() (*AvalaraAccountCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TaxjarAccountCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AvalaraAccountCreateDataRelationships and assigns it to the Relationships field.
func (o *TaxjarAccountCreateData) SetRelationships(v AvalaraAccountCreateDataRelationships) {
	o.Relationships = &v
}

func (o TaxjarAccountCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxjarAccountCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableTaxjarAccountCreateData struct {
	value *TaxjarAccountCreateData
	isSet bool
}

func (v NullableTaxjarAccountCreateData) Get() *TaxjarAccountCreateData {
	return v.value
}

func (v *NullableTaxjarAccountCreateData) Set(val *TaxjarAccountCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxjarAccountCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxjarAccountCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxjarAccountCreateData(val *TaxjarAccountCreateData) *NullableTaxjarAccountCreateData {
	return &NullableTaxjarAccountCreateData{value: val, isSet: true}
}

func (v NullableTaxjarAccountCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxjarAccountCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
