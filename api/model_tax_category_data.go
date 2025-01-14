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

// checks if the TaxCategoryData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxCategoryData{}

// TaxCategoryData struct for TaxCategoryData
type TaxCategoryData struct {
	// The resource's type
	Type          interface{}                                            `json:"type"`
	Attributes    GETTaxCategoriesTaxCategoryId200ResponseDataAttributes `json:"attributes"`
	Relationships *TaxCategoryDataRelationships                          `json:"relationships,omitempty"`
}

// NewTaxCategoryData instantiates a new TaxCategoryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxCategoryData(type_ interface{}, attributes GETTaxCategoriesTaxCategoryId200ResponseDataAttributes) *TaxCategoryData {
	this := TaxCategoryData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewTaxCategoryDataWithDefaults instantiates a new TaxCategoryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxCategoryDataWithDefaults() *TaxCategoryData {
	this := TaxCategoryData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *TaxCategoryData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxCategoryData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaxCategoryData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *TaxCategoryData) GetAttributes() GETTaxCategoriesTaxCategoryId200ResponseDataAttributes {
	if o == nil {
		var ret GETTaxCategoriesTaxCategoryId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TaxCategoryData) GetAttributesOk() (*GETTaxCategoriesTaxCategoryId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TaxCategoryData) SetAttributes(v GETTaxCategoriesTaxCategoryId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TaxCategoryData) GetRelationships() TaxCategoryDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret TaxCategoryDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCategoryData) GetRelationshipsOk() (*TaxCategoryDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TaxCategoryData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given TaxCategoryDataRelationships and assigns it to the Relationships field.
func (o *TaxCategoryData) SetRelationships(v TaxCategoryDataRelationships) {
	o.Relationships = &v
}

func (o TaxCategoryData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxCategoryData) ToMap() (map[string]interface{}, error) {
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

type NullableTaxCategoryData struct {
	value *TaxCategoryData
	isSet bool
}

func (v NullableTaxCategoryData) Get() *TaxCategoryData {
	return v.value
}

func (v *NullableTaxCategoryData) Set(val *TaxCategoryData) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxCategoryData) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxCategoryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxCategoryData(val *TaxCategoryData) *NullableTaxCategoryData {
	return &NullableTaxCategoryData{value: val, isSet: true}
}

func (v NullableTaxCategoryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxCategoryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
