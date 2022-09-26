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

// TaxCalculatorResponseData struct for TaxCalculatorResponseData
type TaxCalculatorResponseData struct {
	// The resource's id
	Id *string `json:"id,omitempty"`
	// The resource's type
	Type          *string                                  `json:"type,omitempty"`
	Links         *AddressResponseDataLinks                `json:"links,omitempty"`
	Attributes    *Attributes                              `json:"attributes,omitempty"`
	Relationships *AvalaraAccountResponseDataRelationships `json:"relationships,omitempty"`
}

// NewTaxCalculatorResponseData instantiates a new TaxCalculatorResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxCalculatorResponseData() *TaxCalculatorResponseData {
	this := TaxCalculatorResponseData{}
	return &this
}

// NewTaxCalculatorResponseDataWithDefaults instantiates a new TaxCalculatorResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxCalculatorResponseDataWithDefaults() *TaxCalculatorResponseData {
	this := TaxCalculatorResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaxCalculatorResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCalculatorResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaxCalculatorResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TaxCalculatorResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TaxCalculatorResponseData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCalculatorResponseData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TaxCalculatorResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TaxCalculatorResponseData) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TaxCalculatorResponseData) GetLinks() AddressResponseDataLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCalculatorResponseData) GetLinksOk() (*AddressResponseDataLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TaxCalculatorResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataLinks and assigns it to the Links field.
func (o *TaxCalculatorResponseData) SetLinks(v AddressResponseDataLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TaxCalculatorResponseData) GetAttributes() Attributes {
	if o == nil || o.Attributes == nil {
		var ret Attributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCalculatorResponseData) GetAttributesOk() (*Attributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TaxCalculatorResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given Attributes and assigns it to the Attributes field.
func (o *TaxCalculatorResponseData) SetAttributes(v Attributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TaxCalculatorResponseData) GetRelationships() AvalaraAccountResponseDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AvalaraAccountResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCalculatorResponseData) GetRelationshipsOk() (*AvalaraAccountResponseDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TaxCalculatorResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AvalaraAccountResponseDataRelationships and assigns it to the Relationships field.
func (o *TaxCalculatorResponseData) SetRelationships(v AvalaraAccountResponseDataRelationships) {
	o.Relationships = &v
}

func (o TaxCalculatorResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableTaxCalculatorResponseData struct {
	value *TaxCalculatorResponseData
	isSet bool
}

func (v NullableTaxCalculatorResponseData) Get() *TaxCalculatorResponseData {
	return v.value
}

func (v *NullableTaxCalculatorResponseData) Set(val *TaxCalculatorResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxCalculatorResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxCalculatorResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxCalculatorResponseData(val *TaxCalculatorResponseData) *NullableTaxCalculatorResponseData {
	return &NullableTaxCalculatorResponseData{value: val, isSet: true}
}

func (v NullableTaxCalculatorResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxCalculatorResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
